package main

import (
	"fmt"
	"go_code/grammar/function/defer/code/utils"
)
func main()  {
	/*
	res := utils.DeferDemo(10,20)
	fmt.Println("main res =",res)

	res2 := utils.DeferDemo2(10,20)
	fmt.Println("main res =",res2)

	res3 := utils.DeferDemo3(10,20)
	fmt.Println("main res =",res3)
*/
  	fmt.Println("------------------------------------------------")
	res4 := utils.DeferDemo4(10,20)
	fmt.Println("main res =",res4)
	fmt.Println("------------------------")
	res5 := utils.DeferDemo5(10,20)
	fmt.Println("main res =",res5)
	fmt.Println("------------------------------------------------")
}