package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := cancelRequest(ctx) // имитация отмены запроса
		if err != nil {
			cancel()
		}
	}()
	doRequest(ctx, "https://ya.ru")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond) // если увеличить время то запрос пройдет и канал done не успеет закрыться.
	return fmt.Errorf("fail request")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response completed, status code=%d", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("request too long")

	}
}
