package main

import (
	"fmt"
)

func main() {
	const A = 0
	const B = 1
	const C = 2
	fmt.Println("A=", A, "B=", B, "C=", C)
	const (
		A1 = 0
		B1 = 1
		C1 = 2
	)
	fmt.Println("A1=", A1, "B1=", B1, "C1=", C1)
	const (
		A2 = 0
		B2
		C2
	)
	fmt.Println("A2=", A2, "B2=", B2, "C2=", C2)

	const factor = 3
	i := 20000
	i *= factor
	j := int16(20)
	i += int(j)
	k := uint8(0)
	k = uint8(i)
	fmt.Println(i, j, k)
}
