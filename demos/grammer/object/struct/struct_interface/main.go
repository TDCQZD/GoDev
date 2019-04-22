package main

import (
	"fmt"
)

// 嵌入interface可以使得一个struct具有interface的接口，
// 而不需要实现interface中的有声明的函数。
type II interface {
	F1()
	F2()
}

type SS struct {
	vv int
	II
}

func (ss *SS) F1() {
	fmt.Printf("in SS::F1()\n")
}

func (ss *SS) F2() {

	fmt.Printf("in SS::F2()\n")
}

func main() {
	var ss SS = SS{}
	var ii II = &ss

	fmt.Printf("ss:value=[%v]\n", ss)
	fmt.Printf("ii:value=[%v]\n", ii)
	ss.F2()
}
