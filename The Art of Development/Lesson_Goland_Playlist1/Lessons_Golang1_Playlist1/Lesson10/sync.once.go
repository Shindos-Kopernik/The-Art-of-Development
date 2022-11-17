package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once // Примитив синхронизации, который позволяет выполнить функцию только один раз.
	done := make(chan bool)

	for i := 0; i < 30; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("the only one")
			})
			done <- true
		}()
	}
	for i := 0; i < 30; i++ {
		<-done
	}

}
