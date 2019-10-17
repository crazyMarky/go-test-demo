package main

import "fmt"

//函数的定义
func sum(x int , y int)(ret int){
	return x+y
}

//可变长参数
func f7(x string , y ...int)  {
	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	ret := sum(10, 20)
	fmt.Println(ret)
	f7("fdfdsdfdf",23,45,6,7)
}
