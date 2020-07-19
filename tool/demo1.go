package tool

import "unicode"
// 判断字符串中汉字的数量
// 怎么判断字符是汉字

func FindChineseCharacter(str string) num int {

	if len(str) == 0{
		return 0
	}

	for _,c:= range str {
		var num int
		if unicode.Is(unicode.Han, c){
			num++
		}
	}
	return 0
}