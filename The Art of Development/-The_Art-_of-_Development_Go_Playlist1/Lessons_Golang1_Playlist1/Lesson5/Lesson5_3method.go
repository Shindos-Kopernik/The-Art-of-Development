package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) movePoint(x, y int) Point { // Было func (p Point) movePoint(p Point, x, y int)
	p.X += x
	p.Y += y
	return p // плохо ниже будем вместо return использовать указатели
}
func (p *Point) movePointPtr(x, y int) { // Было func movePointPtr(p *Point, x, y int)
	p.X += x
	p.Y += y
}
func main() {
	p := Point{1, 2}
	fmt.Println(p)
	fmt.Println(p.movePoint(1, 1)) // fmt.Println(movePoint(p, 1, 1)) пропал т.к. теперь принадлежит структуре
	fmt.Println(p)
	p.movePointPtr(1, 1) // movePointPtr(&p, 1, 1) пропал т.к. теперь принадлежит структуре
	fmt.Println(p)
	v := &p
	v.movePoint(1, 1)
	fmt.Println(p)
	v.movePointPtr(2, 3)
	fmt.Println(p)
	fmt.Println(v)
}
