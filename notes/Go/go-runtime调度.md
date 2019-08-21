# GO runtime调度
Golang进程权限调度包runtime三大函数Gosched、Goexit、GOMAXPROCS.
## runtime.Gosched()
runtime.Gosched()，用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其它等待的任务运行，并在下次某个时候从该位置恢复执行。这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑。

## runtime.Goexit()
runtime.Goexit()，调用此函数会立即使当前的goroutine的运行终止（终止协程），而其它的goroutine并不会受此影响。runtime.Goexit在终止当前goroutine前会先执行此goroutine的还未执行的defer语句。
请注意千万别在主函数调用runtime.Goexit，因为会引发panic。

## runtime.GOMAXPROCS()
runtime.GOMAXPROCS()，用来设置可以并行计算的CPU核数最大值，并返回之前的值。

默认此函数的值与CPU逻辑个数相同，即有多少个goroutine并发执行，当然可以设置它，它的取值是１～256。最好在主函数在开始前设置它，因为设置它会停止当前程序的运行。

> 注意：GO默认是使用一个CPU核的，除非设置runtime.GOMAXPROCS那么在多核环境下，什么情况下设置runtime.GOMAXPROCS会比较好的提高速度呢？适合于CPU密集型、并行度比较高的情景。如果是IO密集型，CPU之间的切换也会带来性能的损失。