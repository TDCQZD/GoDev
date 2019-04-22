package utils

import (
	"sort"
	"math/rand"
	"fmt"
)



type camputer struct{
	
}
type phone struct{
	name string
	price float64
}
type camera struct{
	name string
	price float64
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

func (c camputer) working(usb Usb)  {
	usb.boot()
	usb.shutdown()
	usb.charging()
}

func InterfaceDemo()  {
	p := phone{"小米",6000}
	c := camera{"尼康",20000}
	cap := camputer{}
	cap.working(p)
	cap.working(c)

}
/*
实现对Hero结构体切片的排序: sort.Sort(data Interface)
根据age大小排序
分析
实现原理 sort.Sort(data Interface)
 1、定义结构体 heros，并添加字段 name age
 2、自定义 heros切片类型 HeroSlice
 3、HeroSlice实现len方法,确定切片长度
 4、HeroSlice实现less方法,确定切片排序顺序
 5、HeroSlice实现swap方法,进行数据交换操作
 6、调用。
*/
type heros struct{
	name string 
	age int
}

type HeroSlice []heros

func (h HeroSlice)Len() int {

	return len(h)
}
func (h HeroSlice)Less(i,j int) bool {

	return h[i].age < h[j].age
}
func (h HeroSlice)Swap(i,j int)  {

	//传统写法
	t := h[i]
	h[i] = h[j]
	h[j] = t
	//还有一个简洁的写法
	//hs[i], hs[j] = hs[j], hs[i]

}

func SortDemo(){

	var heroSlice  HeroSlice
	for i := 0; i < 10; i++ {
		hero := heros{
			name :  fmt.Sprintf("梁山英雄%d",rand.Intn(100)),
			age : rand.Intn(100),
		}
		heroSlice = append(heroSlice,hero)
	}
	fmt.Println("heroSlice",heroSlice)
	sort.Sort(heroSlice)
	fmt.Println("heroSlice",heroSlice)
}