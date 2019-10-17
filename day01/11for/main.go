package main

import "fmt"

//for
func main() {
	//99乘法表
	for i := 1;i<=9;i++ {
		for j:=1;j<=i ;j++  {
			fmt.Printf("%dx%d=%d\t",j,i,i*j)
		}
		fmt.Println()
	}

}
