package main

import (
	"errors"
	"fmt"
	"log"
)

type AppError struct {
	Messege string
	Err     error
}

func (ae *AppError) Error() string {
	return ae.Messege

}
func main() {
	divide(1, 0)
	fmt.Println("after panic")
}
func divide(a, b int) {
	defer func() {
		var appErr *AppError
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("panic!", err) // this is AppError handle panic
				} else {
					fmt.Println("custom panic") //this is Other Error handle panic
				}
			default:
				panic("some panic") // this is default GO panic
			}
			log.Println("panic happened:", err)
		}
	}()
	fmt.Println(div(a, b))
}
func div(a, b int) int {
	if b == 0 {
		panic(&AppError{
			Messege: "this is divide by zero custom error",
			Err:     nil,
		})
	}
	return a / b
}
