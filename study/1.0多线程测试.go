package main

import (
	"fmt"
	"time"
)

func printString(value string) {
	for i := 0; i < 10; i++ {
		fmt.Println(value)
		time.Sleep(time.Second) //0.1 second
	}
}

func main() {
	go printString("A")
	go printString("B")
	time.Sleep(time.Second * 10) //暂时挂起主线程，让goroutine的两个线程跑完
}
