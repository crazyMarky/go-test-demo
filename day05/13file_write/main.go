package main

import (
	"fmt"
	"os"
)

//打开文件写内容

func main() {

	/**
	二进制位控制打开文件的行为
	*/
	_, err := os.OpenFile("./xx,txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("fail err  ", err)
		return
	}

}
