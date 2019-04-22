package main

import (
	"flag"
	"fmt"
	"os"
)

func CommandDemo() {
	fmt.Println("命令行参数如下：")
	for index, val := range os.Args {
		fmt.Printf("arg[%d] = %v \n", index, val)
	}
}

func FlagDemo() {
	//命令行是：test.exe -u root -p root123 -h localhost -port 8080
	var user, pwd, host, port string
	flag.StringVar(&user, "u", "user", "用户名")
	flag.StringVar(&pwd, "p", "123456", "密码")
	flag.StringVar(&host, "h", "127.0.0.1", "IP地址")
	flag.StringVar(&port, "port", "80", "端口号")

	flag.Parse() //解析注册的flag

	fmt.Println("flag解析命令行参数如下：")
	fmt.Printf("user = %v \n", user)
	fmt.Printf("pwd = %v \n", pwd)
	fmt.Printf("host = %v \n", host)
	fmt.Printf("port = %v \n", port)

}
