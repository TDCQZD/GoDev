package utils

import (
	"fmt"
)

func MapDemo(){
	 var mapStudent map[string]map[string]string

	 mapStudent = make(map[string]map[string]string)

	 mapStudent["No1"] = make(map[string]string)
	 mapStudent["No1"]["name"] = "张三"
	 mapStudent["No1"]["age"] = "24"
	 mapStudent["No1"]["sex"] = "男"
	 mapStudent["No2"] = make(map[string]string)
	 mapStudent["No2"]["name"] = "张三"
	 mapStudent["No2"]["age"] = "24"
	 mapStudent["No2"]["sex"] = "男"
	 mapStudent["No3"] = make(map[string]string)
	 mapStudent["No3"]["name"] = "张三"
	 mapStudent["No3"]["age"] = "24"
	 mapStudent["No3"]["sex"] = "男"

	fmt.Println("mapStudent=",mapStudent)
	/*map遍历——for range*/

	for key1,value1 := range mapStudent {
		fmt.Println("学号：",key1)
		for key2 ,value2 := range value1 {
			fmt.Printf("\t%v=%v \n",key2,value2)
		}
	}
}
/*map声明*/

/*map增删改查*/


/*map切片*/

/**/

func MapName(user map[string]map[string]string,name string){
	// var mapName map[string]map[string]string

	// mapName = make(map[string]map[string]string)
	if user[name] == nil {
		user[name] = make(map[string]string)
		user[name]["nickName"] = "张三"
	    user[name]["pwd"] = "24"
	}else {
		user[name]["pwd"] = "888"
	}
	fmt.Println("user=",user)
   
}

type stu struct{
	Name string
	Age int
	IsMarrid bool

}
func MapDemo2(){
	var mapStudent map[string]stu

	mapStudent = make(map[string]stu)

	mapStudent["No1"]= stu{"张三",18,false}
	mapStudent["No2"]= stu{"李四",20,true}

	stu1 := stu{"赵柳",20,true}
	stu2 := stu{"周杰",30,true}
	mapStudent["No3"]= stu1
	mapStudent["No4"]= stu2
	fmt.Println("mapStudent=",mapStudent)
}
