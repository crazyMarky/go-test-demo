package main

import "xiaomingkeji.top/laonanhai/day02/utils"

// make 创建切片

func main() {
	//长度是5，容量是10
	s1 := make([]int,5,10)
	utils.PrintLenAndCapInt(s1)

	//长度是0，容量是10
	s2 := make([]int,0,10)
	utils.PrintLenAndCapInt(s2)

	//判断切片为空，应该使用len(s) == 0
	s3 := []int{1,3,5}
	s4 := s3	//s3和s4都指向同一个底层数组
	s4[1] = 99
	utils.PrintLenAndCapInt(s3)
	utils.PrintLenAndCapInt(s4)

}


