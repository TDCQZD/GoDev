package main

import (
	"fmt"
	// "go_code/grammar/function/package/code/details"
	de "go_code/grammar/function/package/code/details"
)
/*包的注意事项和细节说明
 	1) 在给一个文件打包时，该包对应一个文件夹，比如这里的 utils 文件夹对应的包名就是utils, 
        文件的包名通常和文件所在的文件夹名一致，一般为小写字母
    2) 当一个文件要使用其它包函数或变量时，需要先引入对应的包

        引入方式1：import  "包名"
        引入方式2：
        import  (
            "包名1"
            "包名2"
        ) 
        package 指令在 文件第一行，然后是 import 指令。
        在import 包时，路径从 $GOPATH 的  src 下开始，不用带src , 编译器会自动从src下开始引入
    3) 为了让其它包的文件，可以访问到本包的函数，则该函数名的首字母需要大写，类似其它语言的public ,这样才能跨包访问。比如 utils.go 的, 全局变量也遵守该规则。
    
    4) 在访问其它包函数，变量时，其语法是 包名.函数名， 比如这里的 main.go文件中
    
    5) 如果包名较长，Go支持给包取别名， 注意细节：取别名后，原来的包名就不能使用
    
    6) 在同一包下，不能有相同的函数名（也不能有相同的全局变量名），否则报重复定义
    
	7) 如果你要编译成一个可执行程序文件，就需要将这个包声明为 main , 即 package main .
	这个就是一个语法规范，如果你是写一个库 ，包名可以自定义
	
	如使用命令：go build -o bin/package.exe go_code/grammar/function/package/code/main
	在E:\GoProject\bin 目录下生成package.exe，可提供其他人调用
*/
func main()  {
	fmt.Println("主函数")
	// details.Yunsun()

	// details.Bm()
	de.Bm()
}