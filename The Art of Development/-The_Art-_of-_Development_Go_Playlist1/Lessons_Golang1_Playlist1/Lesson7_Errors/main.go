package main

import "fmt"

type AppError struct {
	Err    error
	Custom string
	Field  int
}

func (e *AppError) Error() string {
	fmt.Println(e.Custom)
	return e.Err.Error()
}

func main() {
	err := m()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func m() error {
	return &AppError{
		Err:    fmt.Errorf("my error"),
		Custom: "value here",
		Field:  89,
	}
}
