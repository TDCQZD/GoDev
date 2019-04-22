package main

import (
	"fmt"
	"log"
	_"io"
	"net"//做网络socket开发时,net包含有网络相关的方法和函数
)

func Server()  {
		// Listen函数创建的服务端
		//tcp : 网络协议
		//192.168.191.1:8888 / :8888 本机IP和端口
		l, err := net.Listen("tcp", "192.168.20.23:8888")
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()//延时关闭listen
		////循环等待客户端访问
		for {
			
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("访问客户端信息： con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())

			go handleConnection(conn)

			// go func(c net.Conn) {
				
			// 	io.Copy(c, c)
			
			// 	c.Close()
			// }(conn)
			
		}
}
//服务端处理从客户端接受的数据
func handleConnection(c net.Conn){	
	defer c.Close()//关闭conn
	
	for {
		
		//1. 等待客户端通过conn发送信息
		//2. 如果客户端没有wrtie[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息\n", c.RemoteAddr().String())
		buf := make([]byte, 1024 )
		n, err := c.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}
		
		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main()  {
	Server()
}