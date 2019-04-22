package utils

import (
	"encoding/json"
	"fmt"
)

/*结构体*/
type cats struct{
    name string
    age int
	price float64
	loves []string
}


func StructDemo()  {
	// 在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值(默认值)，
	var c cats
	fmt.Println("c=",c)
	fmt.Println("c.voves=",c.loves)
	// 不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，不影响另外一个
    c = cats{
		name : "白猫",
		age : 2,
		price : 1000,
	}
	fmt.Println("c=",c)
	c1 := cats{
		name : "黑猫",
		age : 2,
		price : 1000,
	}
	fmt.Println("c1=",c1)
	c1.name = "花猫"
	fmt.Println("c1=",c1)

}
/*结构体类型转换*/
type A struct{
	name string
}
type B struct{
	name string
}
type C A
func StructDemo1(){
	//结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名字、个数和类型)
	var a A
	var b B
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	b = B(a)
	fmt.Println("b=",b)
	
	// 结构体进行type重新定义(相当于取别名)，Golang认为是新的数据类型，但是相互间可以强转
	// var c C
	// c = a
}

type Students struct {
	Name string `json:"myName"` // 给strut的字段，指定一个tag标签
	Age int	`json:"myAge"`
}
func StructDemo2(){	
	//演示struct的每个字段上，可以写上一个tag, 
	//该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
	stu1 := Students{"二狗子", 20}
	stu1Str, err := json.Marshal(stu1) // 将stu1序列化成一个字符串
	if err == nil {
		fmt.Println("stu1Str=", string(stu1Str))
	}else {
		fmt.Println("序列化错误=", err)
	}

}
/*继承和接口的区别*/
//定义猴子结构体
type Monkey struct {
	Name string
}
func (m Monkey) climbing() {//方法
	fmt.Println(m.Name, " 生来会爬树...")
}


//定义一个接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

type LittleMonkey struct {//继承
	Monkey
}

func (lm LittleMonkey) Flying(){
	fmt.Println(lm.Name, " 通过学习会飞翔...")
}

func (lm LittleMonkey) Swimming(){
	fmt.Println(lm.Name, " 通过学习会游泳...")
}

func ExtendAndInterface(){	
	littleMonkey := LittleMonkey{
		Monkey{"悟空"},
	}
	littleMonkey.climbing()//继承
	littleMonkey.Flying()//接口
	littleMonkey.Swimming()//接口
}
