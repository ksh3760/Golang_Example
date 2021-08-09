package main

import (
	"fmt"
	"math"
)

func main() {
	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)
}

// Interface 정의
type Shape interface {
	area() float64
	perimeter() float64
}

// Rect 정의
type Rect struct {
	width, height float64
}

// Circle 정의
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printIt(v interface{}) {
	fmt.Println(v) // Tom
}
