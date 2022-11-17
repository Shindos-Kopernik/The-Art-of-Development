package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")
	arr := make([]byte, 8)
	for {
		n, err := r.Read(arr)
		fmt.Printf("n = %d err = %v arr = %v\n", n, err, arr) // %d - Якорь для числа, %v - Якорь для интерфейса
		fmt.Printf("arr n bytes: %q", arr[:n])
		if err == io.EOF {
			break
		}
	}
}
