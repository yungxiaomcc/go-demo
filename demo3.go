package main

import "fmt"

//基本类型
//1.整形 int8 uint8

//2.特殊整形
//uint 32位系统上是uint32,64为系统为uint64
//int 32位系统上是int32,64为系统为int64
//uintptr 无符号整型，用于存放一个指针

//3.go 中无法直接定义二进制数

func main() {

	//十进制
	var a int = 10
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	//八进制以0开头
	var b int = 077
	fmt.Printf("%o \n", b)
	fmt.Printf("%d\n", b)
	fmt.Printf("%b\n", b)
	//十六进制 0x
	var c int = 0x7c
	fmt.Printf("%o \n", c)
	fmt.Printf("%d\n", c)
	fmt.Printf("%b\n", c)
	fmt.Printf("%x\n", c)
	// println 默认将int型整数以十进制输出
	fmt.Println(c)
}
