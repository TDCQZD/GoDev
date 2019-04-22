package utils

import (
	"strconv"
	"fmt"
	"time"
)

func TimeDateDemo()  {
	fmt.Println("------------获取当前时间-----------")
	/*1、获取当前时间*/
	td :=time.Now()//返回类型就是time.Time
	fmt.Printf("td的类型是%T td的返回值是%v \n",td,td )
	fmt.Println("------------------------------------")

	fmt.Println("---------获取到其它的日期信息----------")
	fmt.Println("当前年份：",td.Year())
	fmt.Println("当前月份：",td.Month())
	fmt.Println("当前月份：",int(td.Month()))
	fmt.Println("当前日份：",td.Day())
	fmt.Println("当前日的时：",td.Hour())
	fmt.Println("当前日的分：",td.Minute())
	fmt.Println("当前日的秒：",td.Second())
	fmt.Println("------------------------------------")
   /*格式化日期时间*/
   fmt.Println("------------格式化日期时间一-----------")
	// 格式化的第一种方式
	
	fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n", 
	td.Year(), td.Month(), td.Day(), 
	td.Hour(), td.Minute(), td.Second())
	fmt.Println("------------------------------------")
	// 格式化的第二种方式：固定格式
	fmt.Println("------------格式化日期时间二-----------")
	fmt.Printf(td.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(td.Format("15:04:05 2006/01/02"))
	fmt.Println()
	fmt.Printf(td.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(td.Format("15:04:05"))
	fmt.Println()
	fmt.Println("------------------------------------")

	//休眠
	fmt.Println("--------------休眠-------------------")
	// time.Sleep(1000 * time.Millisecond)  //休眠1秒
	time.Sleep( time.Second)  //休眠1秒
	fmt.Println("------休眠1秒-----")
	fmt.Println("------------------------------------")
	
	//获取当前unix 时间戳  和 unixnano 时间戳。(作用是可以获取随机数字)
	fmt.Println("-------获取当前unix 时间戳 和 unixnano 时间戳---------")
	fmt.Println("当前unix 时间戳:",td.Unix())
	fmt.Println("unixnano 时间戳:",td.UnixNano())
	fmt.Println("------------------------------------")
	

}
/*
编写一段代码来统计 函数执行的时间。 
*/
func CodeTime()  {
	start :=time.Now().Unix()
	fmt.Println("start时间",start)
	// TimeDateDemo()
	code()
	end :=time.Now().Unix()
	fmt.Println("end时间",end)
	res := end -start
	fmt.Printf("时间差:%d \n",res)
}

func code()  {
	str := "hello"
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
}