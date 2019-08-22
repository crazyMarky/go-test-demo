package main
import (
	"fmt"
)
func main() {

	//从指针获取指针指向的值
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)
	ptr2 :=&ptr;
	fmt.Printf("ptr2 type: %T\n", ptr2)
	fmt.Printf("address: %p\n", ptr2)
	fmt.Printf("value type: %T\n", ptr2)
	fmt.Printf("value: %s\n", ptr2)
}