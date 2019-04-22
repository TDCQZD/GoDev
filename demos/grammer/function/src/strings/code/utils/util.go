package utils

import (
	"strings"
	"strconv"
	"fmt"
)

func StringsFunction()  {
	str :="sadjksad嫌这嫌那analsdjflnlsanfln"
	fmt.Println("源字符串=",str)
// 1.	统计字符串的长度，按字节 len(str)
 	//len默认是按照字节的大小返回
	res := len(str)
	fmt.Println("len默认是按照字节的大小返回字符串的长度=",res)
	str2 :=[]rune(str)
	// 照字符的个数统计，需要先转成[]rune	
	res2 := len(str2)
	fmt.Println("[]rune按照字符的大小返回字符串的长度=",res2)
// 2.	字符串遍历，同时处理有中文的问题 r := []rune(str)
	for index,value  := range str {
		fmt.Printf("普通字符串遍历:str[%d]=%c \n",index,value )
	}
	for index,value  := range str2 {
		fmt.Printf("[]rune(str)字符串遍历:str[%d]=%c \n",index,value )
	}

// 3.	字符串转整数:	 n, err := strconv.Atoi("12") 
	str3 :="123325436"
	res3, err := strconv.Atoi(str3)
		if err != nil {
			fmt.Println("字符串转整数,转换错误err=", err)
		} else {
			fmt.Printf("字符串转整数,num1 类型%T num1的值=%d \n", res3, res3)
		}
// 4.	整数转字符串  str = strconv.Itoa(12345)
	str4 :=123325436
	res4 := strconv.Itoa(str4)
	fmt.Printf("整数转字符串,res4 类型%T res4的值=%s \n", res4, res4)

// 5.	字符串 转 []byte:  var bytes = []byte("hello go")
	   res5 := []byte("hello go")
	   fmt.Printf("字符串 转 []byte,res5 类型%T res5的值=%s \n", res5, res5)
// 6.	[]byte 转 字符串: str = string([]byte{97, 98, 99})
	   res6 := string([]byte{97, 98, 99})
	   fmt.Printf("[]byte 转 字符串,res6 类型%T res6的值=%s \n", res6, res6)
// 7.	10进制转 2, 8, 16进制:  str = strconv.FormatInt(123, 2) // 2-> 8 , 16
	
		res7 := strconv.FormatInt(123, 8) // 2-> 8 , 16
		res8 := strconv.FormatInt(123, 16) // 2-> 8 , 16
		res9 := strconv.FormatInt(123, 2) // 2-> 8 , 16
		fmt.Printf("10进制转 2进制值=%s \n", res9)
		fmt.Printf("10进制转 8进制值=%s \n", res7)
		fmt.Printf("10进制转 16进制值=%s \n", res8)
// 8.	查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
		res10 := strings.Contains("seafood", "foo") //true
		fmt.Printf("查找子串是否在指定的字符串中 =%t \n", res10)
// 9.	统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4
		res11 := strings.Count("ceeeeertyuxdeheese", "e") //4
		fmt.Printf("统计一个字符串有几个指定的子串 =%d \n", res11)
// 10.	不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true
		res12 := strings.EqualFold("abc", "Abc")
		fmt.Printf("不区分大小写的字符串比较 =%t \n", res12)
// 11.	返回子串在字符串第一次出现的index值，如果没有返回-1 : strings.Index("NLT_abc", "abc") // 4
		res13 := strings.Index("NtukhbLT_abc", "abc")
		res14 := strings.Index("NLT_abc", "d")
		fmt.Printf("返回子串在字符串第一次出现的index值 =%d \n", res13)
		fmt.Printf("返回子串在字符串第一次出现的index值 =%d \n", res14)
// 12.	返回子串在字符串最后一次出现的index，如没有返回-1 : strings.LastIndex("go golang", "go")
		res15 := strings.LastIndex("go golang", "go")
		res16 := strings.LastIndex("go golang", "b")
		fmt.Printf("返回子串在字符串最后一次出现的index =%d \n", res15)
		fmt.Printf("返回子串在字符串最后一次出现的index =%d \n", res16)
// 13.	将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go语言", n) n可以指定你希望替换几个，如果n=-1表示全部替换
		res17 := strings.Replace("go go hello", "go", "go语言", 1)
		res18 := strings.Replace("go go hello", "go", "go语言", -1)
		fmt.Printf("将指定的子串替换成 另外一个子串 =%s \n", res17)
		fmt.Printf("将指定的子串替换成 另外一个子串 =%s \n", res18)
// 14.	按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组： strings.Split("hello,wrold,ok", ",") 
		res19 := strings.Split("hello,wrold,ok", ",")
		fmt.Printf("一个字符串拆分成字符串数组 =%v \n", res19)
// 15.	将字符串的字母进行大小写的转换: strings.ToLower("Go") // go strings.ToUpper("Go") // GO
		res20 := strings.ToLower("Go") // go 
		res27 := strings.ToUpper("Go") // GO
		fmt.Printf("将字符串的字母进行小写的转换 =%s \n", res20)
		fmt.Printf("将字符串的字母进行大写的转换 =%s \n", res27)
// 16.	将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn   ")
		res21 := strings.TrimSpace("          tn a lone gopher ntrn   ")
		fmt.Printf("将字符串左右两边的空格去掉 =%s \n", res21)
// 17.	将字符串左右两边指定的字符去掉 ： strings.Trim("! hello! ", " !")  // ["hello"] //将左右两边 ! 和 " "去掉
		res22 := strings.Trim("! hello! ", " !") 
		fmt.Printf("将字符串左右两边指定的字符去掉 =%s \n", res22)
// 18.	将字符串左边指定的字符去掉 ： strings.TrimLeft("! hello! ", " !")  // ["hello"] //将左边 ! 和 " "去掉
		res23 := strings.TrimLeft("! hello! ", " !")
		fmt.Printf("将字符串左边指定的字符去掉 =%s \n", res23)
// 19.	将字符串右边指定的字符去掉 ： strings.TrimRight("! hello! ", " !")  // ["hello"] //将右边 ! 和 " "去掉
		res24 := strings.TrimRight("! hello! ", " !")
		fmt.Printf("将字符串右边指定的字符去掉 =%s \n", res24)
// 20.	判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
		res25 := strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
		fmt.Printf("判断字符串是否以指定的字符串开头 =%t \n", res25)
// 21.	判断字符串是否以指定的字符串结束: strings.HasSuffix("NLT_abc.jpg", "abc") //false
		res26 := strings.HasSuffix("NLT_abc.jpg", "abc") //false
		fmt.Printf(" 判断字符串是否以指定的字符串结束=%t \n", res26)
	}