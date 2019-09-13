package main

import (
	"fmt"
	"strconv"
)

type DrawAPI interface {
	drawCircle()
}
type RedCircle struct {
	Shape
}

func (rc *RedCircle) drawCircle() {
	// fmt.Println("Drawing Circle[ color: red, radius: " + radius + ", x: " + x + ", " + y + "]")
	fmt.Printf("Drawing Circle[ color: red, ")
	rc.draw()
}

type GreenCircle struct {
	Shape
}

func (gc *GreenCircle) drawCircle() {
	// fmt.Println("Drawing Circle[ color: green, radius: " + radius + ", x: " + x + ", " + y + "]")
	fmt.Printf("Drawing Circle[ color: green, radius: ")
	gc.draw()
}

type Shape interface {
	draw()
}

type Circle struct {
	x      int
	y      int
	radius int
	DrawAPI
}

func (c *Circle) draw() {
	fmt.Println("radius: " + strconv.Itoa(c.radius) + ", x: " + strconv.Itoa(c.x) + ", y:" + strconv.Itoa(c.y) + "] ")
}

func main() {
	c := &Circle{x: 100, y: 100, radius: 10}
	red := &RedCircle{Shape: c}
	green := &GreenCircle{Shape: c}

	red.drawCircle()
	green.drawCircle()
}
