package utils

import (
	"fmt"
)

// for使用方式一
func ForDemo1()  {
	for i := 0; i < 10; i++ {
		fmt.Println("for使用方式一",i)
	}
	fmt.Println("---------------------------------------")

}

// for使用方式二 使用for循环实现while的效果
func ForDemo2()  {
	i := 0
	for  i < 10 {	
		fmt.Println("for使用方式二",i)
	}
	i++
	fmt.Println("---------------------------------------")
}

// for使用方式三 使用for循环do..while的效果实现
func ForDemo3()  {
	i := 0
	for  {
		i++
		if i >10 {
			break
		}
		
		fmt.Println("for使用方式三",i)
	}
	fmt.Println("---------------------------------------")

}

//for遍历
func ForDemo4()  {
	str := "qjmkfa中国c东方oomccsm,cfd"
	//for range 按字符进行遍历，如果出现中文，不会乱码
	for i , j := range str {
		fmt.Printf("第%d个下标的值是%c \n",i,j)
	}
	fmt.Println("---------------------------------------")
	// 按字节进行遍历，如果出现中文，会乱码，解决方法切片
	for i1 := 0; i1<len(str); i1++ {
		fmt.Printf("第%d个下标的值是%c \n",i1,str[i1])
	}
	fmt.Println("---------------------------------------")
	//切片解决普通遍历中文乱码
	str2 := []rune(str)
	for i2 := 0; i2<len(str2); i2++ {
		fmt.Printf("第%d个下标的值是%c \n",i2,str2[i2])
	}
	fmt.Println("---------------------------------------")
}

// 1) 打印1~100之间所有是9的倍数的整数的个数及总和
/*
分析：
1、for循环遍历1~100
2、是9的倍数的整数 if i%9 ==0
2、变量 sum 加2中i count 计数
*/
 func ForDemo5()  {
	var count = 0 
	var sum = 0
	 for i := 1; i < 101; i++ {
		 if i % 9 == 0 {
			count++
			sum += i
		 }
	 }
	 fmt.Printf("1~100之间所有是9的倍数的整数的个数为:%d及总和:%d \n",count,sum)
 }
/*
1) 统计3个班成绩情况，每个班有5名同学，求出各个班的平均分和所有班级的平均分[学生的成绩从键盘输入]
分析：
1、变量 classromNum 班级 studentNum 每班的学生人数
2、第一层for循环 循环班级
3、第二层for循环 循环班级人数 计算平均分
4、结束第一层循环，计算平均成绩
*/
func ForDemo6()  {

	classromNum := 3
	studentNum := 5
	var grade float64 
	var classromAver, classromSum float64  //每个班级的平均成绩 总分
	var allClassromAver , allClassromSum float64 //所有班级的平均分 总分
	for i := 1; i <= classromNum; i++ {
		for j := 1; j <= studentNum; j++ {
		fmt.Printf("请输入第%d个班级第%d位学生成绩\n",i,j)
		fmt.Scanln(&grade)
		classromSum +=	grade
		}
		classromAver = classromSum / float64(studentNum)
		fmt.Printf("第%d个班级共计%d名学生平均成绩%f\n",i,studentNum,classromAver)
		allClassromSum +=classromSum
		classromSum=0
	}
	num := classromNum * studentNum
	allClassromAver = allClassromSum / float64(num)
	fmt.Printf("%d个班级共计所有%d学生平均成绩%f\n",classromNum,num,allClassromAver)
	
}


	/*
	编写一个程序，可以接收一个整数,表示层数，打印出金字塔
		*
	   ***
	  *****
	  分析：
	  1、两层循环
	  2、第一层：行数
	  3、第二层：循环打印空格和*
	  4、只需要计算*前面的空格 和*的个数
	  5、实现空心：在*的逻辑上修改，处理第一个和最后一个*保留，其余替换为空格
	  6、最后一行 h =i 所有* 保留
	*/

func ForDemo7()  {

	h := 7

	for i := 1; i <= h; i++ {
		for k := 1; k <= h - i; k++ {
			fmt.Printf(" ")		
		}
		for j := 1; j <= 2 * i-1; j++ {
			if j == 1 || j == 2 * i-1 || i == h{
				fmt.Printf("*")
			}else {
				fmt.Printf(" ")	
			}
			
		}
		fmt.Printf("\n")
	}
	
}

func ForDemo9()  {
	/*
		*
	   ***
	  *****

	*/

	//分析
	//1. 先打印出
	/*

		****
		****
		****

	*/



	//2 打印
	/*
		*        //第1行打印 1 *   即： 2 * i - 1 
		***      //第2行打印 3 *
		*****    //第3行打印 5 *

	*/

	//3 打印

	/*
	    *     //第1行打印 空格  2   = 总的行数  - i = 3 - 1 
	   ***    //第2行打印 空格  1   = 3 - 2
	  *****   //第3行打印 空格  0   = 3 - 3
	*/

	//4 打印 空心的 ， 什么情况下打印 * ,什么情况下打印空格

	/*
	    *
	   * *
	  *****
	*/

	//将 3 改成一个变量
	var totalLayer = 20

	//代码实现
	for i := 1; i <= totalLayer; i ++ {  //i表示 第几行 
		//先打印出空格
		for k := 1; k <= totalLayer - i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2 * i - 1; j++ {  //j 表示列
			//什么情况下打印 * ,什么情况下打印空格
			//在当前这个情况下的第一个和最后一个打印*,其它打印空格
			if j == 1 || j ==  2 * i - 1 || i ==  totalLayer {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

// 4) 打印出九九乘法表
/*
分析：1、两层循环
2、第一层：循环1~9
3、第二层：循环1~i
*/
func ForDemo8()  {
	for i := 1 ; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d       ",j,i,j*i)
		}
		fmt.Printf("\n")
	}
}