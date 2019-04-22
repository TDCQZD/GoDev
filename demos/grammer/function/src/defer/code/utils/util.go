package utils

import (
	"fmt"
)
//defer 实例
func DeferDemo(n1 ,n2 int)  int {
	
	defer fmt.Println("defer1 test",n1)
	defer fmt.Println("defer2 test" ,n2)
	res := n1 +n2
	fmt.Println("utils res =",res)
	return res
}
//defer 实例 扩展1
func DeferDemo2(n1 ,n2 int)  int {
	
	defer fmt.Println("defer1 test",n1)
	defer fmt.Println("defer2 test" ,n2)
	n1++
	n2++
	res := n1 +n2
	fmt.Println("utils res =",res)
	return res
}

//defer 实例 扩展2
func DeferDemo3(n1 ,n2 int)  int {
	defer fmt.Println("ok1 n1=", n1) // sum defer栈  （6） ok1 n1 = 10
	defer fmt.Println("ok2 n2=", n2) // sum defer栈   （5） ok2 n2 = 20
	n1++ // n1 = 11
	n2++ // n2 = 21
	res := n1 + n2 // res = 32
	fmt.Println("ok3 res=", res) // (1) ok3 res = 32
	test(9) 
	return res
}
func test(num int) {
	num++ // num = 10
	defer fmt.Println(num) // test的defer 栈 (4) 10
	num++ //num =11
	defer fmt.Println(num) //test的defer 栈 （3） 11
	fmt.Println("test num=", num) //(2) test num = 11
}

//defer 返回值 注意
func DeferDemo4(n1 int, n2 int)  (res int) {

		defer func() {
			res = res + 2
			fmt.Println("defer res=", res) //32! 37!
		}()
		res = n1 + n2 // 30
		fmt.Println("sum res=", res)//30
		//(1) 将这个 res 这值先付给 返回的列表值
		//[(1)如果返回的形参列表有名字，则等价于赋值  (2) 如果返回列表没有名字，则相当于将返回值，保存到一个临时变量] 
		//(2) 再defer 
		return res + 5 // res=30 , 35 // res = res + 5
	

}

func DeferDemo5(n1 int, n2 int)   int {

	res :=0
	defer func() {
		res = res + 2
		fmt.Println("defer res=", res) //32! 37!
	}()
	res = n1 + n2 // 30
	fmt.Println("sum res=", res)//30
	//(1) 将这个 res 这值先付给 返回的列表值
	//[(1)如果返回的形参列表有名字，则等价于赋值  (2) 如果返回列表没有名字，则相当于将返回值，保存到一个临时变量] 
	//(2) 再defer 
	return res + 5 // res=30 , 35 // res = res + 5


}
/*
// defer的最佳实践——关闭文件资源
func DeferDemo6() {
	// 关闭文件资源
	file =openfile(文件名)
	defer file.close()
	// 其它代码
}

// defer的最佳实践——关闭数据库资源
func DeferDemo7() {
	// 关闭数据库资源
	connect =openDataBase()
	defer connect.close()
	// 其它代码
}
*/

