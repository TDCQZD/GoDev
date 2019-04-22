package utils

import (
	"fmt"
)


//画内存变化图

func Test(n int)  {
	if n > 2{
		n--
		Test(n)
	}
	fmt.Printf("n= %d \n" ,n)
}

func Test1(n int)  {
	if n > 2{
		n--
		Test1(n)
	}else{
	fmt.Printf("n= %d \n" ,n)
	}
}
/*
求函数值
已知 f(1)=3; f(n) = 2*f(n-1)+1; 
请使用递归的思想编程，求出 f(n)的值?
分析：
1、当 n=1时，f(n)=3
2、从n=2开始，数学公式是f(n)=2f(n-1)+1
*/

func Test2(n int) int {
	if n == 1{
		return 3
	}else{
	return 2 * (Test2(n-1)) + 1
	}
}

/*
猴子吃桃子问题
有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！以后每天猴子都吃其中的一半，然后再多吃一个。
当到第十天时，想再吃时（还没吃），发现只有1个桃子了。问题：最初共多少个桃子？
分析：(反向分析)
数学表达式：f(n)=f(n-1)/2 + 1  ====>f(n-1)=2(f(n)+1)===>f(n)=2(f(n+1)+1)
1、当n=10 f(10)=1
2、当n=9 f(9)=2(1+1)=4
3、当n=8 f(8)=2(4+1)=10
.
.
.
4、当 n=1时，即可求出桃子总数
*/

func Peachs(n int ) int  {
	if n == 10{
		return 1
	}else{
		return 2*(Peachs(n+1) + 1)
	}
}