package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sayHello(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
func sayHello(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}
	close(exit) // Закрыли канал.
	// exit <- 0  TODO В закрытый канал нельзя не чего послать!!!
}
