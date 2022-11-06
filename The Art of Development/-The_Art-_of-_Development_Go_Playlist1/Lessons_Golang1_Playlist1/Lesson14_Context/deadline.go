package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	duration := 1 * time.Second
	ctx := context.Background()
	d := time.Now().Add(duration)
	ctx, cancel := context.WithDeadline(ctx, d)
	// ctx := context.WithValue(ctx, "test", "hello") Положить значения по ключу, используют для конкретного значения,
	// т.е. только на запрос
	defer cancel()

	doRequest(ctx, "https://ya.ru")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(1500 * time.Millisecond)
	return fmt.Errorf("fail request")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response completed, status code=%d", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("request too long")

	}
}
