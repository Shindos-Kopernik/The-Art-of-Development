package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type InterfaceHere interface {
	Sum() int
}

func main() {
	var os *structHere
	var i InterfaceHere
	i = os
	//fmt.Println(i.Sum())
	fmt.Printf("(%v, %T)\n", i, i)
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}
