# Gin
 Gin 是一个 go 写的 web 框架，具有高性能的优点。
 
 ## 安装
 > 新版本的 Gin 需要 Go 1.6 或者更高版本，并且很快就会要求升级到 Go 1.7.
 ```
 $ go get -u github.com/gin-gonic/gin
 ```
 将 gin 引入到代码中
 ```
 import "github.com/gin-gonic/gin"
 ```
## 使用 Govendor 工具创建项目
1、 govendor(安装)
```
$ go get github.com/kardianos/govendor
```
2、创建项目文件夹并进入文件夹
```
$ mkdir -p $GOPATH/src/github.com/myusername/project && cd "$_"
```
3、初始化项目并添加 gin
```
$ govendor init

$ govendor fetch github.com/gin-gonic/gin@v1.3
```
4、编写代码

示例：复制一个模板到你的项目
```
$ curl https://raw.githubusercontent.com/gin-gonic/gin/master/examples/basic/main.go > main.go
```
5、运行项目
```
$ go run main.go
```
## 使用 jsoniter解析json数据
Gin默认使用encoding/json解析json数据，但您可以通过go build -tags=更改为使用jsoniter。
```
$ go build -tags=jsoniter .
```
## 参考资料
 * 官方地址:https://github.com/gin-gonic/gin
 * https://godoc.org/github.com/gin-gonic/gin
 * https://github.com/skyhee/gin-doc-cn
 * https://www.kancloud.cn/shuangdeyu/gin_book/949411
 * https://github.com/julienschmidt/httprouter
 * https://github.com/json-iterator/go