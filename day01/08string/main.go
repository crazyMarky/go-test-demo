package main

import (
	"fmt"
	"strings"
)
/*
string底层是由byte数组构成，且是不可改变的字符数组，
go默认使用的是utf-8格式的unicode字符，
字符代表着一个rune类型，
rune类型对应的是type rune = int32 ，
几乎在所有方面都等同于int32，
字符串复变量赋值后就变成了byte数组，
英文为一个字节，中文是三个字节。

 */
func main() {
	path := "D:\\gopath\\src\\xiaomingkeji.top\\laonanhai"
	fmt.Println(path)

	//多行的字符串且原样输出
	s := `合适的金黄色的
                  大`

	s2 := `D:\gopath\src\xiaomingkeji.top\laonanhai`
 	fmt.Println(s)
	fmt.Println(s2)

	//字符串的长度
	fmt.Println(len(s2))
	//字符串拼接返回
	sprintf := fmt.Sprintf("%s%s", s, s2)
	fmt.Println(sprintf)
	//分割字符串
	ret := strings.Split(path,`\`)
	//遍历数组
	for _, v := range ret {
		fmt.Println(v)
	}
	//前缀与后缀判断
	fmt.Println(strings.HasPrefix(s,"合适"))
	fmt.Println(strings.HasSuffix(s,"大"))

	s4 := "abdcadfdacddfcbs"

	fmt.Println(strings.Index(s4,"a"))
	fmt.Println(strings.LastIndex(s4,"b"))
	//数组拼接成字符串
	fmt.Println(strings.Join(ret,"--"))

}
