package main

import "fmt"

func main() {
	var i1, i2, i3 int
	go putStr(i1, i2, i3)
	fmt.Printf("pustr: ")
	fmt.Println(i1, i2, i3)
	go getStr(i1, i2, i3)
}

func putStr(i1, i2, i3 int) {
	i1 = 1
	i2 = 2
	i3 = 3
}

func getStr(i1, i2, i3 int) {
	if i2 == 2 {
		fmt.Println(i1, i2, i3)
		return
	}
}
