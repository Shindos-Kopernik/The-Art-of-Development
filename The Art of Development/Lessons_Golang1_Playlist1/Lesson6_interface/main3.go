package main

import "fmt"

func main() {
	var a interface{} // Интерфейс как отдельный тип данных и может хранить любые типы данных (int, strung, итд)
	a = "jelly"
	fmt.Println(a)
	fmt.Printf("(%v, %T)\n\n", a, a)
	a = 42
	fmt.Println(a)
	fmt.Printf("(%v, %T)\n\n", a, a)
}
