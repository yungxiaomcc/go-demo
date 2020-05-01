package main

import (
	"fmt"
)

//函数外只能声明语句

//单独声明
var name string

//批量声明
var (
	age  int
	addr string
	isOk bool
)

func main() {

	//go 中非全局变量声明后必须使用，否则编译不通过
	var car string = "jili"
	//类型推导,不用声明类型
	// var car2 = "jili"
	//简短变量声明,只能在函数中使用
	// car3 := "jili"
	//匿名变量 _ 用于忽略值
	fmt.Printf("car is %s", car)

}
