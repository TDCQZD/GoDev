package utils

import (
	"fmt"
)

func util()  {
	
}
func FloatDemo() {
	
	var n1 float32 = -1238888887788999.0000901 //-123.0000901
	var n2 float64 = -1238888887788999.0000901 //-123.0000901
	var n3 float32 = 1.2 // n3 就是 float64

	var n4 float32
	var n5 = 1.1 //n5 = float64
	n4 = n3 + float32(n5) //在go运算中，如果是变量，则必须考虑其类型，如果是常量(字面量),考虑范围

	fmt.Printf("n1=%v n2=%v n4=%v \n", n1, n2, n4)

	//浮点型常量有两种表示形式
	//十进制数形式：如：5.12       .512   (必须有小数点）
	//科学计数法形式:如：5.12e2  = 5.12 * 10的2次方     5.12E-2  = 5.12/10的2次方   

	n6 := 5.12
	n7 := 0.512
	n8 := .512
	fmt.Printf("n6=%v n7=%v n8=%v\n",n6, n7, n8)
	// 512 
	n9 := 5.12e2 // 科学计数法  == 5.12 * 10的2次方  ==> 512
	n10 := 5.12e-2 // 科学计数法  == 5.12 / 10的2次方 ==> 0.0512
	fmt.Printf("n9=%v n10=%v\n",n9, n10)
}