package main

import "fmt"

func deferDemo(){
	fmt.Println(11)
	defer  fmt.Println(22) //延迟执行到即将返回时再执行
	fmt.Println(33)
	defer  fmt.Println(44) //延迟执行到即将返回时再执行
}

func main() {
	deferDemo()
}
