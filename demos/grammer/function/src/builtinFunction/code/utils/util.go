package utils

import (
	"fmt"
)

func BuiltinDemo()  {
	
	/*new：用来分配内存，主要用来分配值类型，比如int、float32,struct..*/
	var n float64 =33.3333
	fmt.Printf("n的类型=%T n的值为 %f \n",n, n)

	//new 给值类型分配内存
	var m *float64 = new (float64)
	*m = 44.333
	fmt.Printf("m 的类型= %T \nm指向的值=%v \nm本身的值=%p \nm=%p\n", 
	m, *m, m, &m)
	//通过指针改变变量值
	m1 := m
	*m1 = 45.6666
	fmt.Printf("m指向值=%v \n", *m)
	*m++
	fmt.Printf("m指向值=%v \n", *m) // 5.56
	fmt.Printf("m1指向值=%v \n", *m1) // 5.56
}