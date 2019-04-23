# Gopm

> Requirements:Go development environment: >= go1.2

## 安装

```
go get -u github.com/gpmgo/gopm
```

## 命令
###  list命令
> 帮助信息：gopm list -h or gopm help list:
`gopm list`
- 功能：根据当前项目并结合 gopmfile 列出所有依赖和相应的版本。

 选项
* --tags：应用构建 tags。
* --test, -t：列举出 test 文件的依赖。
* --verbose, -v：显示详细信息。
###  gen命令
> gopm gen -h or gopm help gen:
`gopm gen`
-功能：根据当前 Go 项目生成 gopmfile。
-说明：获取依赖并写入 gopmfile。

选项
* --tags：在列举依赖时应用构建 tags。
* --verbose, -v：显示详细信息。
###  get命令
- 功能：根据 gopmfile 拉取远程包及其依赖到本地仓库。

- 说明：如果未传入任何参数，则 gopm 根据在当前目录的项目来进行依赖包的拉取。如果发现 gopmfile 文件，则会应用相关规则。

` gopm get <import path>@[<tag|commit|branch>:<value>]`
- 功能：拉取指定版本的远程包及其依赖到本地仓库。

- 说明：该命令可接受一个或多个参数附带或不带指定版本。

示例：
- 最新版本：gopm get github.com/go-xorm/xorm.
- 固定分支（branch）：gopm get github.com/go-xorm/xorm@branch:master。
- 指定标签（tag）：gopm get github.com/go-xorm/xorm@tag:v0.2.3。
- 某个提交（commit）：gopm get github.com/go-xorm/xorm@commit:6ffffe9。

` gopm get <package name>@[<tag|commit|branch>:<value>]`
- 功能：拉取指定版本的远程包及其依赖到本地仓库；但可使用项目名称代替完整的导入路径。

- 说明：该命令为包导入路径的快捷版。

示例：
- 最新版本：gopm get xorm。
- 固定分支（branch）：gopm get xorm@branch:master。
- 指定标签（tag）：gopm get xorm@tag:v0.2.3。
- 某个提交（commit）：gopm get xorm@commit:6ffffe9。

选项
* --tags：应用构建 tags。
* --download, -d：仅下载当前指定的包。
* --update, -u：检查更新所有包。
* --gopath, -g	：下载所有包至 GOPATH 中。
* --remote, -r：将所有包下载至 gopm 本地仓库。
* --verbose, -v：显示详细信息。
### bin 命令
` gopm bin <import path>@[<tag|commit|branch>:<value>]`
功能：下载指定版本的远程包及其依赖并构建可执行文件。

说明：无需手动处理源代码即可快速构建二进制文件。

示例：
- 最新版本：gopm bin github.com/gpmgo/gopm.
- 固定分支（branch）：gopm bin github.com/gpmgo/gopm@branch:master。
- 指定标签（tag）：gopm bin github.com/gpmgo/gopm@tag:tag:v0.1.0。
- 某个提交（commit）：gopm bin github.com/gpmgo/gopm@commit:23ce93a。

` gopm bin <package name>@[<tag|commit|branch>:<value>]`

- 功能：下载指定版本的远程包及其依赖并构建可执行文件；但可使用项目名称代替完整的导入路径。

- 说明：该命令为包导入路径的快捷版。

示例：
- 最新版本：gopm bin gopm。
- 固定分支（branch）：gopm bin gopm@branch:master。
- 指定标签（tag）：gopm bin gopm@tag:v0.1.0。
- 某个提交（commit）：gopm bin gopm@commit:23ce93a。

选项
* --tags：应用构建 tags。
* -dir, -d：构建可执行文件到指定目录。
* --update, -u：检查更新所有包。
* --verbose, -v：显示详细信息。
### run 命令
`gopm run <go run commands>`

- 功能：根据 gopmfile 链接依赖并执行 go run。
- 示例：gopm run main.go。

选项
* --tags：应用构建 tags。
* --verbose, -v：显示详细信息。
### test 命令
`gopm test <go test commands>`

- 功能：根据 gopmfile 链接依赖并执行 go test。
- 示例：gopm test。

选项

* --tags：应用构建 tags。
* --verbose, -v：显示详细信息。
### build 命令
`gopm build <go build commands>`
- 功能：根据 gopmfile 链接依赖并执行 go build。
- 说明：下载丢失的依赖并链接，然后构建二进制。
- 示例：gopm build。

选项
* --tags：应用构建 tags。
* --update, -u：在构建之前检查包和依赖更新。
* --verbose, -v：显示详细信息。
### install 命令
`gopm install`
- 功能：根据当前项目 gopmfile 链接依赖并执行 go install。
- 说明：下载丢失的依赖然后链接并安装它们。
- 示例：gopm install。

`gopm install <import path>`

- 功能：根据指定导入路径 gopmfile 链接依赖并执行 go install。
- 说明：下载丢失的依赖然后链接并安装它们。
- 示例：gopm install github.com/Unknwon/com。
选项
* --tags：应用构建 tags。
* --verbose, -v：显示详细信息。
### clean 命令
`gopm clean`
- 功能：清除由 gopm 产生的临时文件。
- 示例：gopm clean。
选项
* --verbose, -v：显示详细信息。
### update 命令
`gopm update`
- 功能：检查更新最新的 gopm 资源，包括 gopm 自身。
- 说明：检查资源的版本并下载最新版本。
- 示例：gopm update：
选项
* --verbose, -v：显示详细信息。
## gopmfile
gopmfile 需放在项目根目录下，名称为 .gopmfile。

样例 gopmfile 文件：
```
[target]
path = github.com/gpmgo/gopm

[deps]
github.com/codegangsta/cli = branch:master

[res]
include = conf|etc|public|scripts|templates
```
* target -> path 指示项目名称或导入路径。
* deps 节包含了特殊（非最新）版本的依赖。
* res 在执行 gopm bin 命令时自动打包的资源。
### 如何编写包版本
有五种可能的包版本组合：

- 空白：表示使用最新版本的依赖进行构建。
- `/path/to/my/project`：绝对或者相对的文件路径，例如：d:\projects\xorm。
- `branch:<value>`：固定分支，例如：branch:master。
- `tag:<value>`：指定标签，例如：tag:v0.9.0。
- `commit:<value>`：某个提交，例如：commit:6ffffe9。一般来说只需要 SHA 的前 7 个字母就可以确定一个提交。

## 快速入门
1. 创建 main.go 文件:
```
$ mkdir ~/demo
$ cd ~/demo
$ touch main.go
```
```
package main

import (
	"github.com/astaxie/beego"
)

func main() {
	println("Beego version:", beego.VERSION)
}
```
2. 在项目目录中创建 .gopmfile 文件
```
$ touch .gopmfile
```
```
[target]
path = demo # path 指定项目名称
```
>  如果您的导入路径类似 github.com/gpmgo/gopm 则应该使用 path = github.com/gpmgo/gopm 而不是 path = gopm。

3. 自定义配置
示例：添加beego
```
[target]
path = demo

[deps]
github.com/astaxie/beego = tag:v0.9.0
```
4. 下载依赖
```
$ gopm get
```
5. 构建项目
```
$ gopm build

```
当命令执行成功时，不会输出任何内容。但如果您想要查看详细信息，则可以通过选项 `-v `实现：
```
$ gopm build -v
```
## 参考资料
* https://github.com/gpmgo/gopm
* https://github.com/gpmgo/docs/blob/master/zh-CN/README.md
* https://gopm.io/
