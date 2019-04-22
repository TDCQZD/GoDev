package utils

import (
	"fmt"
)

type Person struct{
	Name string
}

func sayOk1(p *Person) {

}

func sayOk2(p Person) {

}

func (p Person) test01() {
	p.Name = "scott"
}

func (p *Person) test02() {
	p.Name = "scott2"   // p.Name 等价  (*p).Name 
	(*p).Name = "scott3" //标准的访问形式
}

//方法和函数区别
func functionAndMethod() {

	var p Person = Person{"tom"}
	//sayOk1(&p) // ok
	//sayOk1(p) //error
	//sayOk2(p) //ok
	//sayOk2(&p) // error

	p.test01()  //ok
	(&p).test01() // ok , 仍然等价于  p.test01(), 从形式上看是使用&p,但是本质仍然是 p
	fmt.Println("main p.Name=", p.Name) // tom √ scott 


	//对于 func (p *Persons) test02() 的调用形式
	(&p).test02() // ok!!!
	p.test02()// ok!! 仍然等价于  (&p).test01(), 从形式上看是使用p,但是本质仍然是 &p
	fmt.Println("main p.Name=", p.Name) // tom , scott2 √

}
