package main

import (
	"fmt"
)

type Rect struct {
	x, y          float64
	width, height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

//func NewRect(a, b, c, d float64) *Rect {
//	return &Rect(a, b, c, d)
//}

func main() {
	rect1 := new(Rect)
	rect2 := &Rect{0, 0, 4, 5}
	//rect3 := NewRect(2, 2, 5, 6)
	fmt.Println(rect1.Area())
	fmt.Println(rect2.Area())
	//fmt.Println(rect3.Area())
}
