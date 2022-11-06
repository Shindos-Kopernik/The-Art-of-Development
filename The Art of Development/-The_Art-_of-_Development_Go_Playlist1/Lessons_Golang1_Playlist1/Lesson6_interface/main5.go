package main

import "fmt"

func main() {
	var a interface{} = true

	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	default:
		fmt.Printf("unknown type %T\n", a)
	}
}
