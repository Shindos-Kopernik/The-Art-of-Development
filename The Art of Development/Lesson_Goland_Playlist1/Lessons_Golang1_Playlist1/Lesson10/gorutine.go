package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // примитив синхронизации, который позваляет дождаться завершения горутин в оди момент времени
	for i := 0; i < 10; i++ {

		wg.Add(1)
		k := i
		go func() {
			defer wg.Done()
			fmt.Printf("%d goroutine working ..\n", k)
			time.Sleep(300 * time.Millisecond)
		}()
	}
	wg.Wait()
	fmt.Println("all done")
}
