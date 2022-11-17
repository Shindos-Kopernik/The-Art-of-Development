package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex // Mutex защищает с и должен стоять выше значения, которое он защищает (организацион. договоренность)
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock() // Не держите блокировку дольше чем это требуется
	c.c[key]++  // TODO код между Lock & Unlock называется "критичиская секция"(КС). Не рекоменд. в КС держать много кода
	// КС - код эксклюзивного доступа
	c.mu.Unlock()
	// fmt.Println("done")
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() // Использовать когда в return критический код
	return c.c[key]
}

func main() {
	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(c.Value(key))
}
