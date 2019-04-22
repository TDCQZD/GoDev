package utils

import (
	"fmt"
)

func TypeConverDemo()  {
	var n1 int = 90
	var n2 float64 = float64(n1)

	n2 = n2 + 1.5

	var n3 int32 = int32(n2) //当将一个float类型转成 int时，是直接去掉小数点后的部分

	fmt.Printf("n1 type=%T val=%v n2 type=%T val=%v n3 type=%T val=%v\n", 
		n1, n1, n2, n2, n3, n3)


    
	var num1 int16 = 10
	var num2 int32 = 20
	var num3 int64 = 30

	num1 = int16(num2) // int32-->int16
	num3 = int64(num2) // int32-->int64

	fmt.Println(num1, num2, num3)
 // 被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化！
	fmt.Printf("num1 type=%T num1=%v num2 type=%T num2=%v num3 type=%T num3=%v\n", 
	num1, num1, num2, num2, num3, num3)

	//在转换中，比如将 int64  转成 int8  ，如果值超出范围，编译时不会报错，
	//只是转换的结果是按溢出处理，和我们希望的结果不一样,

	var num4 int64 = 123456
	var num5 int8
	//int64->int8
	num5 = int8(num4) // 转换后 num4 仍然是 int64,而且值还是 123456!
	fmt.Printf("num4=%d \n", num4)
	fmt.Printf("num5=%d \n", num5)


	// var n1 int32 = 12
	// var n3 int8
	// var n4 int8
	// n4 = int8(n1) + 127  //  编译能够过,  结果是溢出处理
	// n3 = int8(n1) + 128 //   编译不过 128> int8最大的值127
	// fmt.Println(n3)
}