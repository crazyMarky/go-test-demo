package main

import (
	"fmt"
	"unicode/utf8"
)

/**

 */

func main() {

	s := "傻乎乎的的饭卡手动阀了"

	len2 := len(s)
	fmt.Println(len2)

	s2 := "白萝卜"
	s3 := []rune(s2)	//转换成rune切片
	s3[0] = '黑'
	fmt.Println(string(s3))
	//字符串的区别
	s4:="哈"
	s5:='哈'	//rune(int32)
	s6 := "H"	//string
	s7 := 'H'	//byte(uint8)
	fmt.Printf("%T,%T\n",s4,s5)
	fmt.Printf("%T,%T\n",s6,s7)

	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n",f)

	c1 := "hello沙河小王子"
	l := len(c1)
	bytes := []byte(c1)
	btyeLen := len(bytes)
	i := btyeLen - l
	fmt.Println(i/3)
	count := utf8.RuneCountInString(c1)
	fmt.Println(count)
}
