package utils

import (
	"fmt"
)

//算数操作符
func ArithmeticOperator()  {
	// / 的使用 

	var n1 int = 10
	var n2 int = 4
	res := n1 / n2 // ==?
	fmt.Printf("res=%v\n", res) // 2

	 var n3 float64
	 n3 = 10 / 4.0 // 2.5

	var n4 float64 = 10.0
	var n5 float64
	n5 = n4 / 4

	fmt.Printf("n5=%f\n", n5)
	fmt.Printf("n3=%f\n", n3)

	// % 的使用
	//取模的一个公式  a % b = a - a / b * b
	fmt.Printf("10 模 3= %v\n", (10 % 3)) // 1
	//-10 % 3 = -10 - (-10) / 3 * 3 =  -10 + 9 = -1
	fmt.Printf("-10 模 3= %v\n", (-10 % 3)) 
	//10 % -3 = 10 - 10 / (-3) * -3 = 10 - 9 = 1  
	fmt.Printf("10 模 -3= %v\n", (10 % -3)) 
	// -10 % -3 = (-10) - (-10) / (-3) * -3 = -10 + 9 = -1 
	fmt.Printf("-10 模 -3= %v\n", (-10 % -3)) 

	//++ 和 --
	//在go中 有这样几个特点
	//(1) 只有  a++  a-- , 没有其他使用方法
	//(2) 使用时，只能当做一个独立的语句时候用

	var num1 int = 10
	num1++ // 等价于  num1 = num1 + 1
	num1++
	num1-- // 等价于  num1 = num1 - 1
	fmt.Printf("num1=%d\n", num1) // 11

}
//关系操作符
func RelationalOperator(){
	var n int = 8
	var m int = 9
	fmt.Println(n >= m)
	fmt.Println(n <= m)
	fmt.Println(n == m)
	fmt.Println(n > m)
	fmt.Println(n < m)
	fmt.Println(n != m)
	flag := n >= m
	fmt.Println(flag)
}
//逻辑运算符 
func LogicalOperators()  {
	var A bool = true
	var B bool = false
	fmt.Println("A&&B", A && B) // false
	fmt.Println("A||B", A || B) // true
	fmt.Println("!(A||B)", !(A || B)) // false
	fmt.Println("!A", !!A) // true
}
// 赋值运算符的基本使用。
func getVal(n int) int {
	return 2 * n
}
func AssignmentOperators()  {
	
// 赋值基本案例
// 有两个变量，a和b，要求将其进行交换，最终打印结果 

	var a int = 10
	var b int = 20
	//1. 使用中间变量，完成交换
	// temp := a 
	// a = b // 
	// b = temp // 
	// fmt.Printf("a=%d b=%d\n", a, b)

	//2. 不使用中间变量，完成交换[面试题]
	a = a + b // a = a + b
	b = a - b // b = a + b - b ==> b =a  
	a = a - b // a = a + b - a ==> a = b
	fmt.Printf("a=%d b=%d\n", a, b)

	//赋值运算符的左边 只能是变量,右边 可以是变量、表达式、常量值

	var num1 = 109
	var num2 int
	num2 = num1 // 右边 可以是变量
	num2 = 2 * 3 + 1 * 7 // 右边 是 表达式
	num2 = getVal(10) + 90 * 10
	fmt.Println("num2=", num2)

	// += 的使用案例

	num2 += 30 // num3 = num3 + 30


}