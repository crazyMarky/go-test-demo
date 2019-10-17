package main

import "fmt"
/**
不能转换其他类型
 */
func main() {
	b1 := true
	var b2 bool
	fmt.Printf("%T\n",b1)
	fmt.Printf("%T,value:%v\n",b2,b2)

}
