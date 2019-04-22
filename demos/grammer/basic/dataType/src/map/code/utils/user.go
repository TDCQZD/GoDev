package utils

import (
	"fmt"
)

func UserInfor()  {
	var mapName map[string]map[string]string
    var name string
	mapName = make(map[string]map[string]string)
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
		// fmt.Println("请选择1-5")	
		fmt.Scanln(&num)
		switch num {
		case 1:
			fmt.Println("请输入用户姓名")
			fmt.Scanln(&name)
			addUser(mapName,name)
			fmt.Println("------------------------------")
			fmt.Scanln(&num)
			name = ""	
		case 2:
			fmt.Println("请输入用户姓名")
			fmt.Scanln(&name)
			searchUser(mapName,name)
			fmt.Println("------------------------------")
			fmt.Scanln(&num)
			name = ""	
		case 3:
			fmt.Println("请输入用户姓名")
			fmt.Scanln(&name)
			updateUser(mapName,name)
			fmt.Println("------------------------------")
			fmt.Scanln(&num)	
			name = ""
		case 4:
			searchAllUser(mapName)
			fmt.Println("------------------------------")
		}
		if num == 5{
			break
		}
	}
}
func searchAllUser(user map[string]map[string]string)  {
	if len(user) <= 0 {
		fmt.Println("不存在用户!!!请先添加！")
	}else{
		for key,value := range user {
			fmt.Println("name=",key)
			for key1,val := range value {
				fmt.Printf("\t%v=%v \n",key1,val)
			}
		}
	}

}
func updateUser(user map[string]map[string]string,name string)  {
	if user[name] == nil {
		fmt.Println("该用户不存在!!!请先添加！")
	}else {
		fmt.Println("请输入nickName和pwd")	
		var nickName,pwd  string
		fmt.Scanf("%s %s",&nickName,&pwd)
	    user[name]["nickName"] = nickName
	    user[name]["pwd"] = pwd	
	}
}

func searchUser(user map[string]map[string]string,name string)  {
	if user[name] == nil {
		fmt.Println("该用户不存在!!!请先添加！")
	}else {
		for key,value := range user {
			if key == name {
				for key1,val := range value {
					fmt.Printf("%s = %s\n",key1,val)
				}
			}
		}
		
	}
	
}

func addUser(user map[string]map[string]string,name string)  {
	// fmt.Println("user=",user[name])
	if user[name] == nil {
		user[name] = make(map[string]string)
		 fmt.Println("请输入nickName和pwd")	
		 var nickName,pwd  string
		 
		 fmt.Scanf("%s %s",&nickName,&pwd)
		user[name]["nickName"] = nickName
	    user[name]["pwd"] = pwd
	}else {
		// user[name]["pwd"] = "888"
		fmt.Println("用户已存在!!!")	
	}
	fmt.Println("user=",user)
}