package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) method() {
	fmt.Println("Call Point Method")
}
func main() {
	pointsMap := map[string]int{
		"x": 1,
		"y": 2,
	}

}
