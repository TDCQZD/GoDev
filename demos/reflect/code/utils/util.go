package utils

import (
	"fmt"
	"reflect"
)

func util()  {
	
}

//反射操作方法-基本数据类型float64
func reflectOper(b interface{})  {
	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTye := reflect.TypeOf(b)
	fmt.Println("rtype=",rTye)
	rKind := rTye.Kind()
	fmt.Println("rKind=",rKind)
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=",rVal)
	// 将 reflect.Value 转成 interface{}
	iV := rVal.Interface()
	fmt.Println("iV=",iV)
	//将 interface{} 通过断言转成需要的类型
	num := iV.(float64)
	fmt.Println("num=",num)
}

type Student struct{
	Name string
	Age int

}

//反射操作方法-结构体
func reflectOPerStruct(b interface{})  {
	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	
	fmt.Println("rVal=", rVal)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

func RefelctDemo(){
	// var num float64 = 12.0
	// reflectOper(num)//对普通变量进行反射
	stu :=Student{
		Name : "张三",
		Age : 30,
	}
	reflectOPerStruct(stu)//对结构体进行反射
}


