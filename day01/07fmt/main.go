package main

import "fmt"

/**
各种参数输出
 */
func main() {
	var n = 100

	fmt.Printf("%T\n",n) //类型
	fmt.Printf("%v\n",n) //数值
	fmt.Printf("%b\n",n) //二进制
	fmt.Printf("%d\n",n) //十进制
	fmt.Printf("%o\n",n) //八进制
	fmt.Printf("%x\n",n) //十六进制

	var s = "哈哈哈啥"
	fmt.Printf("%s\n",s)
	fmt.Printf("%v\n",s)
	fmt.Printf("%s\n",&s)

}
