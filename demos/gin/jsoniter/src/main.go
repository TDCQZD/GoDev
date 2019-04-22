package main

import (
	"encoding/json"

	"fmt"

	"github.com/json-iterator/go"
)

type ColorGroup struct {
	ID     int `json:"id"`
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	var group1, group2 ColorGroup
	// 使用encoding/json解析json数据 
	b2, _ := json.Marshal(group)
	fmt.Println(string(b2))
	json.Unmarshal(b2, &group2)
	fmt.Printf("%v \n", group2)
	fmt.Println("---------------jsoniter-----------------------")
	// 使用jsoniter解析json数据 
	b1, _ := jsoniter.Marshal(&group)
	fmt.Println(string(b1))
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(b1, &group1)
	fmt.Printf("%v \n", group1)
}
