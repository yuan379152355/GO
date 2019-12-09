package main

import "fmt"

//import "math"
func main() {
	var flag bool
	flag = true
	fmt.Println("Result=", flag)

	var fvalue1 float32

	fvalue1 = 12
	fvalue2 := 12.0

	//fvalue1 = fvalue2    错误，fvalue默认为float64，不能直接赋值
	fvalue1 = float32(fvalue2)

	fmt.Println("Result=", fvalue1)

	var f1, f2 float32
	f1 = 1.000001222
	f2 = 1.000001111
	// p为自定义比较精度
	f := func(fv1, fv2, p float32) bool {
		return (fv1 - fv2) > p
	}
	fmt.Println("f1 > f2", f(f1, f2, 0.0000001))

	// 复数  complex
	var valuecomplex1 complex64

	valuecomplex1 = 3.2 + 12i
	//valuecomplex2 := complex(3.2, 12)

	fmt.Println(real(valuecomplex1), imag(valuecomplex1))

	// 字符串
	str := "Hello world" // 初始化后不能被修改
	//str[0] = 'X'           // error: cannot assign to str[0]
	fmt.Println(str[:7])

	// 数组
	array := [10]int{2, 4, 3, 1, 5, 7, 4, 5, 3, 2}
	array[0] = 8
	for _, v := range array { //range关键字
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 数组切片定义
	mySlice1 := make([]int, 5)
	//mySlice2 := make([]int, 5,2)
	mySlice3 := []int{1, 2, 3, 4, 5}
	//array := []int{1,23,,4,5,6,7,8,89,8}
	//mySlice4 := array[:7]
	mySlice1 = append(mySlice1, 1, 2, 3)
	for _, v := range mySlice1 { //range关键字
		fmt.Print(v, " ")
	}
	fmt.Println()
	mySlice1 = append(mySlice1, mySlice3...)
	//我们在第二个参数mySlice3后面加了三个点，即一个省略号，如果没有这个省 略号的话，会有编译错误，因为按append()的语义，从第二个参数起的所有参数都是待附加的 元素。因为mySlice中的元素类型为int，所以直接传递mySlice3是行不通的。加上省略号相 当于把mySlice3包含的所有元素打散后传入
	for _, v := range mySlice1 { //range关键字
		fmt.Print(v, " ")
	}
	fmt.Println()

	Slice1 := []int{1, 2, 3, 4, 5}
	Slice2 := []int{3, 2, 1}
	//copy(Slice1,Slice2)
	//fmt.Println(Slice1)
	copy(Slice2, Slice1)
	fmt.Println(Slice2)
}
