package utils

import (
	
	"fmt"
)

func util()  {
	
}
// 被测试函数
func getSum(n int) (sum int) {
	
	for i := 0; i <= n; i++ {
		sum +=i
	}
	return
} 

func TestGetSum()  {
	//传统的测试方法，就是在main函数中使用看看结果是否正确
	res := getSum(10)

	if res == 55 {
		fmt.Println("测试成功")
	}else{
		fmt.Println("测试失败")
	}
}

