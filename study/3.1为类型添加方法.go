package main

import (
	"fmt"
)

type Interger int

func (a Interger) Less(b Interger) bool {
	return (a < b)
}

func main() {
	var a Interger = 5
	fmt.Println(a.Less(7))
}
