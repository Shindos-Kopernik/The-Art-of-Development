package main

import (
	"fmt"
	"log"
)

func main() {
	divide(1, 0)
	fmt.Println("after panic")
}

func divide(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic happened:", err)
		} // TODO Такой defer позволяет нам управлять и обрабатывать панику
	}()
	fmt.Println(a / b)
}
