package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//访问一个网站，并写入到index.html

func main() {
	for _, url := range os.Args[1:] {
		perfix := "http://"
		//加上前缀
		if !strings.HasPrefix(url, perfix) {
			url = perfix + url
		}
		resp, error := http.Get(url)
		s := resp.Status
		fmt.Println(s)
		if error != nil {
			fmt.Println("链接失败：" + url)
			os.Exit(1)
		}
		writer, error := os.Create("index.html")
		if error != nil {
			fmt.Println("创建本地文件失败")
			os.Exit(1)
		}
		io.Copy(writer, resp.Body)
		resp.Body.Close()
	}
}
