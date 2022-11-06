package main

import "fmt" // TODO Для многослойного приложения!

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
	err := method1()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}
func method1() error {
	return method2()
}
func method2() error {
	return method3()
}
func method3() error {
	// implementation business logic
	// return fmt.Errorf("some error") TODO Выведет ошибку
	return nil // success
}
