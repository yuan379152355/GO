package main

import (
	"fmt"
)
60020
1024           10
32 * 1024      15   
func main() {
	const (
		A3 = iota
		B3 = iota
		C3
	)
	fmt.Println("A3=", A3, "B3=", B3, "C3=", C3)
	const (
		A4 = iota + 10
		B4
		C4
	)
	fmt.Println("A4=", A4, "B4=", B4, "C4=", C4)

	type BitFlag int
	const (
		Active BitFlag = 1 << iota // 移位
		Send
		Recive
	)
	flag := Active | Send // 位或
	fmt.Println("flag=", flag)
}
