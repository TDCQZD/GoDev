# interface
在Golang中 多态特性主要是通过接口来体现的。

## 基本介绍
interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量。到某个自定义类型(比如结构体Phone)要使用的时候,在根据具体情况把这些方法写出来(实现)。

## 基本语法
```
type 接口名 interface{

        method1(参数列表) 返回值列表
        method2(参数列表) 返回值列表
        …
}
```
实现接口所有方法：
```
func (t 自定义类型) method1(参数列表) 返回值列表 {
        //方法实现
}
func (t 自定义类型) method2(参数列表) 返回值列表 {
        //方法实现
}
//.... 
```
说明: 
- 1) 接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低偶合的思想。 
- 2) Golang中的接口，不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口。因此，Golang中没有implement这样的关键字

```
type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段Human
	school string
	loan float32
}

type Employee struct {
	Human //匿名字段Human
	company string
	money float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
```
## 空interface
空`interface(interface{})`不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。
```
// 定义a为空接口
var a interface{}
var i int = 5
s := "Hello world"
// a可以存储任意类型的数值
a = i
a = s
```
一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。
## 嵌入interface
源码包container/heap里面有这样的一个定义
```
type Interface interface {
	sort.Interface //嵌入字段sort.Interface
	Push(x interface{}) //a Push method to push elements into the heap
	Pop() interface{} //a Pop elements that pops elements from the heap
}
```
我们看到sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了。也就是下面三个方法：
```
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```
另一个例子就是io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个interface：
```
// io.ReadWriter
type ReadWriter interface {
	Reader
	Writer
}
```