package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//
	// rand.Seed(time.Now().Unix())
	// rand.Seed(time.Now().UnixNano())

	// fmt.Println("My first lucky number is", rand.Intn(10))
	// fmt.Println("My senond lucky number is", rand.Intn(10))

	// 可重复
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println(r.Intn(10))
	}

	//
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 10; i++ {
	// 	x := rand.Intn(10)
	// 	fmt.Println(x)
	// }

	// rad := rand.New(rand.NewSource(time.Now().Unix()))
	// for i := 0; i < rad.Intn(9)+1; i++ {
	// 	fmt.Println(rad.Intn(100))
	// }

	// 
	fmt.Println( GenValidateCode(6)
}
// 6位随机 短信验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
 
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}
 


