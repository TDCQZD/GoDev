# goroutine
## goroutine相比java 多线程NIO有何优势 ？
NIO（非阻塞IO）是一种IO编程模型，Golang中的IO底层实现方式和java NIO模型一致，通俗点说就是都采用了EPOLL。 
你在使用golang读文件的时候，goroutine 会默默的挂起，只是你不知道，当读完毕了，goroutine 再次恢复，但你不用担心，goroutine 的挂起和恢复没有java线程那样可怕，你可以认为goroutine 的挂起和恢复就是保存和恢复几个变量的值，其实也是这样的。

剩下的就是goroutine 和 java线程的区别了，goroutine是用户态的线程切换，java采用的是系统线程切换，用汇编语言描述是一个(java)调用int 80软中断,一个没有。 意味着goroutine更轻量级，可以同时相应成千上万的线程切换，java你创造上千个线程就有些吃力了。

## Goroutine与线程的区别

许多人认为goroutine比线程运行得更快，这是一个误解。Goroutine并不会更快，它只是增加了更多的并发性。当一个goroutine被阻塞（比如等待IO），golang的scheduler会调度其它可以执行的goroutine运行。与线程相比，它有以下几个优点：

- 内存消耗更少：

Goroutine所需要的内存通常只有2kb，而线程则需要1Mb（500倍）。

- 创建与销毁的开销更小

由于线程创建时需要向操作系统申请资源，并且在销毁时将资源归还，因此它的创建和销毁的开销比较大。相比之下，goroutine的创建和销毁是由go语言在运行时自己管理的，因此开销更低。

- 切换开销更小
    * 线程的调度方式是抢占式的
    * goroutine的调度是协同式的
    
这是goroutine与线程的主要区别，也是golang能够实现高并发的主要原因。线程的调度方式是抢占式的，如果一个线程的执行时间超过了分配给它的时间片，就会被其它可执行的线程抢占。在线程切换的过程中需要保存/恢复所有的寄存器信息，比如16个通用寄存器，PC（Program Counter），SP（Stack Pointer），段寄存器等等。

而goroutine的调度是协同式的，它不会直接地与操作系统内核打交道。当goroutine进行切换的时候，之后很少量的寄存器需要保存和恢复（PC和SP）。因此gouroutine的切换效率更高。

## select
golang 的 select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。 

在执行select语句的时候，运行时系统会自上而下地判断每个case中的发送或接收操作是否可以被立即执行(立即执行：意思是当前Goroutine不会因此操作而被阻塞)

select的用法与switch非常类似，由select开始一个新的选择块，每个选择条件由case语句来描述。与switch语句可以选择任何可使用相等比较的条件相比，select有比较多的限制，其中最大的一条限制就是**每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作**。

## goroutine的超时控制
### 1、利用select实现goroutine的超时控制
select是通过线性扫描的方式监视文件描述符是否有变动. channnel在系统层面来说也是个文件描述符。 在golang里我们可以使用goroutine并发执行任务，接着使用select来监视每个任务的channel情况.  

但如果这几个任务都长时间没有回复channel信息，如果我们又有超时timeout需求，那么我们可以使用起一个goroutine，这个goroutine任务逻辑启动sleep,等sleep之后回复channel信号。 
```
func main() {
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(3 * time.Second) // sleep 3 second
        timeout <- true
    }()
    ch := make(chan int)
    select {
    case <-ch:
    case <-timeout:
        fmt.Println("task is timeout!")
    }
} 
```

## goroutine退出
- 1、channel＋select
```
go func() {
    time.Sleep(1e9)
    timeout <- true
} ()

select {
    case <- ch:
        //从ch中读取数据
    case <-timeout:
        //ch一直没有数据写入，超时触发timeout
}
```
- 2、context包来实现上下文功能
```
  select {
        case <-cxt.Done():
            // do some clean...
    }
```
