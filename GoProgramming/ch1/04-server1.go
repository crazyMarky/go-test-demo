package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hahah", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//读取本地文件
	bytes, err := ioutil.ReadFile("./GoProgramming/ch1/index.html")
	str, _ := os.Getwd()
	fmt.Println(str)
	if err != nil {
		fmt.Println("fail to read index.html")
		return
	}
	w.Header().Add("content-type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "%s", bytes)
}
