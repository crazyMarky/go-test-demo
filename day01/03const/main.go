package main

import "fmt"

const pi = 3.14159

const (
	statusOK = 200
	notFound = 404
)

//如果没有赋值，默认赋值上一行的值
const (
	n1 = 100
	n2
	n3
)

//用于枚举
const (
	a1 = iota
	a2
	a3
	_
	a5
)

//同一个const，隔行被声明，都会加
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//多个常量声明在一行，不会+1
const (
	d1,d2 = iota + 1 , iota +2 //
	d3,d4 = iota + 1, iota + 2
)

func main()  {
	fmt.Println("n1",n1)
	fmt.Println("n2",n2)
	fmt.Println("n3",n3)
	fmt.Println("a3",a3)
	fmt.Println("a5",a5)

	fmt.Println("c1",c1)
	fmt.Println("c2",c2)
	fmt.Println("c3",c3)
	fmt.Println("c4",c4)

	fmt.Println("d1",d1)
	fmt.Println("d2",d2)
	fmt.Println("d3",d3)
	fmt.Println("d4",d4)

	fmt.Println("KB",KB)
	fmt.Println("MB",MB)
	fmt.Println("GB",GB)
	fmt.Println("TB",TB)
	fmt.Println("PB",PB)
}
