package main

import (
	"fmt"
	"go_code/grammar/function/closure/code/utils"
)

func main(){
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))

	fmt.Println("--------闭包方式------")
	f1 := utils.MakeSuffix1(".avi")
	fmt.Println(f1("霹雳惊涛"))
	fmt.Println(f1("霹雳天启"))

	fmt.Println("--------传统方式------")
	res := utils.MakeSuffix2(".avi","霹雳九轮异谱")
	fmt.Println(res)
	res = utils.MakeSuffix2(".avi","霹雳天罪")
	fmt.Println(res)
}


/*分析：
1、AddUpper会返回一个匿名函数
func(x int) int {
		n = n + x
		return n
	}
2、该匿名函数引用到其函数外部一个变n 
3、匿名函数和它引用到的外部变量n,形成一个整体，构成一个闭包
4、 f(1) , f(2) ...即在反复调用闭包
5、 当我们使用闭包时，该匿名函数引用到外部n， 只会被初始化一次，以后调用时，则使用同一n.
6、你可以这样理解，把闭包看成一个类，匿名函数就相当于方法，而引用到的外部变量就是字段。
7、 搞清楚闭包的关键点，就是分析出，该匿名函数引用到外部变量是哪些。
*/
func AddUpper() func (int)  int{
	var n = 10
	return func (x int ) int  {
		n = n + x
		return n
	}
}