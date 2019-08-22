package main

import "fmt"

func main() {
	sum(1,2)
	fmt.Println()
}

func sum(a,b int )(c , d int){
	fmt.Println(a,b,c)
	c=a+b
	fmt.Println(a,b,c)
	d=a*b
	fmt.Println(a,b,c)
	return c,d
}
