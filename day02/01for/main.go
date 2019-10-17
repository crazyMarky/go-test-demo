package main

import "fmt"

func main() {

	for i := 1;i<=9;i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	for i := 1;i<=9;i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}

}
