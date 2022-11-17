package main

import (
	"sync"
)

type Counter struct {
	mu sync.RWMutex // Имеет больше контроля над памятью, чем обычный Mutex
	// используется тогда, когда абсолютно уверены, что код кретической секции изменяет охраняемые данные
	c map[string]int
}

func (c *Counter) CountMe() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func (c *Counter) CountMeAgent() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c
}

// func main() {
	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(c.Value(key))
}
