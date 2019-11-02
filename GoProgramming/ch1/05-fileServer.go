package main

import (
	"log"
	"net/http"
)

func main() {
	httpserver()
}

func httpserver() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
