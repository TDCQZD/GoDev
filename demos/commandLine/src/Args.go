package main

import (
	"fmt"
	"os"
)

// Args 查看命令行输入参数
func Args() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
