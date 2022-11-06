package main

import "fmt"

func main() {
	var a interface{} = "hello"

	s := a.(string)
	fmt.Println(s)

	s, ok := a.(string)
	fmt.Println(s, ok)
	f, ok := a.(float32)
	fmt.Println(f, ok)

	g := a.(float64) // panic
	fmt.Println(g)

}
