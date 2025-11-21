package structsinterfaces

import "math"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// 1. Define a Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Declare structs
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// 3. Constructors
func NewCircle(radius float64) Circle {
	return Circle{Radius: radius}
}

func NewRectangle(width, height float64) Rectangle {
	return Rectangle{Width: width, Height: height}
}

// 4. Implement interface Shape for Circle
// Area = pi * r^2
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter = 2 * pi * r
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 4. Implement interface Shape for Rectangle
// Area = w * h
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter = 2 * (w + h)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
