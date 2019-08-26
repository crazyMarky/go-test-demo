package main

import "fmt"

/**
	打印值，体会一下go的声明函数和返回值
 */
func main() {
	sum(1,2)
}

func sum(a,b int )(c , d int){
	fmt.Println(a,b,c)
	c=a+b
	fmt.Println(a,b,c)
	d=a*b
	fmt.Println(a,b,c)
	return c,d
}
