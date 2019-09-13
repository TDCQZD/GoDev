package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct{}

func (c *Circle) draw() {
	fmt.Println("Circle::draw()")
}

type Square struct{}

func (s *Square) draw() {
	fmt.Println("Square::draw()")
}

type ShapeFactory struct{}

func (sf *ShapeFactory) getShape(shapeType string) Shape {
	switch shapeType {
	case "Circle":
		return &Circle{}
	case "Square":
		return &Square{}
	default:
		fmt.Println("Error :: Input shape type is not exist!")
		return nil

	}
}
func (sf *ShapeFactory) getColor(colorType string) Color {
	return nil
}

type Color interface {
	fill()
}
type Red struct{}

func (red *Red) fill() {
	fmt.Println("Inside Red::fill() method")
}

type Blue struct{}

func (blue *Blue) fill() {
	fmt.Println("Inside Blue::fill() method")
}

type ColorFactory struct{}

func (cf *ColorFactory) getColor(colorType string) Color {
	switch colorType {
	case "Red":
		return &Red{}
	case "Blue":
		return &Blue{}
	default:
		fmt.Println("Error :: Input color type is not exist!")
		return nil

	}
}
func (cf *ColorFactory) getShape(shapeType string) Shape {
	return nil
}

type AbstractFactory interface {
	getShape(string) Shape
	getColor(string) Color
}

func FactoryProducer(factoryType string) AbstractFactory {
	switch factoryType {
	case "Shape":
		return &ShapeFactory{}
	case "Color":
		return &ColorFactory{}
	default:
		fmt.Println("Error :: Input factory type is not exist!")
		return nil
	}
}
func main() {
	as := FactoryProducer("Shape")
	as.getShape("Circle").draw()
	as.getShape("Square").draw()
	fmt.Println("----------------------------------")
	as = FactoryProducer("Color")
	as.getColor("Blue").fill()
	as.getColor("Red").fill()

}
