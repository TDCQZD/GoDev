# pprof
Profile是一个调用栈踪迹的集合，显示导致特定事件（如内存分配）的实例的调用栈序列。包可以创建并维护它们自己的profile；它一般用于跟踪必须被显式关闭的资源，例如文件或网络连接。

一个Profile的方法可被多个Go程同时调用。

pprof 由两部分组成：
- runtime/pprof 每个 Go 程序都内置的包
- go tool pprof 用于解析 profile 文件

pprof 支持好几种类型的分析
- CPU 分析
- 内存分析
- 阻塞分析
- 锁竞争分析

通过它的HTTP服务端提供pprof可视化工具期望格式的运行时剖面文件数据服务。
## 使用 pprof
解析使用 go pprof 子命令：
```
go tool pprof /path/to/your/profile
```
注意 : 如果你已经使用 Go 一段时间了，你可能会被告知pprof有两个参数。从 Go 1.9 开始，profile 文件包含展示 profile 所需的所有信息。你不再需要生成 profile 的二进制文件了。