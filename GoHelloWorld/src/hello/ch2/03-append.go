package main

import "fmt"

/**
append()为切片添加元素
 */
func main() {
	var a []int
	a = append(a, 1) // 追加1个元素
	print(a)

	a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	print(a)

	a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
	print(a)
}

func print(s []int)  {
	fmt.Printf("len: %d  cap: %d number: %d\n", len(s), cap(s), s)
}