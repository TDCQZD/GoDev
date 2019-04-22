package main

import (
	"fmt"
	"go_code/grammar/function/init/code/utils"
)
var IP =getIP()
func getIP()  string {
	fmt.Println("main---------getIP")
	return "127.0.0.1"
}
func init()  {
	fmt.Println("main---------init")
}
func main()  {
	fmt.Println("main1---------",utils.IP)
	utils.Test()
	fmt.Println("main---------main")
	fmt.Println("main---------IP",IP)
	fmt.Println("main---------",utils.IP)
	utils.Test()
}