# Govendor
> 注意：使用govendor，必须保证你的工程项目放在GOPATH/src目录下
## 特性
- 可以采用govendor add/update复制现有的依赖从$GOPATH

- 如果要忽视vendor/*/，可采用govendor 同步恢复依赖

- 可直接通过govendor fetch控制器添加新的依赖或者更新现有依赖

- 可采用govendor migrate 实现系统间迁移

- 支持Linux, OS X, Windows，甚至现有所有操作系统

- 支持Git，Hg，SVN，BZR（必须指定一个路径）

## 安装
```
go get -u github.com/kardianos/govendor
```
## Govendor 命令
|命令|功能|
| :------| :------| 
|init|	初始化 vendor 目录,创建 "vendor" 文件夹和 "vendor.json"文件|
|list|	列出所有的依赖包|
|add	|从 $GOPATH 添加包,添加包到 vendor 目录，如govendor add +external 添加所有外部包|
|add PKG_PATH	|添加指定的依赖包到 vendor 目录|
|update	|从 $GOPATH 更新包,从 $GOPATH 更新依赖包到 vendor 目录|
|remove	|从 vendor 管理中删除依赖|
|status	|列出所有缺失、过期和修改过的包|
|fetch	|从远端仓库中增加或更新 vendor 文件中依赖的包|
|sync	|使用vendor.json文件中的修订将包从远程存储库中提取到vendor文件夹中|
|get	|类似 go get 目录，拉取依赖包到 vendor 目录|
|license  |列出已发现的给定状态或导入路径的许可证.|
|shell|    对于大项目，运行一个 "shell" 使多个子命令更有效率.|


```
# 初始化项目
govendor init


# 添加 GOPATH 中已存在的文件到 vendor.
govendor add +external
govendor add +e 

# 列出项目依赖列表.
govendor list

# 查看一个包在哪些地方被使用
govendor list -v fmt

# 指定要获取的特定版本或修订版本
govendor fetch golang.org/x/net/context@a4bbce9fcae005b22ae5443f6af064d80a6f5a55
govendor fetch golang.org/x/net/context@v1   # Get latest v1.*.* tag or branch.
govendor fetch golang.org/x/net/context@=v1  # Get the tag or branch named "v1".

# 将一个包更新到最新，并指定上一个版本的约束
govendor fetch golang.org/x/net/context

# 仅仅格式化您自己的仓库
govendor fmt +local

# 仅构建你仓库中的任何内容
govendor install +local

# 仅仅测试你自己的仓库
govendor test +local

```


## 参考资料
* https://github.com/kardianos/govendor
* https://blog.csdn.net/cyberspecter/article/details/83345760