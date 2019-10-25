package main

import "fmt"

//接口的实现
//1必须实现所有的方法才算是实现了接口

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫冻")
}
func (c cat) eat(s string) {
	fmt.Println("猫吃鱼")
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡冻")
}

func (c chicken) eat(s string) {
	fmt.Println("鸡吃饲料")
}

func main() {
	var a1 animal

	bc := cat{
		name: "小猫",
		feet: 4,
	}

	a1 = bc
	fmt.Println(a1)

	c := chicken{
		feet: 2,
	}
	fmt.Printf("%T\n", bc)
	fmt.Printf("%T\n", c)
}
