# chat_gin

## 项目描述
golang实现的聊天室demo

三种形式：
- refresh
- longpolling（轮询）
- websocket


## 创建项目
```
$ mkdir chat_gin && cd chat_gin
$ govendor init
$ govendor fetch github.com/gin-gonic/gin@v1.3
```
## 测试
```
go run main.go
```
**效果图**
![](./imgs/chat.png)

## 问题
- Refresh、LongPolling 显示有问题

## 参考资料
* https://github.com/confucianzuoyuan/go-tutorials/tree/master/%E4%BD%BF%E7%94%A8Golang%E5%92%8CWebSocket%E5%AE%9E%E7%8E%B0%E5%AE%9E%E6%97%B6%E8%81%8A%E5%A4%A9%E5%AE%A4/upgrade-chat-room