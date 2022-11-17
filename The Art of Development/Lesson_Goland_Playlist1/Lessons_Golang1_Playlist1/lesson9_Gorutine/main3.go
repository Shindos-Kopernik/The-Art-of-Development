package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sayHello(ch)
	s := <-ch // Чтение из канала
	fmt.Println(s)
	// Читаем из канала

}
func say(word string) {
	fmt.Println(word)
}

func sayHello(exit chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		say("hello")

	}
	exit <- "s" // Запись в канал
}
