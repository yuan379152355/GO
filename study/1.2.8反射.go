package main 
 
import (
	"fmt"
	"reflect"
) 
 
type Bird struct {
	Name string     
	LifeExpectance int
}
func (b *Bird) Fly() {
	fmt.Println("Bird==>Fly")
}

//反射（reflection）是在Java语言出现后迅速流行起来的一种概念。
//通过反射，你可以获取对 象类型的详细信息，并可动态操作对象。
//反射是把双刃剑，功能强大但代码可读性并不理想。若非必要，我们并不推荐使用反射
func main() {
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f:= s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

// 第9章中简要介绍反射的基本使用方法和注意事项