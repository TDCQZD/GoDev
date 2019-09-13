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

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{}
}

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

func main() {
	circle := NewShapeFactory().getShape("Circle")
	circle.draw()
	square := NewShapeFactory().getShape("Square")
	square.draw()
	factory := NewShapeFactory().getShape("rectangle")
	if factory != nil {
		factory.draw()
	}
}
