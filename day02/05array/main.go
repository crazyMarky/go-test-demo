package main

import "fmt"

func main() {

	//数组
	var a1 [3]bool
	var a2 [4]bool
	a3 := [...]bool{false,false}

	fmt.Printf("a1:%T a2:%T a3:%T\n",a1,a2,a3)

	for _, v := range a1 {
		fmt.Println(v)
	}

	//bool值
	a1 = [3]bool{true,true,true}
	//索引给值
	a4 := [9]int{0,1,2,7:7,9}
	fmt.Println(a1)
	fmt.Println(a4)
	//数组的遍历
	citys := [...]string{"北京","上海","南京"}
	for _, v := range citys {
		fmt.Println(v)
	}

	//多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println(a11)
	//多维数组的遍历
	//for _, v1 := range a11 {
	//	//fmt.Println(v1)
	//	for _, v2 := range v1 {
	//		fmt.Println(v2)
	//	}
	//}

	//数组是值类型
	b1 := [3]int{1,2,3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1,b2)
}
