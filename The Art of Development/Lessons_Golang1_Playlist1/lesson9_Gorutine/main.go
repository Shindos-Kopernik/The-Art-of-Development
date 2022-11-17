package main

import (
	"fmt"
	"time"
)

func main() {
	go say("hello")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	time.Sleep(2 * time.Second)
}

func say(word string) {
	time.Sleep(1 * time.Second)
	fmt.Println(word)
}
