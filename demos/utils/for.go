package main

import "fmt"

func main() {
	i := 0
	j := 0
	for {
		if i > 100 {
			break
		}
		fmt.Println("i:", i)
		i++
	}
	for {
		if j > 100 {
			break
		}
		fmt.Println("j:", j)
		j++
	}
}
