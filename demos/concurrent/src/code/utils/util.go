package utils

import (
	"time"
	"fmt"
)

func util()  {
	
}

// 判断素数(1~100)


func PrimeNumber() bool {
	falg := true
	var arr  [] int
	for i := 1; i < 200000; i++ {
		falg = numberIsPrime(i)
		if falg {
			arr=append(arr,i)
		}
	}
	// fmt.Println("素数arr=",arr)
	fmt.Println("素数arr长度",len(arr))
    return true
}


func test(n int,m int,arr  [] int)  {
	falg := true
	
	for i := n; i < m; i++ {
		falg = numberIsPrime(i)
		if falg {
			arr=append(arr,i)
		}
	}
}

// 判断一个数是否是素数？
// 即只能被1或者自身整除的自然数（不包括1），称为素数/质数。
func numberIsPrime(num int) bool{
	for i := 2; i < num; i++ {
		if (num % i == 0) {
			return false;
		}
	}
	return true;

}

func CodeTime()  {
	start :=time.Now().Unix()
	fmt.Println("start时间",start)
	PrimeNumber()
	end :=time.Now().Unix()
	fmt.Println("end时间",end)
	res := end -start
	fmt.Printf("时间差:%d \n",res)
}