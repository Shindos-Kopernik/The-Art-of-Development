package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	//gjson
	json := `{"name": {"first": "Artur", "last": "Karapetov"}, "age": 47}`
	value := gjson.Get(json, "name.first")
	fmt.Println(value.String())
}
