package main

import "fmt"

/**
	遍历字符串，用ASCII和Unicode遍历的不同
 */
func main() {

	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c  %d\n", theme[i], theme[i])
	}


	for _, s := range theme {
		fmt.Printf("Unicode: %c  %d\n", s, s)
	}

	fmt.Print("%c %d ",0xcccc,0xcccc)
	
}
