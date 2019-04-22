package utils

import (
	"fmt"
)


// 使用defer 和 recover来处理错误
func ErrorDemo() (errCount int){

	defer func() {
		if err := recover(); err != nil {
			//说明err有错误
			fmt.Println("发送短信，邮件给管理员...", err)
			errCount++

		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
	return 
}

// 自定义错误
func ErrorDemo2() {
	// 测试的panic 是否会被捕获
	defer func() {
		err := recover() // 程序不会崩溃
		fmt.Println("test defer err=", err) 
	}()

	err := readConfig("config2.ini")
	if err != nil {

		//我们以后的代码，需要去判断到底是什么错误类型，然后给予不同的处理机制
		if err == FILE_NOT_EXIST {
			fmt.Println("错误为FILE_NOT_EXIST , 处理方式为...", err)
			// 如果你希望发生了 FILE_NOT_EXIST， 你认为这个错误是致命，你希望抛出错误，并
			// 退出程序，就使用panic 完成(前提是你没有 defer + recover 处理)
			panic(err)
		} else if err == FILE_IO_ERROR {
			fmt.Println("错误为FILE_IO_ERROR , 处理方式为...", FILE_NOT_EXIST)
		}

		//fmt.Println("readconfig err=", err)
	} else {
		fmt.Println("success")
	}
}


func readConfig(fileName string) (err error) {
	//假设我们要求这个文件名必须 config.ini , 抛出一个错误
	if fileName == "config.ini" {
		return nil // 没错
	} else {
		//返回一个自定义的异常
		//return errors.New("file not exist")
		return FILE_NOT_EXIST
	}
}



