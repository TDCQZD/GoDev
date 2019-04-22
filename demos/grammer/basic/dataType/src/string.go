package main

import "fmt"

func main() {
	s := "hello"
	// []byte
	c := []byte(s) // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c) // 再转换回 string 类型
	fmt.Printf("%s\n", s2)
	// +
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)

	// ss:= "sdfghjk"
	// +"weryuio"
}
