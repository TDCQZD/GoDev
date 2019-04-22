package utils

import (
	"fmt"
)

var IP =getIP()
func getIP()  string {
	fmt.Println("utils---------getIP---------")
	return "127.0.0.1"
}
func init()  {
	fmt.Println("utils---------init---------")
}
func Test()  {
	fmt.Println("utils---------test---------")
}