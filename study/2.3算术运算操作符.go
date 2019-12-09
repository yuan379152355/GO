package main

import (
	"fmt"
)

func main() {
	var x := 3
	y := 2
	u := 1

	//^x
	x %= y
	x &= y
	x >>= u
	fmt.Println(x)
	//fmt.Println(x %= y)
	//fmt.Println(x >>= u)
}
