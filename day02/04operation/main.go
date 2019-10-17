package main

import "fmt"

func main() {
	var (
		a = 50	// 0101
		b = 2   // 0010
		c = 9	// 1001
	)

	fmt.Printf("%b\t%d\n",a&b,a&b)
	fmt.Printf("%b\t%d\n",a|b,a|b)
	fmt.Printf("%b\t%d\n",a^b,a^b)
	fmt.Printf("%b\t%d\n",a^c,a^c)
	fmt.Printf("%b\t%d\n",a<<10,a<<10)
	fmt.Printf("%b\t%d\n",a,a)
	fmt.Printf("%b\t%d\n",a>>3,a>>3)

	var m = int8(1)

	fmt.Printf("%b\t%d\n",m<<4,m<<4)
}
