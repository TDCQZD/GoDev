package main

import "fmt"

type Shape interface {
	Draw()
}
type Red struct{}

func (red *Red) Draw() {
	fmt.Println("Border Color: Red")
}

type White struct{}

func (white *White) Draw() {
	fmt.Println("Center Color: White")
}

type Rectangle struct {
	red   *Red
	white *White
}

func (rectangle *Rectangle) Draw() {
	fmt.Println("Shape: Rectangle")
	rectangle.red.Draw()
	rectangle.white.Draw()
}
func main() {
	shape := &Rectangle{&Red{}, &White{}}
	shape.Draw()
}
