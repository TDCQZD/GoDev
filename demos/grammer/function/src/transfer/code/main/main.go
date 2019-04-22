package main

import (
	"fmt"
	transfer "go_code/grammar/function/transfer/code/others"
)
/*
编写函数，计算两个数的和和差，并返回结果
*/
func main(){
	var num1 float64
	var num2 float64
	fmt.Println("请输入运算的两个数字")
	fmt.Scanf("%f %f \n",&num1 ,&num2)
	sum ,dec := transfer.Calculation(num1,num2)
	fmt.Printf("计算结果sum= %.1f des=%.1f \n",sum ,dec)
}