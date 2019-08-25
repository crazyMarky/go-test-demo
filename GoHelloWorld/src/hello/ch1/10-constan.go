package main

import "fmt"

func main() {
	const (
		e  = 2.7182818
		pi = 3.1415926
	)

	//iota 常量生成器
	//常量声明可以使用 iota 常量生成器初始化
	// 它用于生成一组以相似规则初始化的常量
	// 但是不用每行都写一遍初始化表达式。
	// 在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一。
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Thursday)

	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)
}
