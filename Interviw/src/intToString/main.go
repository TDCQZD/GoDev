package main

import (
    "fmt"
)
// 数组和字符串之间的相互转换[]byte/string
func main() {
    str := "hello"
    arr := []byte(str)
    fmt.Println(arr)
    str2 := string(arr[:])
    fmt.Println(str2)
}