package main

import "fmt"

type Point struct {
	Y int
	X int
}

func (p Point) method() {
	fmt.Println("Call Point Method")
}
func main() {
	pointsMap := map[string]Point{
		"b": {
			Y: 13,
			X: 15,
		},
	}
	otherPointsMap := make(map[int]Point)
	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	//fmt.Println(pointsMap)
	//fmt.Println(pointsMap["a"])
	otherPointsMap[1] = Point{1, 2}
	//fmt.Println(otherPointsMap)
	//fmt.Println(otherPointsMap[1])
	var oneMorePointsMap map[string]Point
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{
			"a": {Y: 1, X: 2},
		}
	}
	//fmt.Println(oneMorePointsMap["a"].Y)
	//fmt.Println(oneMorePointsMap["a"].X)
	oneMorePointsMap["a"].method()
	// Проверка есть ли такой ключ в map

	key := "a"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key=%s exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Println("error")
		fmt.Println(value)
	}

}
