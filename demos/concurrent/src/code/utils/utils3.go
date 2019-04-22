package utils

import (
	"runtime"
	"fmt"
)

func CpuDemo()  {
	//获取当前系统cpu数量
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多个cpu，设置num-1的cpu运行程序
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}