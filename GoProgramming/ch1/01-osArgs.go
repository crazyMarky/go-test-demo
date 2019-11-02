package main

import (
	"fmt"
	"os"
	"strings"
)

//展示命令行参数
//os.Args 从1 开始
func main() {
	echo1()
	echo2()
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	args := os.Args
	fmt.Println(strings.Join(args[1:], " - "))
}

//练习：尝试测量可能低效的程序
