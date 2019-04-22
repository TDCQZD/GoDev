package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

/* Go 原子计数器
Go里面的管理协程状态的主要机制就是通道通讯。
使用sync/atomic包来实现由多个goroutine访问的原子计数器。
*/

func main() {

	var ops uint64 = 0 // 一个无符号整数表示计数器(正数)

	for i := 0; i < 50; i++ { // 启动50个goroutine，每个增量计数器大约是1毫秒
		go func() {

			for {
				// 为了原子地递增计数器，这里使用AddUint64()函数，在ops计数器的内存地址上使用＆语法。
				atomic.AddUint64(&ops, 1)
				// 在增量之间等待一秒，允许一些操作累积。
				time.Sleep(time.Millisecond)

				// 允许其他的协程来处理
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	// 为了安全地使用计数器，同时它仍然被其他goroutine更新，通过LoadUint64提取一个当前值的副本到opsFinal
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
