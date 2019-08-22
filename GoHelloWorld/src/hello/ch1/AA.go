package main


import (
"fmt"
)

const c = "C"

var v int = 5

type T struct{
	a int;
}

func init() { // initialization of package
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}


func (t T) Method1() {
	fmt.Println(t)
}

func Func1() { // exported function Func1
	//...
}