package utils

import (
	"fmt"
)

type student struct{}

type action interface{
	sport()
	
}
func (s student) sport()  {
	
}
// 1、接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量(实例)
func DetailsDemo1(){
 
		var a action
		fmt.Println("a=",a)
		var s student
		a=s
		fmt.Println("s=",s)
}
// 2、接口中所有的方法都没有方法体,即都是没有实现的方法。
// 3、在Golang中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口。
type student1 struct{}

type action1 interface{
	// a int
	sport()
	draw()
}
func (s student1) sport()  {
	
}
func DetailsDemo2(){
	
	var a action1
	fmt.Println("a=",a)
	var s student1
	// 如果没有将接口的所有方法实现，会出现恐慌
	// a=s
	fmt.Println("s=",s)
}
// 4、一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型。 
func DetailsDemo3(){
		var a action1
		fmt.Println("a=",a)
		var s student
		// a=s //stuent没有实现action1下方法，会出现恐慌
		fmt.Println("s=",s)
}
// 5、只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
type integer int
func (i integer) sport()  {
	
}
func DetailsDemo4(){
	    var a action
		fmt.Println("a=",a)
		var i integer
		a=i
		fmt.Println("a=",a)
}
// 6、一个自定义类型可以实现多个接口 
func DetailsDemo5(){}
// 7、Golang接口中不能有任何变量
type action2 interface{
	sport()
	// age integer
}
// 8、一个接口(比如A接口)可以继承多个别的接口(比如B,C接口)，这时如果要实现A接口，也必须将B,C接口的方法也全部实现。
type AInter interface {
	sayOk1()
}

type BInter interface {
	sayOk2()
	// sayOk1() //继承多个接口时，每个接口的方法不能重复
}

type CInter interface {
	AInter
	BInter
	sayOk3()
}

type Monster struct {

}

//让Monster 实现 CInter
func (m Monster) sayOk1() {
	fmt.Println("sayOk1()")
}

func (m Monster) sayOk2() {
	fmt.Println("sayOk2()")
}

func (m Monster) sayOk3() {
	fmt.Println("sayOk3()")
}

func DetailsDemo6(){	

	var monster Monster
	var c CInter
	c = monster
	fmt.Println("c=", c)

	var bInter BInter
	bInter = monster
	fmt.Println("bInter=", bInter)	
}
// 9、interface类型默认是一个指针(引用类型)，如果没有对interface初始化就使用，那么会输出nil 
func DetailsDemo7(){
	var c CInter
	fmt.Println("c=", c)
}
// 10、空接口interface{} 没有任何方法，所以所有类型都实现了空接口
func DetailsDemo8(){

	var m interface {}
	num1 := 10
	num2 := 12.00
	num3 := true
	num4 := "sjdskj"
	//空接口interface{}可以接受任何类型
	m = num1
	fmt.Println(m)
	m = num2
	fmt.Println(m)
	m = num3
	fmt.Println(m)
	m = num4
	fmt.Println(m)
}