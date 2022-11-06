package main

import "fmt"

type name struct {
	A, B int
}

func (n *name) method() {
	fmt.Println("ok")
	fmt.Println(n.A)
}
func main() {
	n := &name{1, 2}
	n = nil
	n.method()
}
