## bee 工具

## bee 工具的安装
```
go get github.com/beego/bee
```
> 安装完之后，bee 可执行文件默认存放在 $GOPATH/bin 里面，所以您需要把 $GOPATH/bin 添加到您的环境变量中，才可以进行下一步。
## bee 工具命令详解
### new 命令
new 命令是新建一个 Web 项目，在命令行下执行 `bee new <项目名>` 就可以创建一个新的项目。但是注意该命令必须在 $GOPATH/src 下执行。最后会在 $GOPATH/src 相应目录.
```
bee new beegoNewDemo
```
### api 命令
 api 命令就是用来创建 API 应用的，
```
bee api beegoApiDemo
```
从目录可以看到和 Web 项目相比，少了 static 和 views 目录，多了一个 test 模块，用来做单元测试的。

同时，该命令还支持一些自定义参数自动连接数据库创建相关 model 和 controller:
```
bee api [appname] [-tables=""] [-driver=mysql] [-conn="root:<password>@tcp(127.0.0.1:3306)/test"]
```
如果 conn 参数为空则创建一个示例项目，否则将基于链接信息链接数据库创建项目。
### run 命令
bee run 命令是监控 beego 的项目，通过 fsnotify监控文件系统。但是注意该命令必须在 $GOPATH/src/appname 下执行。
```
bee run 
```
> 测试：http://localhost:8080/
### pack 命令
pack 目录用来发布应用的时候打包，会把项目打包成 zip 包，这样我们部署的时候直接把打包之后的项目上传，解压就可以部署了
```
bee pack
```
### bale 命令
这个命令目前仅限内部使用，具体实现方案未完善，主要用来压缩所有的静态文件变成一个变量申明文件，全部编译到二进制文件里面，用户发布的时候携带静态文件，包括 js、css、img 和 views。最后在启动运行时进行非覆盖式的自解压。
### version 命令
这个命令是动态获取 bee、beego 和 Go 的版本，这样一旦用户出现错误，可以通过该命令来查看当前的版本
```
 bee version
```

### generate 命令
这个命令是用来自动化的生成代码的，包含了从数据库一键生成 model，还包含了 scaffold 的，通过这个命令，让大家开发代码不再慢.
```
bee generate scaffold [scaffoldname] [-fields=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
```

### migrate 命令
这个命令是应用的数据库迁移命令，主要是用来每次应用升级，降级的SQL管理。
```
bee migrate [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    run all outstanding migrations
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

bee migrate rollback [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    rollback the last migration operation
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

bee migrate reset [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    rollback all migrations
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

bee migrate refresh [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    rollback all migrations and run them all again
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
```

###  dockerize 命令
这个命令可以通过生成Dockerfile文件来实现docker化你的应用。
``` 生成一个以1.6.4版本Go环境为基础镜像的Dockerfile,并暴露9000端口:
bee dockerize -image="library/golang:1.6.4" -expose=9000
```
## bee 工具配置文件
 bee 工具的源码目录下有一个 bee.json 文件，这个文件是针对 bee 工具的一些行为进行配置。该功能还未完全开发完成，不过其中的一些选项已经可以使用：

- "version": 0：配置文件版本，用于对比是否发生不兼容的配置格式版本。
- "go_install": false：如果您的包均使用完整的导入路径（例如：github.com/user/repo/subpkg）,则可以启用该选项来进行 go install 操作，加快构建操作。
- "watch_ext": []：用于监控其它类型的文件（默认只监控后缀为 .go 的文件）。
- "dir_structure":{}：如果您的目录名与默认的 MVC 架构的不同，则可以使用该选项进行修改。
- "cmd_args": []：如果您需要在每次启动时加入启动参数，则可以使用该选项。
- "envs": []：如果您需要在每次启动时设置临时环境变量参数，则可以使用该选项。