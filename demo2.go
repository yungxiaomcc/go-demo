package main

import "fmt"

//常量定义无法修改，程序运行期间不修改
const pi = 3.1412

//批量
const (
	statusOk = 200
	notFound = 404
)

//
const (
	n1 = 200 //200
	n2       //200
	n3       //200
)

//
const (
	k1 = iota //0
	k2        //1
	k3        //2
	k4        //3
)

// 同上
const (
	t1 = iota //0
	t2 = iota //1
	t3 = iota //2
	t4 = iota //3
)

//
const (
	p1 = iota //0
	_
	p2 //1
	p3 //2
	p4 //3
)

//插队
const (
	j1 = iota //0
	j2 = 100  //100
	j3        //100
	j4        //100
)

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

//多常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //1,2
	d3, d4 = iota + 1, iota + 2 //2,3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//iota 常量计数器，只能用在常量表达式中
//iota在const关键字中出现将被重置为0，const中每新增一行常量
//声明 将使iota计数一次，iota可理解为const语句块中的行级索引
//
//
func main() {
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)

}
