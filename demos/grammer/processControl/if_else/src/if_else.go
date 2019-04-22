package utils

import (
	"fmt"
	"math"
)

func main() {
	/*	fmt.Println("请输入年龄")
			var paramter1 int
			fmt.Scanln(&paramter1)
			utils.SingleBranch(paramter1)
			utils.DoubleBranch(paramter1)
			utils.SingleAndDoubleBranch1()
			var paramter2 ,paramter22 int
			fmt.Println("请输入2个int32型变量")
			fmt.Scanf("%d %d",&paramter2,&paramter22)
			utils.SingleAndDoubleBranch2(paramter2,paramter22)
			var paramter3,paramter4 float64
			fmt.Println("请输入2个float64型变量")
			fmt.Scanf("%f %f",&paramter3,&paramter4)
			utils.SingleAndDoubleBranch3(paramter3,paramter4)
			fmt.Println("请输入2个int32型变量")
			var paramter5,paramter6 int32
			fmt.Scanf("%d %d",&paramter5,&paramter6)
			utils.SingleAndDoubleBranch4(paramter5,paramter6)
			fmt.Println("请输入年份")
			var paramter7 int
			fmt.Scanln(&paramter7)
			utils.SingleAndDoubleBranch5(paramter7)

		   var paramter8 ,paramter9 ,paramter10 float64
			fmt.Println("请输入方程a b c的值")
			fmt.Scanf("%f %f %f",&paramter8,&paramter9,&paramter10)
			utils.MoreBranch(paramter8 ,paramter9 ,paramter10)

			var grade int
			var sex string
			fmt.Println("请输入百米赛的成绩和性别(F/M)")
			fmt.Scanf("%d %s %f",&grade,&sex)
			utils.NestingBranch(grade ,sex )
	*/
	var mouth, age int
	// fmt.Println("请输入月份")
	// fmt.Scanln(&mouth)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	fmt.Println("请输入月份和年龄")
	fmt.Scanf("%d %d", &mouth, &age)
	// fmt.Println(mouth,age)
	SalePrices(mouth, age)
}

/*
编写一个程序,可以输入人的年龄,
如果该同志的年龄大于18岁,则输出 "你年龄大
于18,要对自己的行为负责!"
*/
//单分支
func SingleBranch(age int) {
	fmt.Println("-----------------单分支------------------------")
	if age > 18 {
		fmt.Println("你年龄大于18,要对自己的行为负责!")
	}
	fmt.Println("-----------------单分支------------------------")
}

//双分支
func DoubleBranch(age int) {
	fmt.Println("-----------------双分支-------------------------")
	if age > 18 {
		fmt.Println("你年龄大于18,要对自己的行为负责!")
	} else {
		fmt.Println("你年龄小于于18,暂时不需要对自己的行为负责!")
	}

	fmt.Println("-----------------双分支-------------------------")
}

/*单分支和双分支练习题*/
func SingleAndDoubleBranch1() {
	fmt.Println("-----------------练习题1-------------------------")
	var x int = 4
	var y int = 1
	if x > 2 {
		if y > 2 {
			fmt.Println(x + y)
		}
		fmt.Println("atguigu")
	} else {
		fmt.Println("x is = ", x)
	}
	fmt.Println("-----------------练习题1-------------------------")
}

/* 1)	编写程序，声明2个int32型变量并赋值。判断两数之和，如果大于等于50，打印“hello world!
分析：1、声明2个int32型变量并赋值
	 2、计算两数之和
	 3、单分支语句
*/
func SingleAndDoubleBranch2(n, m int) {
	fmt.Println("-----------------练习题2-------------------------")
	// n := 23
	// m := 30
	sum := m + n
	if sum > 50 {
		fmt.Println("hello world")
	} else {
		fmt.Println("不满足两数之和大于等于50")
	}
	fmt.Println("-----------------练习题2-------------------------")

}

/* 2)	编写程序，声明2个float64型变量并赋值。判断第一个数大于10.0，且第2个数小于20.0，打印两数之和。
分析：1、声明变量并赋值
2、单分支判断连个变量
*/
func SingleAndDoubleBranch3(n, m float64) {
	fmt.Println("-----------------练习题3-------------------------")
	// var n float64
	// var m float64
	// n = 12.3
	// m = 13.3

	if n > 10.0 && m < 20.0 {
		sum := n + m
		fmt.Printf("两数和为%0.1f \n", sum)
	} else {
		fmt.Println("不满足第一个数大于10.0，且第2个数小于20.0")

	}

	fmt.Println("-----------------练习题3-------------------------")
}

