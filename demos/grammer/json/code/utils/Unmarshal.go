package utils

import (
	"fmt"
	"encoding/json"
)



//反序列化struct
func UnMarshalStruct() {

	str := `{"Name":"红孩儿","Age":8,"Skill":"吐火"}`

	//定义一个结构体
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmarshal err=", err)
	}

	fmt.Printf("反序列化struct: monster=%T, monster=%v\n", monster, monster)


}

//反序列化map
func UnMarshalMap() {

	str := `{"Age":500,"Name":"牛魔王","Skill":"牛魔拳~","lover":["铁扇公主","玉面狐狸"]}`

	var myMap map[string]interface{}

	//反序列化，不需要对 myMap 对其make ，因为 json.Unmarshal会make

	err := json.Unmarshal([]byte(str), &myMap)
	if err != nil {
		fmt.Println("unmarshal err=", err)
	}

	fmt.Printf("反序列化map: myMap=%T, myMap=%v\n", myMap, myMap)
}

//反序列化slice
func UnMarshalSlice() {
	
	str := `[{"Age":500,"Name":"牛魔王","Skill":"牛魔拳~","lover":["铁扇公主","玉面狐狸"]},{"Age":300,"Name":"金角大王","Skill":"晃金绳~","lover":["银角大王","铜脚大王"]}]`

	var mySlice []map[string]interface{}

	//反序列化，不需要对 mySlice 对其make ，因为 json.Unmarshal会make

	err := json.Unmarshal([]byte(str), &mySlice)
	if err != nil {
		fmt.Println("unmarshal err=", err)
	}

	fmt.Printf("反序列化slice: mySlice=%T, mySlice=%v\n", mySlice, mySlice)

}




