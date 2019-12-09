package main

import "fmt"

func main() {
	i, j := 10, 11
	// 匿名函数
	f := func(m, n int) int {
		return m + n
	}
	fmt.Println(f(i, j))
}
