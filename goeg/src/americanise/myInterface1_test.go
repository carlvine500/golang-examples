package main

import (
	"fmt"
	"testing"
)
func init() {
	fmt.Println("init Rectangle")
}
type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func Test23(t *testing.T) {
	var s Shape = Rectangle{10.0, 5.0}
	fmt.Println("Area of rectangle: ", s.area())
}
