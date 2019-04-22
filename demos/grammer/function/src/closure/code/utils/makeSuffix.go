package utils

import (
	"strings"
)

/*
1)编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，并返回一个闭
2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg后缀，则返回原文件名。
3)要求使用闭包的方式完成
4)strings.HasSuffix //提示会使用到函数

*/
// 使用闭包
func MakeSuffix1(suffix string) func(string) string  {
	return func (str string) string  {
		if strings.HasSuffix(str,suffix) {
			return str
		}else {
			return str + suffix
		}
	}
}

// 传统方式
func MakeSuffix2(suffix string, str string) string {
	
		
	if strings.HasSuffix(str, suffix) {
		return str
	} else {
		return str + suffix
	}

}
