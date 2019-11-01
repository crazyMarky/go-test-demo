package main

import "fmt"

func f1()  {
	fmt.Println("fjsdjkklsd")
}

func f2()  int{
	fmt.Println("fjsdjkklsd")
	return 10
}

//函数作为参数
func f3(f func())  {
	f()
}

//函数可以是返回值
func f5() func(int,int) int {
	ret := func(a,b int) int { return a+b}
	return ret
}

func f6() (int,int)  {
	return 1,2
}

func main() {
	//函数也是类型
	a := f1
	fmt.Printf("%T\n",a)
	a2 := f2
	fmt.Printf("%T\n",a2)
	f3(a)
	a3 := f5()

	fmt.Println(a3(1,2))

	i, _ := f6()

	fmt.Println(i)


}
