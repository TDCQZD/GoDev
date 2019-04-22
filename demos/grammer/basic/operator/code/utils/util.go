package utils

import (
	"fmt"
)
/*算术运算符*/
// 假如还有97天放假，问：xx个星期零xx天
func OPeratorDemo1()  {
	var i = 97
	w := i / 7
	d := i % 7

	fmt.Printf("还有97天放假剩余 %d 个星期零 %d 天 \n",w,d)

	var t = 235.5
	dd := 5.0 / 9 * ( t - 32 )
	fmt.Printf("%.1f华氏温度对应的摄氏温度=%.1f ",t,dd)
}

// 定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100),请求出华氏温度对应的摄氏温度。[比如华氏温度：232.5]
func OPeratorDemo2()  {


	var t = 235.5
	dd := 5.0 / 9 * ( t - 32 )
	fmt.Printf("%.1f华氏温度对应的摄氏温度=%.1f ",t,dd)
}

/*赋值运算符*/
// 面试题：有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
/*
分析：
1、res=a+b
2、 a=res - b
3、b =res - b
*/
// 不允许使用中间变量
func OPeratorDemo3()  {
	a := 20
	b := 30
	fmt.Printf("交换前a=%d b=%d \n",a,b)
	a = a + b
	b = a - b
	a = a - b 
	fmt.Printf("交换后a=%d b=%d \n",a,b) 
}
//使用中间变量
func OPeratorDemo6()  {
	a := 20
	b := 30
	fmt.Printf("交换前a=%d b=%d \n",a,b)
	temp := a 
	a =  b 
	b = temp
	fmt.Printf("交换后a=%d b=%d \n",a,b) 
}

// 案例1：求两个数的最大值
func OPeratorDemo4()  {
	a := 10 
	b := 20
	var max int
	var min int
	if a > b{
		max = a
		min = b
	}else{
		max = b
		min = a
	}
	fmt.Printf("两个数中最大值max=%d 最小值=%d \n",max,min) 
}
// 案例2：求三个数的最大值
func OPeratorDemo5()  {
	a := 10 
	b := 20
	c := 30
	max ,min := dealData(a,b,c)
	fmt.Printf("三个数中最大值max=%d 最小值=%d \n",max,min) 
}

func dealData(n1 int, n2 int, n3 int) (m int, n int) {
	var max int
	var min int
	if n1 > n2 {
		max = n1
		min = n2
	} else {
		max = n2
		min = n1
	}
	if max < n3 {
		max = n3
	}
	if min > n3 {
		min = n3
	}

	return max, min
}