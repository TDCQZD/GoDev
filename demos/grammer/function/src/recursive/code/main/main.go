package main

import (
	"fmt"
   "go_code/grammar/function/recursive/code/utils"
)

func main()  {
	utils.Test(6)
	fmt.Println("---------------------------")
	utils.Test1(6)
	fmt.Println("---------------------------")
	num := utils.Fbnqs(3)//斐波那契数列
	fmt.Println(num)
	fmt.Println("---------------------------")
	num1 :=utils.Test2(3)
	fmt.Println(num1)

	fmt.Println("---------------------------")
	num2 :=utils.Peachs(8)
	fmt.Printf("桃子总数为=%d \n",num2)
}
	