package main

import "fmt"

//字符串，原生的go类型
//go 语言中双引号表示字符出啊，单引号包裹字符


/*
字符串中的转义符


*/

func main() {
	a := '杨'
	b := 'a'
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Printf("%T\n", b)
	//打印出来是一个整型,变量中存储的是对应字符的unicode码位
	fmt.Println(b)
    //字符串中转义符
	path := "\"d:\\go\\src\\code\""
	fmt.Println(path)
	//多行字符串, 反引号中的内容原样输出，不需要转义
	s2 := `
	窗前明月光，
	疑是地上霜
	`
	fmt.Println(s2)

}
 