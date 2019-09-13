# chat_beego
通过两种形式展示了如何实现一个在线聊天室应用：

- 使用长轮询模式。
- 使用 WebSocket 模式。
以上两种模式均默认将数据存储在内存中，因此每次启动都会被重置。但您也可以通过修改 conf/app.conf 中的设置来启用数据库

## 新建项目
```
bee new chat_beego
```
## 测试
```
bee run 
```

## 问题
```
Handler crashed with error can't find templatefile in the path:views/appcontroller/join.tpl
```
## 参考资料
* https://github.com/beego/samples/tree/master/WebIM