package main

import "fmt"

/**
	获取数组声明及使用
 */
func main() {

	var numbers2  [5]int
	numbers2[0] = 2
	numbers2[1] = 3
	numbers2[2] = numbers2[0] + 2

	//遍历方式1：只获取键
	for e := range numbers2 {
		fmt.Println(e,numbers2[e])
	}
	fmt.Println("-------------------------我是分割线-------------------------")
	//遍历方式2：键值对
	for k, v := range numbers2 {
		fmt.Println(k,v)
	}
	fmt.Println("-------------------------我是分割线-------------------------")
	//获取数组的长度
	fmt.Print(len(numbers2))




}
