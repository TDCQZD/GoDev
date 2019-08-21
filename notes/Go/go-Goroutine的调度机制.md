# Goroutine的调度机制
Golang 在语言层面支持多线程，代码可以通过 go 关键字来启动 Goroutine ，调用者不需要关心调用栈的大小，函数上下文等等信息就可以完成并发或者并行操作，加快了我们的开发速度。 
## Scheduler Structure
整个调度模型由 Goroutine/Processor/Machine 以及全局调度信息 sched 组成

### 1、 Goroutine
Goroutine 是 Golang 里的 线程 ，同样也是可调度的单元。
```
// src/runtime/runtime2.go
type g struct {
        ....
        m       *m
        sched gobuf
        goid   int64
        ....
}

type gobuf struct {
        sp   uintptr
        pc   uintptr
        ....
}
```
runtime 为 Goroutine 引入了类似 PID 的属性 goid ，使得每一个 Goroutine 都有全局唯一的 goid 标识。
> 注意：不过官方并没有提供接口能 直接 访问当前 Goroutine 的 goid
在调度过程中，runtime 需要 Goroutine 释放当前的计算资源，为了保证下次能恢复现场，执行的上下文现场（指令地址 和 Stack Pointer 等）将会存储在 gobuf 这个数据结构中。

整体来说，Goroutine 仅代表任务的内容以及上下文，并不是具体的执行单元。
### 2、 Machine
Machine 是 OS Thread，它负责执行 Goroutine。
```
// src/runtime/runtime2.go

type m struct {
        ....
        g0      *g     // goroutine with scheduling stack
        curg    *g     // current running goroutine

        tls     [6]uintptr // thread-local storage (for x86 extern register)
        p       puintptr // attached p for executing go code (nil if not executing go code)
        ....
}
```
runtime 在做调度工作或者和当前 Goroutine 无关的任务时，Golang 会切换调用栈来进行相关的任务，就好像 Linux 的进程进入系统调用时会切换到内核态的调用栈一样，这么做也是为了避免影响到调度以及垃圾回收的扫描。

Machine 一般会调用 systemstack 函数 来切换调用栈。 

从名字可以看出，Golang 对外部 go code 的调用栈称之为 user stack ，而将运行核心 runtime 部分代码的调用栈称之为 system stack。 

Machine 需要维护这两个调用栈的上下文，所以 m 中 g0 用来代表 runtime 内部逻辑，而 curg 则是我们平时写的代码

因为调用栈可以来回地切换，Machine 需要知道当前运行的调用栈信息，所以 Golang 会利用 Thread Local Storage 或者指定寄存器来存储当前运行的 g。

settls 汇编代码会将 g 的地址放到 m.tls 中，这样 Machine 就可以通过 getg 取出当前运行的 Goroutine。

> 不同平台 settls 的行为有一定差别。
```
// src/runtime/sys_linux_amd64.s

// set tls base to DI
TEXT runtime·settls(SB),NOSPLIT,$32
#ifdef GOOS_android
        // Same as in sys_darwin_386.s:/ugliness, different constant.
        // DI currently holds m->tls, which must be fs:0x1d0.
        // See cgo/gcc_android_amd64.c for the derivation of the constant.
        SUBQ    $0x1d0, DI  // In android, the tls base·
#else
        ADDQ    $8, DI  // ELF wants to use -8(FS)
#endif
        MOVQ    DI, SI
        MOVQ    $0x1002, DI     // ARCH_SET_FS
        MOVQ    $158, AX        // arch_prctl
        SYSCALL
        CMPQ    AX, $0xfffffffffffff001
        JLS     2(PC)
        MOVL    $0xf1, 0xf1  // crash
        RET

// src/runtime/stubs.go

// getg returns the pointer to the current g.
// The compiler rewrites calls to this function into instructions
// that fetch the g directly (from TLS or from the dedicated register).
func getg() *g

// src/runtime/go_tls.h 

#ifdef GOARCH_amd64
#define get_tls(r)      MOVQ TLS, r
#define g(r)    0(r)(TLS*1)
#endif
```
但是 Machine 想要执行一个 Goroutine，必须要绑定 Processor。

runtime 内部有些函数执行时会直接绑定 Machine，并不需要 Processor，比如 sysmon 。
### 3、Processor
Processor 可以理解成处理器，它会维护着本地 Goroutine 队列 runq ，并在新的 Goroutine 入队列时分配唯一的 goid。
```
type p struct {
        ...
        m           muintptr   // back-link to associated m (nil if idle)

        // Cache of goroutine ids, amortizes accesses to runtime·sched.goidgen.
        goidcache    uint64
        goidcacheend uint64

        // Queue of runnable goroutines. Accessed without lock.
        runqhead uint32
        runqtail uint32
        runq     [256]guintptr
        ...
}
```
Processor 的数目代表着 runtime 能同时处理 Goroutine 的数目，GOMAXPROCS 环境变量是用来指定 Processor 的数目，默认状态会是 CPU 的个数。

也正是因为 Processor 的存在，runtime 并不需要做一个集中式的 Goroutine 调度，每一个 Machine 都会在 Processor 本地队列、Global Runnable Queue 或者其他 Processor 队列中找 Goroutine 执行，减少全局锁对性能的影响.
### 4、全局调度信息 sched
全局调度信息 sched 会记录当前 Global Runnable Queue、当前空闲的 Machine 和空闲 Processor 的数目等等
- goidgen：
- nmspinning：
```
// src/runtime/runtime2.go

var (
        ...
        sched      schedt
        ...
)

type schedt struct {
        // accessed atomically. keep at top to ensure alignment on 32-bit systems.
        goidgen  uint64

        lock mutex

        midle        muintptr // idle m's waiting for work
        nmidle       int32    // number of idle m's waiting for work
        maxmcount    int32    // maximum number of m's allowed (or die)

        pidle      puintptr // idle p's
        npidle     uint32
        nmspinning uint32 // See "Worker thread parking/unparking" comment in proc.go.

        // Global runnable queue.
        runqhead guintptr
        runqtail guintptr
        runqsize int32
        ....
}
``` 