package utils

import (
	"fmt"
)

func main() {
	// utils.ContinueDemo1()
	// utils.ContinueDemo2()
	ContinueDemo3()
}

//continue实现 打印1——100之内的奇数[要求使用for循环+continue
/*
分析：
1、for 循环 1~100
2、奇数 %2 ！=0
*/
func ContinueDemo1() {
	for i := 1; i < 101; i++ {
		if i%2 == 0 { //是偶数
			continue
		}
		fmt.Printf("1~100的奇数有%d \n", i)
	}
}

// 从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入为0时结束程序【使用for循环 ，break, continue完成】
/*
分析：
1、键盘输入数
2、for 正数和负数
*/
func ContinueDemo2() {
	positiveNum := 0
	negativeNUm := 0
	var num int
	for {
		fmt.Println("请输入整数")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num < 0 {
			negativeNUm++
			continue
		}
		positiveNum++

	}
	fmt.Printf("正数的个数为%d,负数的个数是%d \n", positiveNum, negativeNUm)
}

/*
某人有100,000元,每经过一次路口，需要交费,规则如下:
当现金>50000时,每次交5%
当现金<=50000时,每次交1000
编程计算该人可以经过多少次路口,使用 for + break方式完成

分析：
1、变量 count account
2、if  account>50 account<50
3、account<=0 时结束
*/
func ContinueDemo3() {

	account := 100000.0
	count := 0
	for {
		if account > 50000 {
			account = account - account*0.05
			count++
			fmt.Printf("有100,000元经过第%d次路口剩余%f \n", count, account)
		} else if account > 1000 && account < 50000 {
			account = account - 1000
			count++
			fmt.Printf("有100,000元经过第%d次路口剩余%f \n", count, account)
		} else {
			break
		}
	}
	fmt.Printf("有100,000元可以经过%d次路口 \n", count)

}
