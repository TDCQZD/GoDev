package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// str := "XHelloWorldX"
	// content := str[1 : 1]
	// fmt.Println(content)

	// A := []int{4, 2, 3}
	// k := 1
	// n := largestSumAfterKNegations(A, k)
	// fmt.Println("sum->", n)
	// AMap := make(map[int]int, 201)
	// for t := 0; t < 201; t++ {
	// 	AMap[t] = 0
	// }
	// fmt.Println("num->", len(AMap))
	// s := []string{"bob,689,1910,barcelona", "alex,696,122,bangkok", "bob,832,1726,barcelona", "bob,820,596,bangkok", "chalicefy,217,669,barcelona", "bob,175,221,amsterdam"}
	// res := invalidTransactions(s)
	// fmt.Println(res)
	res := fairCandySwap([]int{1, 1}, []int{2, 2})
	fmt.Println(res)
}
func fairCandySwap(A []int, B []int) []int {
	sa, sb := 0, 0            // sum of A, B respectively
	setB := make(map[int]int) // If Alice gives x, she expects to receive x + delta
	for _, x := range A {
		sa += x
	}
	for _, y := range B {
		sb += y
		setB[y] = 0
	}

	delta := (sb - sa) / 2
	fmt.Println(delta, setB)
	for _, x := range A {
		if _, ok := setB[x+delta]; ok {
			return []int{x, x + delta}
		}
	}
	return nil
}
func invalidTransactions(transactions []string) []string {
	if nil == transactions || len(transactions) < 1 {
		return nil
	}
	inval := []string{}
	stringMap := make(map[string]string)
	for i := 0; i < len(transactions); i++ {
		s := strings.Split(transactions[i], ",")
		sToi, _ := strconv.Atoi(s[2])

		if ms, ok := stringMap[s[0]]; ok {
			mapEle := strings.Split(ms, ",")
			mapEle1, _ := strconv.Atoi(mapEle[1])
			mapEle2, _ := strconv.Atoi(mapEle[2])
			s2, _ := strconv.Atoi(s[1])
			if mapEle[3] != s[3] && abs(s2, mapEle1) < 60 {

				if sToi <= 1000 {
					inval = append(inval, transactions[i])
				}
				if mapEle2 <= 1000 {
					inval = append(inval, ms)
				}

			}
		}
		if sToi > 1000 {
			inval = append(inval, transactions[i])
			fmt.Println("3", inval)
		}
		stringMap[s[0]] = transactions[i]
	}
	return inval
}

func abs(a, b int) int {
	d := a - b
	if d > 0 {
		return d
	} else {
		return 0 - d
	}
}
func largestSumAfterKNegations(A []int, K int) int {
	AMap := make(map[int]int, 201) //-100 <= A[i] <= 100,这个范围的大小是201
	for _, v := range A {
		AMap[v+100]++ //将[-100,100]映射到[0,200]上
	}
	i := 0
	for K > 0 {
		for AMap[i] == 0 { //找到A[]中最小的数字
			i++
		}
		AMap[i]--     //此数字个数-1
		AMap[200-i]++ //其相反数个数+1
		if i > 100 {
			i = 200 - i //若原最小数索引>100,则新的最小数索引应为200-i.(索引即number[]数组的下标)
		}

		K--
	}

	sum := 0
	fmt.Println(len(AMap))
	for j := i; j < len(AMap); j++ { //遍历number[]求和
		fmt.Println((j - 100) * AMap[j])
		sum = sum + (j-100)*AMap[j] //j-100是数字大小,number[j]是该数字出现次数.
	}
	return sum
}
