package main

import "xiaomingkeji.top/laonanhai/day02/utils"

// append()追加
func main() {
	s1 := []string{"北京","上海","南京"}
	utils.PrintLenAndCapStr(s1)
	//调用append函数必须用原切片接受返回值
	//数组放不下了，小于1024容量翻倍,底层数组会搬家，复制到别的地方
	s1 = append(s1,"深圳")
	utils.PrintLenAndCapStr(s1)
	s1 = append(s1,"杭州","武汉","苏州")
	utils.PrintLenAndCapStr(s1)
}



