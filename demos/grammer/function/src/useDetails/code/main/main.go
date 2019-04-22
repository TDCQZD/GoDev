package main 

import (
	"fmt"
	"go_code/grammar/function/useDetails/code/utils"
)

func main()  {
	
	n :=20
	utils.Point(&n)
	fmt.Printf("n的值是%d \n",n)
	fmt.Println("---------------------")
	res := utils.Variable
	res1 := utils.Variable(12,13)
	fmt.Printf("通过变量对函数调用的返回值是%d \n",res1)
	fmt.Println("---------------------")
	//函数可以作为形参
	res2 := utils.FormParameter(res,123,122)
	fmt.Printf("函数可以作为形参，并且调用返回的结果%d \n",res2)
	fmt.Println("---------------------")
	utils.NameDataType()
	// 11)支持对函数返回值命名
	sum ,sub := utils.NameFuncReturn(12314, 123)
	fmt.Printf("对函数返回值命名,调用返回的结果的和：%d 差：%d  \n",sum ,sub )
	fmt.Println("---------------------")
	// 12) 使用 _ 标识符，忽略返回值
	sum1 ,_ := utils.NameFuncReturn(12314, 123)
	fmt.Printf("对函数返回值命名,调用返回的结果的和：%d  \n",sum1 )
	fmt.Println("---------------------")

	res3 := utils.MoreArgs()
	res4 := utils.MoreArgs(1, 2,3,4,5)
	res5 := utils.MoreArgs1(1)
	res6 := utils.MoreArgs1(1, 2,3,4,5)
	fmt.Printf("支持0到正无穷个参数返回的结果的和：%d  \n",res3 )
	fmt.Printf("支持0到正无穷个参数返回的结果的和：%d  \n",res4 )
	fmt.Printf("支持1到正无穷个参数返回的结果的和：%d  \n",res5 )
	fmt.Printf("支持1到正无穷个参数返回的结果的和：%d  \n",res6 )
	fmt.Println("---------------------")

  	m := 30
	fmt.Printf("交换前n的值=%d m的值=%d \n",n,m )
	utils.SwapPoint( &n ,&m)
	fmt.Printf("交换后n的值=%d m的值=%d \n",n,m)
	fmt.Println("---------------------")
	
}