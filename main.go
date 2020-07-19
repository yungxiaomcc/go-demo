package main

import (
	"fmt"

	"github.com/yungxiaomcc/go-demo/tool"
)

func main() {
	str := "杨晓"
	num := tool.FindChineseCharacter(str)
	fmt.Println(num)

	tool.ToolTest()
}
