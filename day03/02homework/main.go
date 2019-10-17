package main

import (
	"fmt"
)

func main() {
	//1、判断字符串中汉字的数量
	// 难点是怎么判断一个字符是一个汉字
	//s1 := "hello沙河"
	//
	//// 1。依次获得字符串的中文
	//var count int
	//for _, c := range s1 {
	//	if unicode.Is(unicode.Han,c) {
	//		count ++
	//	}
	//}

	//fmt.Println(count)

	//回文判断
	// 字符串从右和从左都是一样的
	ss := "a山西运煤车煤运西山a"
	//
	r := make([]rune,0,len(ss))
	for _, c := range ss {
		r = append(r,c)
	}
	for i := range r {
		if  r[i] != r[len(r)-1-i] {
			fmt.Println(r,"不是回文")
			return
		}
		if i==len(r)/2 {
			fmt.Println(r,"是回文")
		}
	}

}
