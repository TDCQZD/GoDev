package utils

import (
	"fmt"
)

func PointDemo()  {
	var num1 int = 888

	//获取到num1的存储空间的地址
	fmt.Printf("num1的存储空间的值 = %v num1的存储空间的地址=%p\n", num1, &num1)
	var num2 int = 999
	//获取到num2的存储空间的地址
	fmt.Printf("num2的存储空间的值 = %v num2的存储空间的地址=%p\n", num2, &num2)
	//指针，我们让一个指针指向 num2
	var ptr1 *int = &num2
	fmt.Printf("ptr 本身存储空间的值=%v  ptr 本身存储空间的地址=%p ptr 指向的值=%v", 
		ptr1, &ptr1, *ptr1)
}

func PointDemo1()  {
	//写一个程序，获取一个int变量num的地址，并显示到终端
	//将num的地址赋给指针 ptr , 并通过ptr去修改num的值.

	var num int = 123
	fmt.Printf("num的地址=%p num=%v \n", &num,num)

	var ptr *int = &num
	*ptr = *ptr + 1
	fmt.Printf("ptr指向的值=%v num=%v\n", *ptr, num) // 124, 124
	fmt.Printf("ptr的存储空间的值 = %v ptr的存储空间的地址=%p\n", ptr, &ptr)
}