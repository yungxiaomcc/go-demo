package main

import "fmt"

//浮点数

func main() {
	//math.MaxFloat32 //最大float32 值
	f1 := 3.12
	//go 语言中默认小数都是float64
	fmt.Printf("%T\n", f1) //foat64,
	//定义float32,需要显式定义
	f2 := float32(3.12312)
	fmt.Printf("%T", f2)
	//float32 和我float64之间不能相互赋值
	// f1 = f2

}
