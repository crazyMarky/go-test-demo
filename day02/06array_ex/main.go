package main

import "fmt"

func main() {
	//数组求和,求出指定两个元素的下标
	a1 := [...]int{1,3,5,7,8}

	sum:=0
	for _, v := range a1 {
		sum +=v
	}
	fmt.Println(sum)

	//找出值相加为8的下标组合
	for i:=0;i<len(a1) ;i++  {
		for j:=i;j<len(a1) ;j++  {
			if a1[i] + a1[j] == 8{
				fmt.Printf("(%d,%d)\n",i,j)
			}
		}
	}


}
