package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 在并行时循环迭代的常见并发模型
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
