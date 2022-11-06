package main

import "fmt"

type appError struct {
	Err    error
	Custom string
	Field  int
}

func (e *appError) Error() string {
	return fmt.Sprintf("%s error", e.Custom)
}
func (e *appError) Unwrap() error { // Unwrap - распокавать ошибку
	return e.Err
}
func main() {
	err := method1()
	if err != nil {
		fmt.Println(err.Unwrap())
		return
	}
	fmt.Println("success")
}
func method1() *appError {
	return method2()
}
func method2() *appError {
	return method3()
}
func method3() *appError {
	// implementation business logic
	return &appError{
		Err:    fmt.Errorf("internal error"),
		Custom: "something goes wrong",
	}
}
