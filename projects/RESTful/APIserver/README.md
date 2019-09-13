# 启动一个最简单的 RESTful API 服务器
## 说明
**注意：**Windows环境导入包不能使用相对路径
## 本节核心内容
1. 启动一个最简单的 RESTful API 服务器
2. 设置 HTTP Header
3. API 服务器健康检查和状态查询
4. 编译并测试 API

## REST Web 框架选择
要编写一个 RESTful 风格的 API 服务器，首先需要一个 RESTful Web 框架，经过调研选择了 GitHub star 数最多的 Gin。采用轻量级的 Gin 框架，具有如下优点：高性能、扩展性强、稳定性强、相对而言比较简洁。

## 加载路由，并启动 HTTP 服务

main.go 中的 main() 函数是 Go 程序的入口函数，在 main() 函数中主要做一些配置文件解析、程序初始化和路由加载之类的事情，最终调用 http.ListenAndServe() 在指定端口启动一个 HTTP 服务器。本小节是一个简单的 HTTP 服务器，仅初始化一个 Gin 实例，加载路由并启动 HTTP 服务器。

### 编写入口函数
编写 main() 函数，main.go 代码：

```

```


### 加载路由
main() 函数通过调用 router.Load 函数来加载路由（函数路径为 router/router.go，具体函数实现参照 demo01/router/router.go）：

```
```

该代码块定义了一个叫 sd 的分组，在该分组下注册了 /health、/disk、/cpu、/ram HTTP 路径，分别路由到 sd.HealthCheck、sd.DiskCheck、sd.CPUCheck、sd.RAMCheck 函数。sd 分组主要用来检查 API Server 的状态：健康状况、服务器硬盘、CPU 和内存使用量。具体函数实现参照 demo01/handler/sd/check.go。

### 设置 HTTP Header
router.Load 函数通过 g.Use() 来为每一个请求设置 Header，在 router/router.go 文件中设置 Header：

```
```

- gin.Recovery()：在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
- middleware.NoCache：强制浏览器不使用缓存
- middleware.Options：浏览器跨域 OPTIONS 请求设置
- middleware.Secure：一些安全设置
>middleware包的实现见 demo01/router/middleware。


##  API 服务器健康状态自检

有时候 API 进程起来不代表 API 服务器正常，问题：API 进程存在，但是服务器却不能对外提供服务。因此在启动 API 服务器时，如果能够最后做一个自检会更好些。apiserver 中也添加了自检程序，在启动 HTTP 端口前 go 一个 pingServer 协程，启动 HTTP 端口后，该协程不断地 ping /sd/health 路径，如果失败次数超过一定次数，则终止 HTTP 服务器进程。通过自检可以最大程度地保证启动后的 API 服务器处于健康状态。自检部分代码位于 main.go 中：

```
```
在 pingServer() 函数中，http.Get 向 http://127.0.0.1:8080/sd/health 发送 HTTP GET 请求，如果函数正确执行并且返回的 HTTP StatusCode 为 200，则说明 API 服务器可用，pingServer 函数输出部署成功提示；如果超过指定次数，pingServer 直接终止 API Server 进程，如下图所示。


##  编译源码
将vendor文件夹中的包拷贝到相应位置。
做检查然后编译。

```
$ gofmt -w .
$ go tool vet .
$ go build -v .
```

>建议每次编译前对 Go 源码进行格式化和代码静态检查，以发现潜在的 Bug 或可疑的构造。

##  cURL 工具测试 API

###  cURL 工具简介

### 启动 API Server

### 发送 HTTP GET 请求