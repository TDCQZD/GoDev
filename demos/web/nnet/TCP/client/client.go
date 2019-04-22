package main

import (
	"strings"
	"os"
	"log"
	"bufio"
	"fmt"
	"net"
)

func Client()  {
	
	conn, err := net.Dial("tcp", "192.168.20.23:8888")
	if err != nil {
		log.Fatal(err)
	}
	
	//客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]
	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = strings.Trim(line,"\r\n")

		if line == "exit" {
			fmt.Println("用户退出客户端")
			break
		}
		//再将line 发送给 服务器
		conent, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据到服务端\n", conent)
	}
}


func main()  {
	Client()
}