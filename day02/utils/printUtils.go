package utils

import "fmt"

type Slice []interface{}

func NewSlice() Slice {
	return make(Slice, 0)
}


func PrintLenAndCapStr(s1 []string)  {
	fmt.Printf("%v == len(s):%d cap(s):%d == %p\n",s1,len(s1), cap(s1),&s1)
}

func PrintLenAndCapInt(s1 []int)  {
	fmt.Printf("%v == len(s):%d cap(s):%d == %p\n",s1,len(s1), cap(s1),&s1)
}
