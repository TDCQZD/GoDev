package utils

import (
	"reflect"
	"fmt"
)

/*反射的最佳实践*/
// 1、使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
/*
分析：
1、结构体 有四个字段 Name Age   Score Sex  
2、声明结构体方法 GetSum  Set Print()
3、TestStruct 处理结构体方法和字段
4、main方法调用
*/

type Monster struct{
	Name string `json:"name"`
	Age int      `json:"age"`
	Score float64 `json:"score"`
	Sex string 

}

func (m Monster)Print()  {
	fmt.Println("---start~----")
	fmt.Println(m)
	fmt.Println("---end~----")
}

func (m Monster)GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m Monster)Set(name string, age int, score float64, sex string) Monster {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
	
	return m
}

func StructOpera(i interface{})  {
	//获取reflect.Type 类型
	rType := reflect.TypeOf(i)
	//获取reflect.Value 类型
	rValue := reflect.ValueOf(i)
	//获取到a对应的类别
	rKind := rValue.Kind()
	//如果传入的不是struct，就退出
    if rKind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	numField := rValue.NumField()
	
	//变量结构体的所有字段

	for i := 0; i < numField; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, rValue.Field(i))
		//获取到struct标签, 注意需要通过reflect.Type来获取tag标签的值
		tagVal := rType.Field(i).Tag.Get("json")
		//如果该字段是tag标签就显示，否则就不显示
		if tagVal !="" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}
	
		
	
	//获取到该结构体有多少个方法
	numMethod := rValue.NumMethod()
	fmt.Printf("struct has %d methods\n", numMethod)
	

	//方法的排序默认是按照 函数名的排序（ASCII码）
	rValue.Method(1).Call(nil) //获取到第二个方法。调用它

    //调用结构体的第1个方法Method(0)
  	//声明了 []reflect.Value
	var params []reflect.Value
	
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(20))
	
	res := rValue.Method(0).Call(params)//传入的参数是 []reflect.Value, 返回[]reflect.Value 
	fmt.Println("res=", res[0].Int()) //返回结果, 返回的结果是 []reflect.Value*/

	var params1 []reflect.Value
	
	params1 = append(params1,reflect.ValueOf("狐狸精"))
	params1 = append(params1,reflect.ValueOf(500))
	params1 = append(params1,reflect.ValueOf(60.0))
	params1 = append(params1,reflect.ValueOf("女"))

	res1 := rValue.Method(2).Call(params1)//传入的参数是 []reflect.Value, 返回[]reflect.Value 
	fmt.Println("res=", res1[0])

}

func StructOperaDemo()  {
	//创建了一个Monster实例
	var m Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	//将Monster实例传递给StructOpera函数
	StructOpera(m)	
	fmt.Println("m=",m)
}
