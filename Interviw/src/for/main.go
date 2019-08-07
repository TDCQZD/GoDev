package main

import (
	"fmt"
)

func main() {
	// print3(4)
	// 判断奇偶数
	// num := 0
	// fmt.Println(num%2 == 0)
}

/* 打印金字塔
  *     打印星星次数=1*2-1 =1   打印空格次数=3-1=2
 ***            打印星星次数=2*2-1=3    打印空格次数=3-2=1
*****   打印星星次数=3*2-1=5    打印空格次数= 3-3=0

*/
func print1(n int) {
	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 打印倒立金字塔
func print2(n int) {
	for i := n; i > 0; i-- {
		//空格
		for j := 0; j <= n-i; j++ {
			fmt.Print(" ")
		}
		//**
		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		//换行
		fmt.Println()

	}
}

/*
打印棱形
*/
func print3(n int) {

	//上三角
	for i := 1; i <= n; i++ { // i表示行
		//第一部分：空格
		for j := 0; j < n-i; j++ { // 5-i
			/*
			   i:1, j:0,1,2,3   j<4
			   i:2, j:0,1,2    j<3
			*/
			fmt.Print(" ")
		}
		//第二部分：*
		for k := 0; k < 2*i-1; k++ { // 2*i-1
			/*
			   i:1, j:0      j<1
			   i:2, j:0,1,2   j<3
			*/
			fmt.Print("*")
		}
		//第三部分：换行
		fmt.Println()
	}
	//下三角
	for i := 1; i < n; i++ {
		//空格
		for j := 0; j < i; j++ {
			/*
			  i：1，j:0,j<1
			  i:2,j:0,1,j<2
			  i:3,j:0,1,2,j<3
			*/
			fmt.Print(" ")
		}
		//**
		for k := 0; k < 2*n-1-2*i; k++ {
			fmt.Print("*")
		}
		//换行
		fmt.Println()

	}
}
