package main

import (
	"fmt"
)

// 要求：可以从控制台接收用户信息，【姓名，年龄，薪水, 是否通过考试 】。
func KeyDemo() {
	var name string
	var age int
	var sal float64
	var isPass bool

	fmt.Println("请输入name")
	fmt.Scanln(&name)

	fmt.Println("请输入age")
	fmt.Scanln(&age)

	fmt.Println("请输入sal")
	fmt.Scanln(&sal)

	fmt.Println("请输入isPass")
	fmt.Scanln(&isPass)

	fmt.Printf("name=%s age=%d sal=%f ispass=%t", name, age, sal, isPass)

	//2. 使用fmt.Scanf， 该方法可以一次性输入多个值

	fmt.Println("请输入name age sal ispass:")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass) //输入时，必须使用空格来区分不同值
	fmt.Printf("name=%s age=%d sal=%f ispass=%t", name, age, sal, isPass)

}
