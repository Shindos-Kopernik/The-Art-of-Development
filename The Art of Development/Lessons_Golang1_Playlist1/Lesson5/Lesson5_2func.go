package main

import "fmt"

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	orderPrice := totalPrice(1) // sum = 1
	fmt.Println(orderPrice(1))  // sum = 1 + 1 = 2
	fmt.Println(orderPrice(1))
}
