package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Print
	// utils.BreakDemo1()
	// utils.BreakDemo2()
	BreakDemo3()
}

// 随机生成1-100的一个数，直到生成了99这个数，看看程序一共用了几次?
/*
分析：1、怎么生成一个随机数？
2、比较
*/

func BreakDemo1() {

	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		radom := rand.Intn(100) + 1
		fmt.Printf("生成的随机数：%d \n", radom)
		count++
		if radom == 99 {
			fmt.Printf("生成了99这个数,共计次数：%d \n", count)
			break
		}
	}
}

// 1) 1-100以内的数求和，求出当和 第一次大于20的当前数是多少
func BreakDemo2() {
	sum := 0
	for i := 1; i < 101; i++ {
		sum += i
		if sum > 20 {
			fmt.Printf("当第一次sum大于20时的数为%d \n", i)
			break
		}

	}

	fmt.Printf("当第一次sum大于20时的数时sum = %d \n", sum)
}

// 2) 实现登录验证，有三次机会，如果用户名为”张无忌” ,密码”888”提示登录成功，否则提示还有几次登录机会.
func BreakDemo3() {
	var account = ""
	var pwd int
	for num := 3; num > 0; num-- {

		// fmt.Println("请输入账户")
		// fmt.Scanln(&account)
		// fmt.Println("请输入密码")
		// fmt.Scanln(&pwd)
		fmt.Println("请输入账户名和密码")
		fmt.Scanf("%s %d \n", &account, &pwd)

		if account == "张无忌" && pwd == 888 {
			fmt.Printf("登陆成功\n")
			break
		} else {
			fmt.Printf("登陆失败，还剩%d次机会\n", num-1)
		}
	}
}
