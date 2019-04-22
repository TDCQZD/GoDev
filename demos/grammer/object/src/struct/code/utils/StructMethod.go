package utils
import (
	_"encoding/json"
	"fmt"
)

/*结构体方法*/
type D struct {
	Num int
}
func (a D) test() {
	fmt.Println(a.Num)
}

func StructMethodDemo4()  {
	//方法的使用
   var a D //创建变量
   a.test() //调用
}
/*
面向对象编程应用步骤

声明(定义)结构体，确定结构体名
编写结构体的字段
编写结构体的方法
*/



type person struct{
	Name string
	Age int
}

func (p person) showInforn()  {
	fmt.Printf("name=%s age=%d\n",p.Name,p.Age)
}

func (p person) speak()  {
	fmt.Printf("%s 是一个好人\n",p.Name)
}

func (p person) jisuan()  {
	sum := 0
	for i := 1; i < 101; i++ {
           sum += i
	}
	fmt.Printf("1+~+100的和为%d \n",sum)
}

func (p person) jisuan2(n int) (sum int) {
	
	for i := 1; i <= n; i++ {
           sum += i
	}
	fmt.Printf("1+~+%d的和为%d \n",n,sum)
	return
}

func (p *person) pointDemo() {
	p.Name="Tom"
	p.Age=100
}
func StructMethodDemo()  {
	p :=person{"张三",18}
	p.showInforn()
	p.speak()
	p.jisuan()
	res :=p.jisuan2(1000)
	fmt.Printf("返回值为 %d \n",res)
	p1 :=&person{"张三",18}
	fmt.Println("------------------------")
	fmt.Printf("指针调用前返回值p= %v \n",*p1)
	p1.pointDemo()//标准调用
	fmt.Printf("指针调用后返回值p= %v \n",*p1)
	
}

/*
Golang中的方法作用在指定的数据类型上的(即：和指定的数据类型绑定)，因此自定义类型，都可以有方法，
而不仅仅是struct， 比如int , float32等都可以有方法
*/
type fl  float64

func (f fl) test()  {
	fmt.Println("f=",f)
}
func (f *fl) test1()  {
	(*f)++
	fmt.Println("f=",f)
	fmt.Println("*f=",*f)
}
func StructMethodDemo1()  {
	var f fl = 12.3425
	f.test()
	f.test1()
}

/*
如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出

*/

type cat struct{
	Name string
	Price float64
}


func (c *cat) String()  string {
	str :=fmt.Sprintf("name=[ %s] price= [%f]",c.Name,c.Price)
	return str
}

func StructMethodDemo2()  {
	c := cat{"劳斯莱斯",1000000000}
	fmt.Println(c)
	res :=	c.String()
	fmt.Printf("res的类型=%T res的值=%v ",res,res)
}



// 编写结构体(MethodUtils)，编程一个方法，方法不需要参数，在方法中打印一个10*8 的矩形，在main方法中调用该方法。
type MethodUtils struct {
       arr [][]int 
}
func (m *MethodUtils) ArrayTranspose(arr *[3][3]int)  {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 1 + i; j < len(arr[i]); j++ {
			temp := arr[i][j]
			arr[i][j] = arr[j][i]
			arr[j][i] = temp
		}
	}
}
func (m MethodUtils) MultipilcationTable(n int){
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ",j,i,j*i)
		}
		fmt.Printf("\n")
	}
}
func StructMethodDemo3()  {
	var m MethodUtils
	m.MultipilcationTable(9)
	
   arrys :=[3][3]int{{1,2,3},{4,5,6},{7,8,9}}
   fmt.Println("arrys=",arrys)
   (&m).ArrayTranspose(&arrys)
   fmt.Println("arrys=",arrys)
}



// 编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形

// 编写一个方法算该矩形的面积(可以接收长len，和宽width)， 将其作为方法返回值。在main方法中调用该方法，接收返回的面积值并打印。

/*
学生案例：
编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型。
结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。
*/

type Student struct {
	name string
	gender string 
	age  int
	id int 
	score float64
}

func (student *Student) say() string {
	str := fmt.Sprintf("name=[%s] gender=[%s]  age=[%d]  id=[%d]  score=[%f] \n",
	student.name,student.gender,student.age,student.id,student.score)
    return str
}
func StudentDemo()  {
	var s Student=Student{
		name : "张三" ,
		gender : "男" ,
		age  : 15,
		id : 001 ,
		score : 88.88,
	}
	(&s).say()
	fmt.Println(s)
	fmt.Println((&s).say())
}
/*
盒子案例

编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
声明一个方法获取立方体的体积。
创建一个Box结构体变量，打印给定尺寸的立方体的体积。
*/
type Box struct {
	hight float64
	width float64
	long float64
}

func (box *Box) volume() float64 {
	return box.hight * box.long * box.width
}

func BoXDemo()  {
	b := Box{12.3,13.4,15.6}
	fmt.Println("立方体的体积=",b.volume())
}