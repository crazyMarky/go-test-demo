package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//readFromFileByio()
	//readFromFileByBufio()
	//readFromFileByIoUtil()
}

//使用字符数组读取数据
func readFromFileByio() {
	//打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file faild,err", err)
		return
	}
	//记得关闭
	defer file.Close()
	for {
		var bytes [1024]byte
		n, err := file.Read(bytes[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("err")
			return
		}
		fmt.Println(string(bytes[:]))
		if n < 1024 {
			return
		}
	}
}

/**
利用bufio一行行的读文件
*/
func readFromFileByBufio() {
	//打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file faild,err", err)
		return
	}
	//记得关闭
	defer file.Close()
	//创建从文件读取内容的对象
	reader := bufio.NewReader(file)
	for {
		//ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。
		// 如果ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。
		// 当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
		line, err := reader.ReadString('\n')
		//注意是字符
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		fmt.Print(line)
	}
}

/**
使用ioutil工具读取文件
*/
func readFromFileByIoUtil() {
	bytes, err := ioutil.ReadFile("./main.go")
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
