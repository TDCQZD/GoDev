package utils

import (
	"fmt"
)

/*
6) 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。
*/
func Point(n *int)  {
	 *n++
	 fmt.Printf("变量n本身的存储空间地址%p \nn存储空间的值是%v \nn存储空间的值指向的数据是%v \n",
	&n,n,*n)
}
/*
8) 在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。
通过该变量可以对函数调用
*/

func Variable(n ,m int) int {
	
	return n + m
}
/*
9) 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
*/
func FormParameter(funParameter func(int, int) int,n int ,m int ) int  {
	
	return funParameter(n,m)
}

// Go支持自定义数据类型，结构体中常用
func NameDataType()  {
	var n = 888
	type D int//自定义变量类型D
	var n1 D = 666
	fmt.Printf("默认int变量赋值：%d \n自定义int变量m赋值结果：%d\n",n,n1)
	  n1 = D(n)
	  fmt.Printf("默认int变量强转为自定义变量：%d\n",n1)
	  fmt.Println("---------------------")
}

// 11) 支持对函数返回值命名
func NameFuncReturn(n ,m int) (sum ,sub int) {
	sum = n + m
	sub = n - m
	return
}
// 13) Go支持可变参数

//支持0到正无穷个参数
func MoreArgs(args... int ) (sum int) {
	for i := range args {
		sum +=args[i]
	}
	return
}
//支持1到正无穷个参数
func MoreArgs1(n int,args... int)  (sum int) {
	sum =n
	for i := 0; i < len(args); i++ {
		sum +=args[i]
	}
	return 
}
//编写一个函数可以交换两个变量的值
func SwapPoint(n ,m *int)  {
	t := *n
	*n = *m
	*m = t
}