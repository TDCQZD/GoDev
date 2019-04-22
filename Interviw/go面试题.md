# go 面试题整理
## 程序代码没问题(不是变量, 资源未释放等问题), 内存占用变得很高, 正常情况下它不应该这么高的, 请问这是什么问题, 如何解决?
> 资源释放的时候, 那一部分内存空间还没还给系统, 导致系统内存升高, 可以用pprof工具查看

这个问题在我们对后台服务进行压力测试时发现，我们模拟大量的用户请求访问后台服务，这时各服务模块能观察到明显的内存占用上升。但是当停止压测时，内存占用并未发生明显的下降。花了很长时间定位问题，使用gprof等各种方法，依然没有发现原因。最后发现原来这时正常的…主要的原因有两个，

- 一是go的垃圾回收有个触发阈值，这个阈值会随着每次内存使用变大而逐渐增大（如初始阈值是10MB则下一次就是20MB，再下一次就成为了40MB…），如果长时间没有触发gc go会主动触发一次（2min）。高峰时内存使用量上去后，除非持续申请内存，靠阈值触发gc已经基本不可能，而是要等最多2min主动gc开始才能触发gc。

- 第二个原因是go语言在向系统交还内存时只是告诉系统这些内存不需要使用了，可以回收；同时操作系统会采取“拖延症”策略，并不是立即回收，而是等到系统内存紧张时才会开始回收这样该程序又重新申请内存时就可以获得极快的分配速度。
## 分别描述三种情况下的输出结果以及导致的原因

1.First Second 均注释

2.First 打开 Second 注释

3.First Second 均打开
```
func main() {
    //runtime.GOMAXPROCS(1) //First
    exit := make(chan int)
    go func() {
        close(exit)
        for {
            if true {
                //println("Looping!")  //Second
            }
        }
    }()
    <-exit
    println("Am I printed?")
}
```
- 第一种情况：打印最后一行，然后退出。channel 关闭，之后，读channel 读出默认值 
- 第二种情况会挂起，
- 第三种情况会一段时间后退出。
> 第二和第三种情况设计到go 的GC 机制，没有函数调用的死循环不会被GC。

**扩展**: [go语言死循环分析](https://juejin.im/entry/5a263c9e51882535c56cc42f)
## [滴滴]select中实现优先级问题
> select中事件是随机触发的
可以使用default实现
```
for msg := range c.incomingMsgChan {
    select {
    case c.memoryMsgChan <- msg:
    default:
        err := WriteMessageToBackend(&msgBuf, msg, c.backend)
        if err != nil {
            // ... handle errors ...
        }
    }
}
Taking advantage of Go’s select statement allows this functionality to be expressed in just a few lines of code: the default case above only executes if memoryMsgChan is full.
利用Go的select语句，只需几行代码即可表示此功能：上述默认情况仅在memoryMsgChan已满时执行。
```
## unsafe.Pointer 面试题
问 fun1 和 fun2 fun3分别输出什么，为什么?
```
func fun1()  {
    a := 2
    c := (*string) (unsafe.Pointer(&a))
    *c = "44"
    fmt.Println(*c)
}
func fun2()  {
    a := "654"
    c := (*string) (unsafe.Pointer(&a))
    *c = "44"
    fmt.Println(*c)
}
func fun3()  {
    a := 3
    c := *(*string) (unsafe.Pointer(&a))
    c = "445"
    fmt.Println(c)
}
```
```
func fun1()  {
    a := 2
    c := (*string) (unsafe.Pointer(&a)) //c为a的unsafe.pointer强制转为string指针，*c指的是a的值，a的值是int型的，这里*c = "44"是强制赋了个字符串,所以赋值是失败的，只是没有提示，因为使用了unsafe
    *c = "44"
    fmt.Println(*c) //这时的c是string指针，*c指的应该是一个字符串，但是里面实际上是int数据，所以操作也是失败的
}
func fun2()  {
    a := "654"
    c := (*string) (unsafe.Pointer(&a)) //根据fun1的解释 a 为字符串型，那整个流程就通了，自然*c打出来的是 44，在*c="44"上面打印出来的是肯定是654
    *c = "44"
    fmt.Println(*c)
}
func fun3()  {
    a := 3
    c := *(*string) (unsafe.Pointer(&a))  //这里获取的是a地址里的string值，但a的值是int型，所以操作失败，没有得到值，但c被定义里了string型变量，c = "445"，给string型变量赋string值 自然是正常的
    c = "445"
    fmt.Println(c)
}
```