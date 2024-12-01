package main

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init TriangleDoubleFace")
}
type TriangleDoubleFace struct {
	Triangle  // 继承了对象的所有方法和属性
	thickness float64
}

func (r TriangleDoubleFace) area() float64 {
	return r.Triangle.area() * 2 // 使用继承的方法求一面再*2
}

func Test2331(t *testing.T) {
	d := TriangleDoubleFace{Triangle{2.0, 3.0}, 1.0}
	fmt.Println("width=", d.width)
	var s ColorShape = d
	fmt.Println("Area of TriangleDoubleFace: ", s.area())
	fmt.Println("Area of TriangleDoubleFace: ", s.color())
}
