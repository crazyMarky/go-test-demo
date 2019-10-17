package main

import (
	"fmt"
	"unicode"
)

func main() {
	//1、判断字符串中汉字的数量
	// 难点是怎么判断一个字符是一个汉字
	s1 := "hello沙河"

	// 1。依次获得字符串的中文
	var count int
	for _, c := range s1 {
		if unicode.Is(unicode.Han,c) {
			count ++
		}
	}

	fmt.Println(count)


}
