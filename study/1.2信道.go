package main

import (
	"fmt"
)

var channel chan int = make(chan int)

func printString(value string) {
	for i := 0; i < 10; i++ {
		fmt.Println(value)
	}
	channel <- 0
}

func main() {
	go printString("A")
	<-channel // 接受到信号才会退出主线程
}
