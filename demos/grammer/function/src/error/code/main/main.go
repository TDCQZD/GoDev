package main

import (
	"fmt"
	_"time"
	"go_code/grammar/function/error/code/utils"
)

func main()  {
	/*
	errCount := 0
	for {
		if errCount < 10 {
			errCount = utils.ErrorDemo() 
		}
		time.Sleep(time.Second)
		fmt.Println("下面的代码和逻辑...")
	}
	*/
	utils.ErrorDemo2() 
	fmt.Println("main() 继续执行...")
}