package main

import "fmt"

/**
切片
- 前闭后开
- 切片本身没有数据，是对底层array的一个view，所以切片作为函数输入，相当于传地址
- 切片是可以向后扩展的，但不可以向前扩展，s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
- 添加元素：append
- 拷贝
- 删除

 */
func main() {
	//前闭后开，三个参数，起始位置，最大长度
	arr := [...]int {0,1,2,3,4,5,6,7}
	fmt.Println("This arr[2:6]: ", arr[2:6] ) // output: This arr[2:6]:  [2 3 4 5]
	fmt.Println("This arr[:6]: ", arr[:6] ) // output: This arr[:6]:  [0 1 2 3 4 5]
	fmt.Println("This arr[2:]: ", arr[2:] ) // output: This arr[2:]:  [2 3 4 5 6 7]
	fmt.Println("This arr[:]: ", arr[:] ) // output: This arr[:]:  [0 1 2 3 4 5 6 7]
	//切片本身没有数据，是对底层array的一个view，所以切片作为函数输入，相当于传地址
	a2 := arr[:]
	fmt.Print(a2)
	//切片是可以向后扩展的，但不可以向前扩展，s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)


}