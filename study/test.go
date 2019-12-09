package main

import (
	"fmt"
	"regexp"
)

func main() {
	//a := "I am learning Go language"
	var str string = "bbbbb.aaaa;2;mmm,nnn_&1_11,12_&2_21,22_&3_31,32_&end"
	fmt.Println([]byte(str))
	//re, _ := regexp.Compile("aaaa(.*?)_&end")
	re, _ := regexp.Compile("aaaa.*?_&end")
	//查找符合正则的第一个
	one := re.Find([]byte(str))
	fmt.Println("Find:", string(one))

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAll([]byte(str), -1)
	for _, value := range all {
		fmt.Println("FindAll", string(value))
	}
}
