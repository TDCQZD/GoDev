package utils

import (
	"fmt"
	"encoding/json"
)


//演示一个结构体序列化的案例
func TestStruct() {

	//先创建要给结构体实例
	monster1 :=  Monster{
		Name : "红孩儿~",
		Age : 8,
		Skill : "吐火",
	}
	var m []Monster
	m = append(m,monster1)
	fmt.Println("源串 = ",&monster1)
	//序列化
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	fmt.Println("序列化的字符串 = ",string(data))

}

//演示map序列化的案例
func TestMap() {

	var myMap map[string]interface{} // map的key是string而值 interface{}
	myMap = make(map[string]interface{}, 1)
	myMap["Name"] = "牛魔王"
	myMap["Age"] = 500
	myMap["lover"] = [2]string{"铁扇公主", "玉面狐狸"}
	myMap["Skill"] = "牛魔拳~"

	//序列化
	data, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	fmt.Println("序列化的字符串 = ",string(data))


}

//演示map序列化的案例
func TestSlice(){

	var mySlice []map[string]interface{}

	var myMap map[string]interface{} // map的key是string而值 interface{}
	myMap = make(map[string]interface{}, 1)
	myMap["Name"] = "牛魔王"
	myMap["Age"] = 500
	myMap["lover"] = [2]string{"铁扇公主", "玉面狐狸"}
	myMap["Skill"] = "牛魔拳~"

	mySlice = append(mySlice, myMap)

	var myMap2 map[string]interface{} // map的key是string而值 interface{}
	myMap2 = make(map[string]interface{}, 1)
	myMap2["Name"] = "金角大王"
	myMap2["Age"] = 300
	myMap2["lover"] = [2]string{"银角大王", "铜脚大王"}
	myMap2["Skill"] = "晃金绳~"

	mySlice = append(mySlice, myMap2)

	//序列化
	data, err := json.Marshal(mySlice)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	fmt.Println("序列化的字符串 = ",string(data))


}

//演示Float序列化的案例
func TestFloat() {
	var num1 float64 = 908.12

	
	//序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
		return
	}
	fmt.Println("序列化的字符串 = ",string(data))

}


