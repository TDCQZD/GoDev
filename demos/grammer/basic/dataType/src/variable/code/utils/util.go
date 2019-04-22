package utils

import (
	"fmt"
)
/**/
/*变量的使用*/
func VariableDemo()  {
	var num int  //声明（定义）变量
	num = 9      //变量赋值
	fmt.Println(num)//变量使用
}
/*Golang变量使用的三种方式*/
func VariableDemo2()  {
	//Golang变量使用的三种方式
	//(1) 第一种：指定变量类型，声明后若不赋值，使用默认值

	var num2 int
	fmt.Println("num2=", num2) // 整数类型的默认值都是 go 0 
	
	//(2) 第二种：根据值自行判定变量类型(类型推导) 

	var num3 = 890 //编译器，就会进行类型推导，num3就是int
	var num4 = 1.2 //编译器，就会进行类型推导，num4就是float...
	fmt.Printf("num3 类型=%T, val= %v\n", num3, num3)
	fmt.Printf("num4 类型=%T, val= %v\n", num4, num4)

	//(3) 第三种：省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误

	num5 := 12 //等价两句话  var num5 int ; num5 = 12
	num6 := "tom"
	fmt.Printf("num5 类型=%T, val= %v\n", num5, num5)
	fmt.Printf("num6 类型=%T, val= %v\n", num6, num6)
}


/*

func VariableDemo() {

}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
func VariableDemo() {
	
}
*/