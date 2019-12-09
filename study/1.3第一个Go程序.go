package main   

import "fmt" // 我们需要使用fmt包中的Println()函数 
 
func main() {          
	fmt.Println("Hello, world. 你好，世界！")  
}

/*一个常规的 函数定义包含以下部分： 
func 函数名(参数列表)(返回值列表) {   
	// 函数体 
}  对应的一个实例如下： 
func Compute(value1 int, value2 float64)(result float64, err error) {  
	// 函数体 
}*/


//Go支持多个返回值。以上的示例函数Compute()返回了两个值，一个叫result，另一个是 err。
//并不是所有返回值都必须赋值。在函数返回时没有被明确赋值的返回值都会被设置为默认 值，比如result会被设为0.0，err会被设为nil。 