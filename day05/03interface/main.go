package main

import "fmt"

type speaker interface {
	speak()
}

type cat struct {
}

type dog struct {
}

type person struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("哈哈哈")
}

func do(x speaker) {
	x.speak()
}

func main() {

	c := cat{}
	d := dog{}
	p := person{}

	do(c)
	do(d)
	do(p)

}
