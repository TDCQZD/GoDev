package utils

import (
	"fmt"
)
// 要求统计1-200000 的数字中，哪些是素数？这个问题在本章开篇就提出了，
// 现在我们有goroutine和channel的知识后，就可以完成了
/* 分析：
1、开启initChan通道存放数据
2、开启prime通道存放素数
3、开启exitChan 存放协程结束标志
4、开启4个协程处理1+~+n
   1)协程方法处理
   2）判读素数
   3）当前协程取值结束，存放true到exitChan
5、主线程for exitchan ，协程全部执行结束，下一步，否则阻塞并且关闭prime
6、循环打印exitchan
*/
func writeDataToChan(initChan chan int)  {
	for i := 1; i <=200000; i++ {
		initChan<- i
	}
	close(initChan)
	
}
func writePrimeTochan(initChan chan int, primeChan chan int, exitChan chan bool)  {
	
	for {
		val, ok := <-initChan
		if !ok {
			break
		}
		if numberIsPrime1(val) {
			primeChan<- val
		}
	}
	exitChan<- true
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
}

func numberIsPrime1(num int) bool{
	for i := 2; i < num; i++ {
		if (num % i == 0) {
			return false;
		}
	}
	return true;

}

func PrimeGoroutine() bool {
	initChan := make(chan int,200000)
	primeChan := make(chan int,200000)
	exitChan := make(chan bool,8)

	go writeDataToChan(initChan)
	for i := 0; i < 8; i++ {
		go writePrimeTochan(initChan, primeChan, exitChan)//需要四个
	}

	go func ()  {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		_, ok := <-primeChan
		if !ok{
			break
		}
		
		// fmt.Printf("素数=%d\n", res)
	}
	return true
}