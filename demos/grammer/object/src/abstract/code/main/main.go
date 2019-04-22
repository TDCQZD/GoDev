package main

import (
	"fmt"
	"go_code/grammar/OOP/abstract/code/utils"
)

func main()  {
	

  account1 := utils.CreateAccount("123456789",20,"000000")
//   account2 := utils.CreateAccount("1234567890",200000,"000000")
  fmt.Println(account1)
//   fmt.Println(account2)
//   account1.SaveMoney(200,"000000")
//   fmt.Println(account1)
//   account1.QueryMoney("000000")
//   account1.TakeMoney(200,"000000")
//   fmt.Println(account1)
//   account1.UpdatePwd("000000","000001")
//   fmt.Println(account1)
//    account2.TransformMoney(account1,1000,"000000")
//    fmt.Println(account1)
//   fmt.Println(account2)
    
	 res1 := account1.GetAccount()
	 res2 := account1.GetMony()
	 res3 := account1.GetPwd()
	 fmt.Println(res1)
	 fmt.Println(res2)
	 fmt.Println(res3)

	 account1.SetAccount("1111111")
	 fmt.Println(account1)
	 account1.SetMony(1111111)
	 fmt.Println(account1)
	 account1.SetPwd("111111")
	 fmt.Println(account1)

}