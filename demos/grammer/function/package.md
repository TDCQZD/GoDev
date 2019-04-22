## package
go的每一个文件都是属于一个包的，也就是说go是以包的形式来管理文件和项目目录结构的.





### 包基本语法
```
package 包名
引入包的基本语法
import "包的路径"
```
### 包的三大作用
1.  区分相同名字的函数、变量等标识符 
2. 当程序文件很多时,可以很好的管理项目 
3. 控制函数、变量等访问范围，即作用域
## import
Go使用import用来导入包文件
```
import(
    "fmt"
)
```
> 在import 包时，路径从 $GOPATH 的 src 下开始(不用带src) , 编译器会自动从src下开始引入

上面这个fmt是Go语言的标准库，其实是去GOROOT环境变量指定目录下去加载该模块，当然Go的import还支持如下两种方式来加载自己写的模块：

1. 相对路径
```
import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
```
2. 绝对路径
```
import “shorturl/model” //加载gopath/src/shorturl/model模块
```
### 特殊的import

1. 点操作
```
 import(
     . "fmt"
 )
```
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的`fmt.Println("hello world")`可以省略的写成`Println("hello world")`

2. 别名操作
```
 import(
     f "fmt"
 )
```
别名操作的话调用包函数时前缀变成了我们的前缀，即`f.Println("hello world")`

3. _操作
```
	import (
	    "database/sql"
	    _ "github.com/ziutek/mymysql/godrv"
	)
```
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。