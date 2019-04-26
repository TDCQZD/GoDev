# dep
Go语言的官方版依赖管理工具
## 安装
```
$ go get -u github.com/golang/dep/cmd/dep
```
## 命令
###  dep init
初始化操作。

备份当前的vender，创建vender目录并下载项目中的所有依赖包，生成Gopkg.lock和Gopkg.toml
以下是两个文件的作用说明，
- Gopkg.toml是清单文件，
- Gopkg.lock是校验描述文件。
尽量不要修改，避免造成两个文件不同步的错误。
### dep ensure
> ` -v` 显示详细过程

#### 添加新依赖项
```
$ dep ensure -add github.com/pkg/errors github.com/foo/bar   # 添加新依赖项
$ dep ensure -add github.com/gin-gonic/gin                   # 获取最新TAG的依赖包
$ dep ensure -add github.com/gin-gonic/gin@master            # 获取指定分支的依赖包
$ dep ensure -add github.com/gin-gonic/gin@V1.2              #获取指定TAG的依赖包
$ dep ensure -add github.com/gin-gonic/gin@^V1.2             # 获取指定TAG以上的依赖包

```
#### 更新依赖项
```
$ dep ensure -update github.com/foo/bar #更新依赖项
$ dep ensure -update # 更新所有依赖项（尽管通常不推荐）
```

### dep status
用来查看项目依赖的详细信息和状态。当缺少依赖项时，`dep status`会告诉您哪些软件包缺失


## 本地缓存
当然dep不会每次都去下载，其工作原理和Mevan一样，会优先在本地仓库搜索，本地仓库未找到即在网络上下载，并添加到本地仓库。

## 快速入门
1. 初始化项目
```
$ dep init
$ ls
Gopkg.toml Gopkg.lock vendor/
```
2. 添加依赖
```
$ dep ensure -add github.com/foo/bar github.com/baz/quux
```
 
## 参考资料
* https://github.com/golang/dep
* https://godoc.org/github.com/golang/dep
* https://github.com/golang/dep/blob/master/docs/Gopkg.toml.md
* https://golang.github.io/dep/
* [godep] :https://github.com/tools/godep

