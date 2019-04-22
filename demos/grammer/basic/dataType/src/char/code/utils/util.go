package utils

import (
	"fmt"
)

func CharDemo()  {
	var char1 byte = 'a' // 97
	var char2 byte = '\t'
	var char3 int = '北' // 对应编码值 21271 
	var char4 byte = 97 

	//如果我们直接输出一个byte，其实是输出该字符对应的编码值[ASCII]
	//如果我们直接输出一个byte，对应的字符 %c
	//总结
	//1. 当我们将一个字符存放到byte ，或者 int时，本质存放的是该字符对应的编码值(整数)
	//2. 如果我们希望输出对应的字符，则使用%c, 如果希望输出 对应的编码值 %v
	//3. 默认是输出编码值
	fmt.Printf("char1 = %v 字%c符是=%c  char3字符是=%c char4对应的字符=%c\n", 
		char1, char2, char1, char3, char4) 

	var n1 int = 38889
	fmt.Printf("n1=%c ", n1) // 输出’韩‘

	var char5 byte = 'a'
	var n10 = char5 + 1  //  98 n10[byte]
	fmt.Printf("n10 类型=%T, 值=%v 对应的字符=%c", n10, n10, n10) // ==> 97+1
}