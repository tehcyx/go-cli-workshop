package geometry

import (
	"math"

	"fmt"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func NewRect(width, height float64) Rect {
	return Rect{Width: width, Height: height}
}

func NewCircle(radius float64) Circle {
	return Circle{Radius: radius}
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func Measure(g Geometry) {
	fmt.Printf("%+v\n", g)
	fmt.Printf("Area: %f\n", g.Area())
	fmt.Printf("Perimeter: %f\n", g.Perimeter())
}
