package utils

import (
	"fmt"
)


type phone struct{
	name string
	
}
type camera struct{
	name string

}
type camputer struct{
	
}
func (c camputer) working(usb Usb)  {
	usb.boot()
	usb.shutdown()
	usb.charging()
}
type Usb interface {
	charging()
	boot()
	shutdown()
}
//实现接口
func (p phone) charging(){
	fmt.Printf("%s手机开始充电。。。。\n",p.name)
}
func (p phone) boot(){
	fmt.Printf("%s手机正在开机。。。。\n",p.name)
}

func (p phone) shutdown(){
	fmt.Printf("%s手机关机了。。。。\n",p.name)
}
func (c camera) charging(){
	fmt.Printf("%s相机开始充电。。。。\n",c.name)
}
func (c camera) boot(){
	fmt.Printf("%s相机正在开机。。。。\n",c.name)
}
func (c camera) shutdown(){
	fmt.Printf("%s相机关机了。。。。\n",c.name)
}




//多态参数

func PloymDemo()  {
	p := phone{"小米"}
	c := camera{"尼康"}
	cap := camputer{}
	cap.working(p)
	cap.working(c)
}

//多态数组
func PloymDemo1()  {
	var usbArr [4]Usb
	usbArr[0] = phone{"vivo"}
	usbArr[1] = camera{"尼康"}
	usbArr[2] = phone{"小米"}
	usbArr[3] = camera{"索尼"}
	fmt.Println("usbArr=", usbArr)
	//遍历usbArr,调用方法
	for i := 0; i < len(usbArr); i++ {
		usbArr[i].charging()
		usbArr[i].boot()
		usbArr[i].shutdown()

	}

}

//类型断言
func TypeAssertion()  {
	var t float32 = 1.2
	var x interface{}
	x = t //ok
	
	//待检测的类型断言
	y , res := x.(float32)
	if res {
		fmt.Println("断言成功!", y)
	} else {
		fmt.Println("断言失败!")
	}

	y1 ,res1 := x.(int32)//转成float
	if res1 {
		fmt.Println("断言成功!", y1)
	} else {
		fmt.Println("断言失败!")
	}

}

/*最佳实践1*/
//类型断言使用
func (p phone) call(){
	fmt.Printf("%s手机正在通话。。。。\n",p.name)
}
func (c camputer) working1(usb Usb)  {
	// 开始使用类型断言来调用Phone特有的方法Call
	ps ,yes :=usb.(phone)
	if yes {
		ps.call()
	}
	usb.boot()
	usb.shutdown()
	usb.charging()
}
func TypeAssertion1()  {	
	p := phone{"小米"}
	c := camera{"尼康"}
	cap := camputer{}
	cap.working1(p)
	cap.working1(c)
}
/*最佳实践2*/
func TypeAssertion2()  {
	var t interface{}  // 这个是一个空接口
	var n int32 = 80 //
	t = n // 将  n 赋值给 t

	//写一代码，判断 t究竟指向什么类型 type-switch 又叫 类型断言
	// i := t.(type)  有两个作用
	// (1) 先尝试将 t 转成实际指向的数据类型
	// (2) 如果转成功，就会将实际指向的值，赋给 i
	// (3)
	switch i := t.(type) { // 
		case float32 :
			fmt.Printf("i的类型%T i的值%v\n", i, i)
		case float64 :
			fmt.Printf("i的类型%T i的值%v\n", i, i)
		case int32 :
			fmt.Printf("i的类型%T i的值%v\n", i, i)
		case bool :
			fmt.Printf("i的类型%T i的值%v\n", i, i)
		case string :
			fmt.Printf("i的类型%T i的值%v\n", i, i)
       //case Phone: //...
		default :
			fmt.Println("类型不确定")
	}

}