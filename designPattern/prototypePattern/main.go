package main

import "fmt"

type Shape struct {
	Id        int
	TypeShape string
}

func (s *Shape) Clone() *Shape {
	res := *s
	return &res

}
func (s *Shape) GetId() int {
	return s.Id
}
func (s *Shape) SetId(id int) {
	s.Id = id
}

func (s *Shape) GetType() string {
	return s.TypeShape
}
func (s *Shape) SetType(shapeType string) {
	s.TypeShape = shapeType
}
func (s *Shape) draw() {
	fmt.Println("Inside::" + s.GetType() + "::draw() method")
}

func main() {
	shape := &Shape{}
	shape.SetType("Circle")
	shape.draw()

	rectangle := shape.Clone()
	rectangle.SetType("Rectangle")
	rectangle.draw()

	square := shape.Clone()
	square.SetType("Square")
	square.draw()
}
