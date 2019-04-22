package utils

import (
	"fmt"
)



type accountXX struct {
	account string
	money float64
	pwd string
}



func CreateAccount(accountName string,Money float64 ,Pwd string) *accountXX{
	
	if verParater(accountName ,Money ,Pwd) {
		return &accountXX{
			account : accountName,
			money :Money,
			pwd : Pwd ,
		}
	}else {
		return nil
	}
	
}

func verParater(accountName string,Money float64 ,Pwd string) (flag bool) {

	if len(accountName) < 6 ||  len(accountName) > 10 {

		fmt.Println("账户不满足6~10位")
		flag = false		
		return 
	}
	if  len(Pwd) != 6 {
	
		fmt.Println("密码必须是6位数字")
		flag = false	
		return 	 
	  }
	if Money < 20 {
	
		fmt.Println("余额必须多于20")
		flag = false
		return 		
	}
	flag = true
	return 
}

func (account *accountXX) SetAccount (accountName string ) {
	if len(accountName) > 5 &&  len(accountName) < 11 {
	 account.account = accountName
	}else{
		fmt.Println("账户不满足6~10位")
		return
	}
}
func (account *accountXX) GetAccount()  string {
	return account.account

}

func (account *accountXX) SetPwd(pwd string ) {
	if  len(pwd) == 6 {
		account.pwd=pwd
	  }else {
		fmt.Println("密码必须是6位数字")
		return
	  }

}
func (account *accountXX) GetPwd()  string {
	return account.pwd

}

func (account *accountXX) SetMony(money float64 ) {
	if money < 20 {
		fmt.Println("余额必须多于20")
		return
	}else{
		account.money +=money
	}
	

}
func (account *accountXX) GetMony()  float64 {
	return account.money

}

func (account *accountXX) SaveMoney(money float64 ,pwd string) {
	if account.queryPwd(pwd) {
		account.money += money
		}else{
			fmt.Println("账户密码错误，请重新输入！")
		}
}

func (account *accountXX) TakeMoney(money float64 ,pwd string) {
	if account.queryPwd(pwd) {
		if account.money < money {
			fmt.Println("账户余额不足！")
			return	
		}else {
			account.money -= money
		}
		
	}else{
		fmt.Println("账户密码错误，请重新输入！")
	}
}

func (account *accountXX) queryPwd(pwd string) (flag bool) {
	if pwd == account.pwd {
      flag = true 
	}else {
		flag = false
	}
	return
}

func (account *accountXX) QueryMoney(pwd string)   {
	if account.queryPwd(pwd) {		
		fmt.Println("账户余额:",account.money)
	}else {
		fmt.Println("账户密码错误，请重新输入！")		
	}
	
}

func (account *accountXX) TransformMoney(tfAccount *accountXX,money float64 ,pwd string) {
	if account.queryPwd(pwd) {
		if account.money < money {
			fmt.Println("账户余额不足！")
			return	
		}else {
			account.money -= money
			tfAccount.money +=money
		}
		
	}else{
		fmt.Println("账户密码错误，请重新输入！")
	}
}
	func (account *accountXX) UpdatePwd(oldPwd string,newPwd string) {
		if account.queryPwd(oldPwd) {
			if oldPwd == newPwd {
				fmt.Println("新旧密码不能重复！")
				return	
			}else if  len(newPwd) != 6 {
				fmt.Println("密码必须是6位数字")	
				return 	 
			  }else {
				account.pwd=newPwd
			}
			
		}else{
			fmt.Println("账户密码错误，请重新输入！")
		}
}