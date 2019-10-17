package main

import "fmt"

func main() {
	//切片的定义
	var s1 []int //定义
	var s2 []string
	fmt.Println(s1,s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//初始化
	s1 = []int{1,2,3,4,5}
	s2 = []string{"哈哈","喵喵","喵1喵"}
	fmt.Println(s1,s2)

	//长度和容量
	//切片的容量是指首个位置到底层数组最末的容量
	fmt.Printf("len(s1):%d cap(s1):%d\n",len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n",len(s2), cap(s2))

	//从数组得到切片
	a1 := [...]int{1,3,5,7,8,11,13}
	s3 := a1[0:4]	//左闭右开
	s4 := a1[1:4]	//左闭右开
	s5 := a1[:4]	//左闭右开
	s6 := a1[3:]	//左闭右开
	fmt.Println(a1,s3)
	fmt.Println(a1,s4)
	fmt.Println(a1,s5)
	fmt.Println(a1,s6)
	fmt.Printf("len(s5):%d cap(s5):%d\n",len(s5), cap(s5))
	fmt.Printf("len(s6):%d cap(s6):%d\n",len(s6), cap(s6))

	//切片再切片
	s8 := s6[3:]
	fmt.Println(s8)
	fmt.Printf("len(s8):%d cap(s8):%d\n",len(s8), cap(s8))

	//切片是引用的底层数组，所以修改底层数组会影响切片的结果
	fmt.Println("s6:",s6)
	a1[5] = 999;
	fmt.Println("s6:",s6)


}
