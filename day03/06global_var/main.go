package main

import "fmt"

//全局变量
var x = 100

func f1()  {
	//函数先在内部查找
	//再查找外部变量，找到为止，找不到报错
	fmt.Println(x)
}

func main() {
	f1()
}
