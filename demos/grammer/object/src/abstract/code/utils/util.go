package utils

import (
	_"fmt"
)



type Account struct {
	Account string
	Money float64
	Pwd string
}
/*
func (account *Account) SetAccount (accountName string ) {
	if (len(accountName) > 5 &&  len(accountName) < 11) {
	 account.Account = accountName
	}else{
		fmt.Println("账户不满足6~10位")
		return
	}
	

}

func (account *Account) GetAccount()  string {
	return account.Account

}

func (account *Account) SetPwd (pwd string ) {
	if  len(pwd) == 6 {
		account.Pwd=pwd
	  }else {
		fmt.Println("密码必须是6位数字")
		return
	  }

}
func (account *Account) GetPwd()  string {
	return account.Account

}

func (account *Account) SetMony (money float64 ) {
	if money < 20 {
		fmt.Println("余额必须多于20")
		return
	}else{
		account.Money +=money
	}
	

}
func (account *Account) GetMony()  float64 {
	return account.Money

}

func (account *Account) SaveMoney(money float64 ,pwd string) {
	if account.queryPwd(pwd) {
		account.Money += money
		}else{
			fmt.Println("账户密码错误，请重新输入！")
		}
}

func (account *Account) TakeMoney(money float64 ,pwd string) {
	if account.queryPwd(pwd) {
		if account.Money < money {
			fmt.Println("账户余额不足！")
			return	
		}else {
			account.Money -= money
		}
		
	}else{
		fmt.Println("账户密码错误，请重新输入！")
	}
}

func (account *Account) queryPwd(pwd string) (flag bool) {
	if pwd == account.Pwd {
      flag = true 
	}else {
		flag = false
	}
	return
}

func (account *Account) QueryMoney(pwd string)   {
	if account.queryPwd(pwd) {		
		fmt.Println("账户余额:",account.Money)
	}else {
		fmt.Println("账户密码错误，请重新输入！")		
	}
	
}

func (account *Account) TransformMoney(money float64 ,pwd string)   {

}*/