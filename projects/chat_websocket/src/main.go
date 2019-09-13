package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket" // websocket连接包
)

// Message 定义消息类型
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// 已连接的客户端集合，map类型
var clients = make(map[*websocket.Conn]bool)

// 用来广播消息的通道
var broadcast = make(chan Message)

// Upgrader 指定用于将HTTP连接升级到WebSocket连接的连接的参数。
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// 创建一个简单的文件服务器, 用来托管静态文件，或者叫前端页面
	fs := http.FileServer(http.Dir("../static"))

	// 首页的url路由
	http.Handle("/", fs)
	// 配置websocket路由
	http.HandleFunc("/ws", handleConnections)

	// 监听从客户端发送到服务端的消息
	go handleMessages()

	// Http 连接异常处理
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// handleConnections 处理Clients连接
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// 建立websocket连接
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	// 将新的连接添加到clients里面
	clients[ws] = true

	for {
		var msg Message
		// 读取客户端发送来的消息，并反序列化json，然后映射为msg。
		err := ws.ReadJSON(&msg)

		// websocket 连接异常处理
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// 将接收到的消息广播出去,发送给每一个客户端
		broadcast <- msg
	}
}

// 监听从客户端发送到服务端的消息
func handleMessages() {
	for {
		// 从broadcast通道获取下一条要广播的信息
		msg := <-broadcast
		// 将msg发送给所有连接到服务端的客户端
		for client := range clients {
			clientMsg := client
			err := clientMsg.WriteJSON(msg)
			// websocket 连接异常处理
			if err != nil {
				log.Printf("error: %v", err)
				clientMsg.Close()
				delete(clients, clientMsg)
			}
		}
	}
}
