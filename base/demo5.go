package base

import "fmt"

//字符串，原生的go类型
/*
字符串中包含的转义符

\r
\n
\'
\"
\\

*/

func demo5() {
	a := '杨'
	b := 'a'
	fmt.Printf("%T\n", a)
	fmt.Print(a)
	fmt.Printf("%T\n", b)
	//打印出来是一个整型,变量中存储的是对应字符的unicode码位
	fmt.Print(b)

}
