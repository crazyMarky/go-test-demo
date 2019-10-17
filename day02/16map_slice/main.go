package main

import "fmt"

// map slice的组合
func main() {
	//元素为map的切片

	var s1 = make([]map[int]string,10)
	//记得要用make分配内存
	s1[0] =  make(map[int]string, 10)
	s1[0][1] = "你好"
	s1[0][2] = "你好2"
	fmt.Println(s1)
}
