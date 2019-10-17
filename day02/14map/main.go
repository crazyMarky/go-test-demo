package main

import "fmt"

// var m1 map
func main() {
	var m1 map[string]int	//只声明了变量，没有在内存中开辟空间

	m1 = make(map[string]int,10)	//分配内存空间，提前估算好改容量

	m1["理想"] = 18
	m1["哈哈"] = 123

	for k, v := range m1 {
		fmt.Println(k,v)
	}

	//如果不存在，就返回零值
	 v,o := m1["123"]

	fmt.Println(v,o)


}
