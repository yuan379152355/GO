package main

import (
	"fmt"
)

type Option func(*Options)

type Options struct{
	A int
	B string
}

func SetA(a int) Option {
	return func(o *Options) {
		o.A = a
		fmt.Println("aaa")
	}
}

func SetB(b string) Option {
	return func (o *Options) {
		o.B = b
		fmt.Println("bbb")
	}
}

func InitOptions(opts ...Option) Options {
	//opt := new(Options)
	fmt.Println(111111)
	opt := Options{
		//A: 0,
		//B:"",
	}
	fmt.Println(222222)
	for _, o := range opts {
		o(&opt)
		fmt.Println(123)
	}
    
	return opt
}

func main(){
	opt := InitOptions(
		SetA(4),
		SetB("ysysysys"),
	)

	fmt.Println(opt.A)
	fmt.Println(opt.B)
	//return 0
}