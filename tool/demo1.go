package tool

import (
	. "fmt"
	"unicode"
)

// 判断字符串中汉字的数量
// 怎么判断字符是汉字

func FindChineseCharacter(str string) int {

	if len(str) == 0 {
		return 0
	}
	var num int
	for _, c := range str {

		if unicode.Is(unicode.Han, c) {
			num++
		}
	}
	return num
}

func ToolTest() {
	Printf("--------------------")
}
