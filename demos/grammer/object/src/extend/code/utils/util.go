package utils

import (
	"fmt"
)
//继承快速入门案例
type student struct {
	name string
	age int
	score float64
}
func (s *student) setGrade(score float64)  {
	s.score=score
}  
func (s *student) showInfor()  {
	fmt.Printf("学生信息如下：姓名：%s 年龄：%d  分数:%f \n",
				s.name,s.age,s.score)
}

type collegeStudent struct{
	student
	writeThesis bool
}
type middleStudent struct{
	student
	college bool
}

type middleStudentPoint struct{
	*student
	college bool
}

func (c *collegeStudent) graduation()  {
	res := "未毕业"
	if c.writeThesis {
		res = "已毕业"
	}

	fmt.Printf("毕业信息如下：姓名：%s 年龄：%d  分数:%f 是否毕业: %s\n",
		c.name,c.age,c.score,res)
}

func (p *primaryStudent) showInfor()  {
	fmt.Printf("学生信息如下：姓名：%s 年龄：%d  分数:%f \n",
	p.name,p.age,p.score)
}
func ExtendDemo()  {
	// 匿名结构体字段访问可以简化
	/*继承结构体的声明方式1*/
	var college collegeStudent
	college.score = 70.5
	college.student.name = "张三"
	college.age = 23
	college.showInfor()
	college.student.setGrade(59.1)
	college.showInfor()
	college.writeThesis = true
	college.graduation()
	/*继承结构体(匿名结构体)的声明方式2*/
	middle := middleStudent{
		student : student{
			name : "王辉",
			age : 44,
			score : 50,
		},
		college :true,
	}
	/*继承结构体的声明方式3*/
	middle1 := middleStudent{
		 student{ "王辉",44,90,},true,
	}

	middle.showInfor()
	middle1.showInfor()
	/*指针*/
  /*继承结构体的声明方式4*/
  middlep :=middleStudentPoint{&student{ "刘亦菲",44,90,},true,}
  /*继承结构体的声明方式5*/
  middlep1 :=middleStudentPoint{
	 	 student : &student{
				name : "小龙女",
				age : 44,
				score : 50,
			},
		college :true,}
		middlep.showInfor()
		middlep1.showInfor()
}
/*
如果一个struct嵌套了一个有名结构体，这种模式就是组合，
如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字
*/

type romm struct{//组合结构体
	 s student
	  num int
}

func ExtendDemo1()  {
	/*组合结构体声明方式1*/
	var romms romm
	// romms.name ="ds" //访问方式是错误的
	 romms.s.name = "小米粒"
	 romms.s.age = 34
	 romms.s.score = 60
	 romms.num = 30
	 fmt.Println(romms)
	/*组合结构体声明方式2*/
	romms2 := romm{
		s : student{"大米粒",23,80},
		num : 40,
	}
	fmt.Println(romms2)

	/*组合结构体声明方式3*/
	romms3 := romm{student{"大米粒",23,80},40}
	fmt.Println(romms3)

	
}
/*结构体声明方式4种*/
func ExtendDemo4()  {
	/*结构体声明方式1-直接声明*/	
	var s student
	s.name = "赵柳"
	s.age = 16
	s.score = 90
	s.showInfor()
	/*结构体声明方式2-字面量*/	
	s1 := student{"黄晓明",18,45}
	s2 := student{name : "赵丽颖",age : 18, score : 30}
	s1.showInfor()
	s2.showInfor()
   /*结构体声明方式3-new关键字*/	
	var s3 *student = new (student)
	//标准的赋值方法
	//因为上面的写法通过指针去访问字段的标准方式，比较麻烦，编译器的底层做了优化
	// (*s3).name  等价于  s3.name 
	// 即：编译器，会自动带 s3.name  ，执行时，加上 * ,就是 (*s3).name
	(*s3).name = "刘备"//标准用法
	s3.age = 32 //简写用法
	s3.score = 90
	s3.showInfor()
	/*结构体声明方式3-指针*/
	var s4 *student = &student{"张宇",18,45}
	s4.showInfor()
}

/*
当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，
如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
*/

type primaryStudent struct{
	student
	age int
}



func ExtendDemo2()  {

primary := primaryStudent{
	student : student{
		name : "小龙",
		age : 9,
		score : 50,
	},
	age :13,
}
primary.student.showInfor()
primary.showInfor()
}
/*
	结构体嵌入两个(或多个)匿名结构体，如两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，
	在访问时，就必须明确指定匿名结构体名字，否则编译报错。
	*/

	type person struct {
		name string
		age int
		score float64
	}
	type user struct {
		person
		student
		score float64
	}
	// 如嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型名来区分。
func ExtendDemo3()  {
	var root user
	// root.age = 23
	root.person.age = 23
	root.score = 23
	root.person.score = 23
	root.student.score = 23
}
	
/*
6) 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值，并取值显示
*/



type Good struct{
	Name string
	Price float64
}
type Band struct{
	Name string
	Address string
}

type Test struct{
	Band
	Good
}
func ExtendDemo5()  {
	var t Test
	t.Band.Name = "zsff"
	t.Good.Name = "qrqf"
}