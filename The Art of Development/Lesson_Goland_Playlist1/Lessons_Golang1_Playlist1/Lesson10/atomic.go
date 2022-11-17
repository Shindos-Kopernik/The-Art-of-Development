package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup // примитив синхронизации, который позваляет дождаться завершения горутин в один момент времени
	var counter uint64
	//var mu sync.Mutex
	for i := 0; i < 10; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				//mu.Lock()
				//counter++
				//mu.Unlock()
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf(" all done, counter=%d\n", counter)
}
