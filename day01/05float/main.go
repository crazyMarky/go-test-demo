package main

import "fmt"

func main() {
	f1 := 1.123456
	fmt.Printf("%T\n",f1) //默认go的小数是float64
	//显式声明float32
	f2 := float32(1.123456)
	fmt.Printf("%T\n",f2)
}
