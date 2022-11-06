package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	json := `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age": 37,
  "children": ["Sara", "Alex", "Jack"]},
"friends": [
{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb","tw"]},
{"first": "Roger", "last": "Craig", "age": 47, "nets": ["fb","tw"]},
{"first": "Jane", "last": "Murphy", "age": 42, "nets": ["ig", "tw"]},
]}`

	value := gjson.Get(json, "children")
	fmt.Println(value.String())
	if !gjson.Valid(json) {
		panic("JSON IS NOT VALID")
	}
}
