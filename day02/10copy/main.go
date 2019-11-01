package main

import "github.com/crazyMarky/go-test-demo/utils"

//copy
func main() {

	a1 := []int{1,3,5}
	a2 := a1
	a3 := make([]int,3,5)
	//拷贝不会互相影响
	copy(a3,a1)
	utils.PrintLenAndCapInt(a1)
	utils.PrintLenAndCapInt(a2)
	utils.PrintLenAndCapInt(a3)
	a1[1] = 999;
	utils.PrintLenAndCapInt(a1)
	utils.PrintLenAndCapInt(a1)
	utils.PrintLenAndCapInt(a2)
	utils.PrintLenAndCapInt(a3)

}
