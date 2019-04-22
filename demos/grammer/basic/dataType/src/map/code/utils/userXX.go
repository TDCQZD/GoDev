package utils

import (
	"fmt"
)

type userXX struct {
	NickName string
	Age int
	Pwd int
}

func UserInfors()  {
	 mapName :=make(map[string]userXX)
    var name string
	fmt.Println("------------菜单--------------")
	fmt.Println("1、添加用户")
	fmt.Println("2、查询用户")
	fmt.Println("3、修改用户")
	fmt.Println("4、显示所有用户")
	fmt.Println("5、退出系统")
	fmt.Println("请选择1-5")
	fmt.Println("------------------------------")
	var num int
	
	for {
		fmt.Scanln(&num)
		switch num {
		case 1:
			fmt.Println("请输入用户姓名")
			fmt.Scanln(&name)
			addUserXX(mapName,name)
			// addUserXXPoint(&mapName,name)
			fmt.Println("------------------------------")
			fmt.Scanln(&num)
			name = ""	
			
		case 2:
			fmt.Println("请输入用户姓名")
			fmt.Scanln(&name)
			searchUserXX(mapName,name)
			fmt.Println("------------------------------")
			// fmt.Scanln(&num)
			name = ""	
			
		case 3:
			fmt.Println("请输入用户姓名")
			fmt.Scanln(&name)
			updateUserXX(mapName,name)
			fmt.Println("------------------------------")
			fmt.Scanln(&num)	
			name = ""
			
		case 4:
			searchAllUserXX(mapName)
			fmt.Println("------------------------------")
		}
		if num == 5{
			break
		}
	}
}
func searchAllUserXX(user map[string]userXX)  {
	if len(user) <= 0 {
		fmt.Println("不存在用户!!!请先添加！")
	}else{
		for key,value := range user {
				// fmt.Printf("%v=%v \n",key,value)	
				fmt.Println("name= ",key)	
				fmt.Println("\t userXX=",value)			
		}
	}

}
func updateUserXX(user map[string]userXX,name string)  {
	if len(user) == 0 {
		fmt.Println("用户不存在!!!请先添加！")
	}else {
		flag :=false
		for key ,_:= range user {
			if key == name {	
				add(user,name)
					
			}else {
				flag = true
			}
		}
		if !flag {
			fmt.Println("该用户不存在!!!请先添加！")
			flag = false
		}
	}

}

func searchUserXX(user map[string]userXX,name string)  {
	if len(user) == 0 {//新增
		fmt.Println("用户不存在!!!请先添加！")
	}else {
		flag :=false
		for key ,value:= range user {
			if key == name {	
				fmt.Println("name= ",key)	
				fmt.Println("\tuserXX=",value)		
			}else {
				flag = true
			}
		}
		if !flag {
			fmt.Println("该用户不存在!!!请先添加！")
			flag = false
		}
	}
	
	
}
func addUserXXPoint(user *map[string]userXX,name string)  {
	
	if len(*user) == 0 {//新增
		add(*user,name)
	}else {//遍历改变密码
		flag :=false
		for key ,value:= range *user {
			if key == name {
				value.Pwd=888	
				(*user)[name] = value			
				flag = true
			}
		}
		if !flag {
			add(*user,name)
			flag = false
		}
	}
	
	fmt.Println("user=",user)
}
func addUserXX(user map[string]userXX,name string)  {
	
	if len(user) == 0 {//新增
		add(user,name)
	}else {//遍历改变密码
		flag :=false
		for key ,value:= range user {
			if key == name {
				value.Pwd=888	
				user[name] = value
				flag = true
			
			}
		}
		if !flag {
			add(user,name)
			flag = false
		}
	}
	
	fmt.Println("user=",user)
}
 func add(user map[string]userXX,name string){
	fmt.Println("请输入nickName、age和pwd")	
	var nickName  string
	var age ,pwd int
	fmt.Scanf("%s %d %d",&nickName,&age,&pwd)
	userInfor :=userXX{nickName,age,pwd}
	user[name] = userInfor
}