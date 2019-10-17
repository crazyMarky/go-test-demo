package main

import "fmt"

//return 不是真正的原子操作
// 第一步：返回赋值
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，那么执行的时间是在第一步和第二步之间

func f1()  int{
	x := 5
	defer func() {
		x++	//修改的是x，不是修改返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5		// 返回值 = x
}

func f3() (y int)  {
	x := 5
	defer func() {
		x++	//修改的是x
	}()
	return x	// 5
}

func f4() (x int)  {
	defer func(x int) {
		//局部变量，不改变外部x
		x++
	}(x)
	return 5	 //5
}


func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
