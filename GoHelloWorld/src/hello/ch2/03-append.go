package main

import "fmt"

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ints := numbers4[2:5:6]
	s1 := append(ints, 10)
	s2 := append(ints, 10,12)
	s3 := append(ints, 10,12,14)
	print(s1)
	print(s2)
	print(s3)
}

func print(s []int)  {
	fmt.Printf("len: %d  cap: %d number: %d\n", len(s), cap(s), s)
}