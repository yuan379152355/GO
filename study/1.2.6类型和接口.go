package main 

import "fmt"

type Bird1 struct{	
}
func (b *Bird1) Fly() {
	fmt.Println("Bird1 ==> Fly")
}

type Bird2 struct{	
}
func (b *Bird2) Fly() {
	fmt.Println("Bird2 ==> Fly")
}

type Bird3 struct{	
}
func (b *Bird3) Fly() {
	fmt.Println("Bird3 ==> Fly")
}

type Bird4 struct{	
}
func (b *Bird4) Fly() {
	fmt.Println("Bird4 ==> Fly")
}

type IFly interface {
	Fly()
}

func main() {
	var fly IFly = new(Bird1)
	fly.Fly()
	fly = new(Bird2)
	fly.Fly()
	fly = new(Bird3)
	fly.Fly()
	fly = new(Bird4)
	fly.Fly()
}