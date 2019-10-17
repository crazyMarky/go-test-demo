package main

import "fmt"

var (
	name string
	age int
	isOk bool
)
//Go的非全局变量声明必须使用
//匿名变量不占用命名空间，不分配内存

func main()  {
	name = "离线"
	name2 := "hahaha"
	fmt.Print(isOk)
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Printf("my name:%s",name)
}