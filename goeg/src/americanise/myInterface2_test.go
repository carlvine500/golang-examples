package main

import (
	"fmt"
	"testing"
)
func init() {
	fmt.Println("init Triangle")
}
type ColorShape interface {
	Shape
	color() string
}

type Triangle struct {
	width  float64
	height float64
}

func (r Triangle) area() float64 {
	return r.width * r.height/2
}

func (r Triangle) color() string {
	return "red"
}


func Test233(t *testing.T) {
	var s ColorShape = Triangle{10.0, 5.0}
	fmt.Println("Area of Triangle: ", s.area())
	fmt.Println("Area of Triangle: ", s.color())
}
