#  Go环境配置

## Go 安装
[GO语言使用之配置篇(二)之Win7搭建go开发环境](https://blog.csdn.net/TDCQZD/article/details/81193572)

[GO语言使用之配置篇(三)之Mac OS X11搭建go开发环境](https://blog.csdn.net/TDCQZD/article/details/81193944)

[GO语言使用之配置篇(四)之Linux搭建go开发环境搭建](https://blog.csdn.net/TDCQZD/article/details/81194146)

## GOPATH 与 GoProject

### GOPATH
Go从1.1版本开始必须设置GOPATH变量，而且不能和Go的安装目录一样，这个目录用来存放Go源码，Go的可运行文件，以及相应的编译之后的包文件。所以这个目录下面有三个子目录：src、bin、pkg

#### GOPATH设置
go 命令依赖一个重要的环境变量：$GOPATH 【Windows系统中环境变量的形式为%GOPATH%】

GOPATH允许多个目录，当有多个目录时，请注意分隔符，多个目录的时候Windows是分号，Linux系统是冒号，当有多个GOPATH时，默认会将go get的内容放在第一个目录下。

以上 $GOPATH 目录约定有三个子目录：

- src 存放源代码（比如：.go .c .h .s等）
- pkg 编译后生成的文件（比如：.a）
- bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）
## GoProject
GOPATH下的src目录就是开发程序的主要目录，所有的源码都是放在这个目录下面，那么一般我们的做法就是一个目录一个项目，例如: $GOPATH/src/mymath 表示mymath这个应用包或者可执行应用，这个根据package是main还是其他来决定，main的话就是可执行应用，其他的话就是应用包，这个会在后续详细介绍package。

当新建应用或者一个代码包时都是在src目录下新建一个文件夹，文件夹名称一般是代码包名称，当然也允许多级目录，例如在src下面新建了目录$GOPATH/src/github.com/astaxie/beedb 那么这个包路径就是"github.com/astaxie/beedb"，包名称是最后一个目录beedb

> 注意：import package的名称和目录名必须保持一致

## go install
上面我们已经建立了自己的应用包，如何进行编译安装呢？有两种方式可以进行安装

1. 只要进入对应的应用包目录，然后执行go install，就可以安装了

2. 在任意的目录执行如下代码go install mymath

## 获取远程包
go语言有一个获取远程包的工具就是`go get`，目前`go get`支持多数开源社区(例如：`github、googlecode、bitbucket、Launchpad`)
```
go get github.com/astaxie/beedb
```
`go get -u` 参数可以自动更新包，而且当`go get`的时候会自动获取该包依赖的其他第三方包

通过这个命令可以获取相应的源码，对应的开源平台采用不同的源码控制工具，例如github采用git、googlecode采用hg，所以要想获取这些源码，必须先安装相应的源码控制工具

通过上面获取的代码在我们本地的源码相应的代码结构如下
```
$GOPATH
  src
   |--github.com
          |-astaxie
              |-beedb
   pkg
    |--相应平台
         |-github.com
               |--astaxie
                    |beedb.a
```
**go get本质上可以理解为首先第一步是通过源码工具clone代码到src下面，然后执行`go install`**

在代码中如何使用远程包，很简单的就是和使用本地包一样，只要在开头import相应的路径就可以
```
import "github.com/astaxie/beedb"
```
## 程序的整体结构
```
bin/
    mathapp
pkg/
    平台名/ 如：darwin_amd64、linux_amd64
         mymath.a
         github.com/
              astaxie/
                   beedb.a
src/
    mathapp
          main.go
      mymath/
          sqrt.go
      github.com/
           astaxie/
                beedb/
                    beedb.go
                    util.go

```
从上面的结构我们可以很清晰的看到，bin目录下面存的是编译之后可执行的文件，pkg下面存放的是应用包，src下面保存的是应用源代码