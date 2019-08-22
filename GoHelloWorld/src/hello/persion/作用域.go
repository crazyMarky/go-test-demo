package main

import (
	"fmt"
	"os"
)

type IP []byte

func main() {
	environ := os.Environ()
	dir, _ := os.Getwd() //
	fmt.Println(dir)
	fmt.Print(environ)
}

func (ip IP) MarshalText() ([]byte, error) {
	if len(ip) == 0 {
		return []byte(""), nil
	}
	return []byte(ip), nil
}
