package main

import (
	"fmt"
	"strconv"
)

type Packing interface {
	Pack(string)
}

type Item struct {
	Name  string
	Price float64
}

func (item *Item) SetName(name string) {
	item.Name = name
}
func (item *Item) SetPrice(price float64) {
	item.Price = price
}
func (item *Item) GetName() string {
	return item.Name
}
func (item *Item) GetPrice() float64 {
	return item.Price
}
func (item *Item) Pack(str string) {
	fmt.Printf("Packing :" + str + ",")
}

type Meal interface {
	SetName(name string) Meal
	SetPrice(price float64) Meal
	Meal() *Item
}

type MealItem struct {
	item *Item
}

func (mb *MealItem) SetName(name string) Meal {
	if mb.item == nil {
		mb.item = &Item{}
	}
	mb.item.SetName(name)
	return mb
}
func (mb *MealItem) SetPrice(price float64) Meal {
	if mb.item == nil {
		mb.item = &Item{}
	}
	mb.item.SetPrice(price)
	return mb
}
func (mb *MealItem) Meal() *Item {
	return mb.item
}

/*
type MealBuilder struct {
	items []*Item
}

func (mb *MealBuilder) ShowItems() {
	if len(mb.items) < 1 {
		return
	}
	for i := 0; i < len(mb.items); i++ {
		name := mb.items[i].GetName()
		price := mb.items[i].GetPrice()
		fmt.Printf("Name: " + name + "," + "Price: " + strconv.FormatFloat(price, 'f', -1, 64) + ";")
	}
}
func (mb *MealBuilder) GetCost() {
	if len(mb.items) < 1 {
		return
	}
	var costs float64
	for i := 0; i < len(mb.items); i++ {
		costs += mb.items[i].GetPrice()
	}
	fmt.Println("Total Cost:" + strconv.FormatFloat(costs, 'f', -1, 64))
}
*/
type Director struct {
	meal Meal
}

func (p Director) CreateVegMeal(name string, price float64) *Item {
	return p.meal.SetName(name).SetPrice(price).Meal()
}
func main() {

	var meal Meal = &MealItem{}
	var director *Director = &Director{meal: meal}
	var item1 *Item = director.CreateVegMeal("Veg Burge", 25.0)
	item1.Pack("Wrapper")
	fmt.Println(item1.GetName() + "," + strconv.FormatFloat(item1.GetPrice(), 'f', -1, 64))

}