/* 3)	【选作】定义两个变量int32，判断二者的和，是否能被3又能被5整除，打印提示信息
分析：
1、定义变量
2、计算两个变量的和
3、能被3又能被5整除实现，num%3==0&&num%5==0
4、输出信息
*/
func SingleAndDoubleBranch4(n, m int32) {
	fmt.Println("-----------------练习题4-------------------------")
	// var n ,m int32
	// n = 9
	// m = 6
	sum := n + m
	if sum%3 == 0 && sum%5 == 0 {
		fmt.Println("输入的两数之和是即能被3又能被5整除")
	} else {
		fmt.Println("输入的两数之和是不能被3同时又能被5整除")
	}
	fmt.Println("-----------------练习题4-------------------------")
}

/*
4)判断一个年份是否是闰年，闰年的条件是符合下面二者之一：
(1)年份能被4整除，但不能被100整除；(2)能被400整除

分析：
1、输入年份
2、判断
*/
func SingleAndDoubleBranch5(year int) {
	fmt.Println("-----------------练习题5-------------------------")
	if (year%4 == 0 && year%10 != 0) || year%400 == 0 {
		fmt.Printf("输入年份:%d是闰年\n", year)
	} else {
		fmt.Printf("输入年份:%d不是闰年\n", year)
	}
	fmt.Println("-----------------练习题5-------------------------")
}

/*多分支练习题*/
/*
求ax2+bx+c=0方程的根。a,b,c分别为函数的参数，如果：b2-4ac>0，则有两个解；
b2-4ac=0，则有一个解；b2-4ac<0，则无解；
提示1：x1=(-b+sqrt(b2-4ac))/2a
               X2=(-b-sqrt(b2-4ac))/2a
提示2：math.Sqrt(num); 可以求平方根 需要引入 math包

测试数据: 3 ,100, 6
分析：
1、输入  a b c 三个变量 float64
2、计算b^2 -4ac 的结果
3、判断解
4、输出
*/

func MoreBranch(a, b, c float64) {

	res := b*b - 4*a*c

	if res > 0 {
		//(-b+sqrt(b2-4ac))/2a
		x1 := (-b + math.Sqrt(res)) / (2 * a)
		x2 := (-b - math.Sqrt(res)) / (2 * a)
		fmt.Printf("方程有两个解x1=%.1f x2=%.1f \n", x1, x2)
	} else if res == 0 {

		// x1 :=(-b + math.Sqrt(res)) /(2 * a)
		x2 := (-b - math.Sqrt(res)) / (2 * a)
		fmt.Printf("方程只有一个解x=%.1f \n", x2)
	} else {
		fmt.Println("方程无解")
	}
}

/*嵌套分支*/
/*
参加百米运动会，如果用时8秒以内进入决赛，否则提示淘汰。
并且根据性别提示进入男子组或女子组。【可以让学员先练习下】, 输入成绩和性别，进行判断

分析：
1、两个变量 成绩 grade float64 性别 sting
2、判断是否进入决赛
3、进入决赛后判断性别
4、输出
*/
func NestingBranch(grade int, sex string) {

	if grade < 8 {
		if sex == "M" {
			fmt.Println("恭喜进入男子决赛组")
		} else if sex == "F" {
			fmt.Println("恭喜进入女子决赛组")
		} else {
			fmt.Println("性别输入错误！！")
		}
	} else {
		fmt.Println("你被淘汰了！")
	}
}

/*嵌套分支
出票系统：根据淡旺季的月份和年龄，打印票价 [学生先做]

4_10 旺季：
	成人（18-60）：全价
	儿童（<18）:半价
	老人（>60）:1/3

淡季：
	成人：40
	其他：20

分析：


*/

func SalePrices(mouth, age int) {
	fmt.Println(mouth, age)
	if (mouth <= 0) || (mouth > 12) {
		fmt.Println("月份输入错误请重新输入")
	} else {
		if mouth > 3 && mouth < 11 {
			if age < 18 {
				fmt.Println("需要支付半价")
			} else if age > 60 {
				fmt.Println("需要支付1/3")
			} else {
				fmt.Println("需要支付全价")
			}
		} else {
			if age >= 18 && age <= 60 {
				fmt.Println("需要支付40")
			} else {
				fmt.Println("需要支付20")
			}
		}
	}
}
