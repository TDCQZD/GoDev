# Glide
glide是Go的包管理工具。支持语义化版本,支持Git、Svn等，支持Go工具链，支持vendor目录，支持从Godep、GB、GPM、Gom倒入，支持私有的Repos和Forks。

glide主要功能：

- 持久化依赖列表至配置文件中，包括依赖版本（支持范围限定）以及私人仓库等
- 持久化关系树至 lock 文件中（类似于 yarn 和 cargo），以重复拉取相同版本依赖
- 兼容 go get 所支持的版本控制系统：Git, Bzr, HG, and SVN
- 支持 GO15VENDOREXPERIMENT 特性，使得不同项目可以依赖相同项目的不同版本
- 可以导入其他工具配置，例如： Godep, GPM, Gom, and GB
## 安装
``` window
$ go get github.com/Masterminds/glide
$ go install github.com/Masterminds/glide
```
``` Linux
$ curl https://glide.sh/get | sh
$ sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
$ sudo apt-get install glide
```

``` Mac OS X
$ curl https://glide.sh/get | sh
$ brew install glide
```
### 验证
```
$ glide
NAME:
   glide - Vendor Package Management for your Go projects.

   Each project should have a 'glide.yaml' file in the project directory. Files
   look something like this:

       package: github.com/Masterminds/glide
       imports:
       - package: github.com/Masterminds/cookoo
         version: 1.1.0
       - package: github.com/kylelemons/go-gypsy
         subpackages:
         - yaml

   For more details on the 'glide.yaml' files see the documentation at
   https://glide.sh/docs/glide.yaml


USAGE:
   glide [global options] command [command options] [arguments...]

VERSION:
   0.13.2-dev

COMMANDS:
     create, init       Initialize a new project, creating a glide.yaml file
     config-wizard, cw  Wizard that makes optional suggestions to improve config in a glide.yaml file.
     get                Install one or more packages into `vendor/` and add dependency to glide.yaml.
     remove, rm         Remove a package from the glide.yaml file, and regenerate the lock file.
     import             Import files from other dependency management systems.
     name               Print the name of this project.
     novendor, nv       List all non-vendor paths in a directory.
     rebuild            Rebuild ('go build') the dependencies
     install, i         Install a project's dependencies
     update, up         Update a project's dependencies
     tree               (Deprecated) Tree prints the dependencies of this project as a tree.
     list               List prints all dependencies that the present code references.
     info               Info prints information about this project
     cache-clear, cc    Clears the Glide cache.
     about              Learn about Glide
     mirror             Manage mirrors
     help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --yaml value, -y value  Set a YAML configuration file. (default: "glide.yaml")
   --quiet, -q             Quiet (no info or debug messages)
   --debug                 Print debug verbose informational messages
   --home value            The location of Glide files (default: "C:\\Users\\ZhangDai\\.glide") [%GLIDE_HOME%]
   --tmp value             The temp directory to use. Defaults to systems temp [%GLIDE_TMP%]
   --no-color              Turn off colored output for log messages
   --help, -h              show help
   --version, -v           print the version
```
### windows 安装 bug
**win7 glide install**
```
 Unable to export dependencies to vendor directory: Error moving files: exit status 1. output:
```
1. 修改 github.com/Masterminds/glide/path/winbug.go 76 行
```
cmd := exec.Command("cmd.exe", "/c", "xcopy /s/y", o, n+"\\")
```
2. 重新编译
```
go get -u github.com/Masterminds/glide
```
3. 再次安装
```
glide install
```
> 详情参考[glide/issues/873](https://github.com/Masterminds/glide/issues/873)


## glide命令
```
# 初始化glide配置
glide create
glide init

# 添加新的包
glide get [package name]

# 根据glide.yaml更新包
glide update
glide up

# 根据glide.yaml安装包
glide install

# 返回当前项目的名称
glide name

# 列出当前项目已安装的包
glide list

# 替换包的镜像
glide mirror set [original] [replacement]
glide mirror set [original] [replacement] --vcs [type]
glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto --vcs git

# 移除包的镜像
glide mirror remove [original]

# 获取包的镜像列表
glide mirror list
```
> glide mirror特别适用于不能访问一些站点，导致很Golang的依赖包不能通过go get下载的情况。可以通过配置将墙了的版本库 URL 映射到没被墙的 URL，甚至也可以映射到本地版本库。

### glide init 
glide初始化
```
glide init
```
在初始化过程中， glide 会询问一些问题。 glide.yaml记载了依赖包的列表及其更新规则，每次执行 glide up 时，都会按照指定的规则（如只下载补丁(patch)不下载升级(minor)）下载新版。

### glide install
安装依赖
```
glide install
```
> 注意:安装完成之后，会生成glide.lock文件，锁定安装包的版本。

### glide up
升级版本
```
glide up
```
glide up 会按照语义化版本规则更新依赖包代码，开发过程中如果需要使用新版代码，可以执行这个命令： 修改一下glide.yaml中的一个Package.
```
- package: github.com/astaxie/beego
  version: 1.8.3
```
执行glide up。
### glide get
添加并下载依赖

除了自动从代码中解析 import 外，glide 还可以通过 glide get 直接下载代码中没有的依赖，与 go get 的用法基本一致：
```
$ glide get github.com/orcaman/concurrent-map 
```

### glide mirror
使用镜像

Golang的依赖包不能通过go get下载：
```
[WARN]    Unable to checkout golang.org/x/crypto
[ERROR]    Update failed for golang.org/x/crypto: Cannot detect VCS
[ERROR]    Failed to do initial checkout of config: Cannot detect VCS
```
可以通过配置将墙了的版本库 URL 映射到没被墙的 URL，甚至也可以映射到本地版本库。 将golang.org映射到github: 修改glide.yaml加入
```
- package: golang.org/x/crypto
```
执行如下命令：
```
$ glide mirror set golang.org/x/crypto github.com/golang/crypto
$ glide up
```

## glide.yaml解析
```
package: github.com/Masterminds/glide
homepage: https://masterminds.github.io/glide
license: MIT
owners:
- name: Matt Butcher
  email: technosophos@gmail.com
  homepage: http://technosophos.com
- name: Matt Farina
  email: matt@mattfarina.com
  homepage: https://www.mattfarina.com
ignore:
- appengine
excludeDirs:
- node_modules
import:
- package: gopkg.in/yaml.v2
- package: github.com/Masterminds/vcs
  version: ^1.2.0
  repo:    git@github.com:Masterminds/vcs
  vcs:     git
- package: github.com/codegangsta/cli
  version: f89effe81c1ece9c5b0fda359ebd9cf65f169a51
- package: github.com/Masterminds/semver
  version: ^1.0.0
# 测试导入包
testImport:
- package: github.com/arschles/assert
```
**glide.yaml中的这些元素的解释如下：**

- package：顶部的 package 是它所在GOPATH的位置，glide 将从该位置下开始导包。

- homepage：该项目的详情页面。

- license：许可证标识，可以是SPDX license字符串或文件路径。

- owners：项目的所有者信息，便于接受漏洞信息。

- ignore：忽略导入的包，注意是包而不是目录。

- excludeDirs：排除扫描依赖的目录。

- import：import 的包列表：

- package：导入包的名称，必填。软件包名称遵循go工具所用的相同模式。这意味着：1、映射到VCS远程位置的软件包名称以.git，.bzr，.hg或.svn结尾。 例如，example.com/foo/pkg.git/subpkg。2、GitHub, BitBucket, Launchpad, IBM Bluemix Services, and Go on Google Source是特殊情况，不需要 VCS 扩展。

- version：可以为semantic version, semantic version range, branch, tag 或者 commit id。

- repo：如果包名称不是repo位置或这是一个私人存储库，它可以去这里。 该软件包将从repo签出并放在软件包名称指定的位置。 这允许使用fork。

- vcs：要使用的VCS，如git，hg，bzr或svn。仅当无法从名称中检测到类型时才需要。例如，以.git或GitHub结尾的仓库可以被检测为Git。 对于Bitbucket的repo，我们可以联系API来发现类型。

- subpackages：在存储库中使用的包的记录。这不包括存储库中的所有包，而是包括正在使用的包。

- os：用于过滤的操作系统的列表。如果设置它将比较当前运行时操作系统与指定的操作系统，并且只有获取匹配的依赖。如果未设置过滤，则跳过。这些名称与构建标志和GOOS环境变量中使用的名称相同。

- arch：用于过滤的体系结构列表。如果设置它将比较当前运行时架构与指定的架构，并且只有在匹配时获取依赖关系。如果未设置过滤，则跳过。名称与构建标志和GOARCH环境变量中使用的名称相同。

- testImport：在导入中未列出的测试中使用的软件包列表。每个包具有与导入下列出的相同的详细信息

## glide版本号指定规则如下：
```
=: equal (aliased to no operator)
!=: not equal
>: greater than
<: less than
>=: greater than or equal to
<=: less than or equal to

1.2 - 1.4.5 which is equivalent to >= 1.2, <= 1.4.5
2.3.4 - 4.5 which is equivalent to >= 2.3.4, <= 4.5
1.2.x is equivalent to >= 1.2.0, < 1.3.0

>= 1.2.x is equivalent to >= 1.2.0
<= 2.x is equivalent to < 3
* is equivalent to >= 0.0.0

~1.2.3 is equivalent to >= 1.2.3, < 1.3.0
~1 is equivalent to >= 1, < 2
~2.3 is equivalent to >= 2.3, < 2.4
~1.2.x is equivalent to >= 1.2.0, < 1.3.0
~1.x is equivalent to >= 1, < 2

^1.2.3 is equivalent to >= 1.2.3, < 2.0.0
^1.2.x is equivalent to >= 1.2.0, < 2.0.0
^2.3 is equivalent to >= 2.3, < 3
^2.x is equivalent to >= 2.0.0, < 3
```


## 参考资料
* https://github.com/Masterminds/glide
* https://glide.sh/
* https://glide.readthedocs.io/en/latest/