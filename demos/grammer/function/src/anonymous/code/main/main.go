package main

import (
	"fmt"
	"go_code/grammar/function/anonymous/code/utils"
)

func main(){
	res1 := utils.Anonymous1(1,2)// 在定义匿名函数时就直接调用
	fmt.Println(res1)
	fmt.Println("-----------------------")
	res2 :=utils.Anonymous2(1,2)// 将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数
	fmt.Println(res2)
	fmt.Println("-----------------------")
	res3 :=utils.Anonymous3(1,2)//全局匿名函数
	fmt.Println(res3)
	fmt.Println("-----------------------")
}