package main

import (
	"fmt"
	"github.com/tidwall/sjson"
)

func main() {
	json := `{
 "name": {"first": "Tom", "last": "Anderson"},
  "age": 37,
  "children": ["Sara", "Alex", "Jack"],
  "fav.movie": "Deer Hunter",
  "friends":[
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb","tw"]},
    {"first": "Roger", "last": "Craig", "age": 47, "nets": ["fb","tw"]},
    {"first": "Jane", "last": "Murphy", "age": 42, "nets": ["ig, "tw"]}
  ]
}`
	value, _ := sjson.Set(json, "name.first", "Artur")
	value2, _ := sjson.Set(json, "children.1", "Ira")
	value3, _ := sjson.Set(json, "children.-1", "Lena")
	value4, _ := sjson.Set(json, "new_obj", map[string]interface{}{"hello": "world"})
	fmt.Println(value)
	fmt.Println(value2)
	fmt.Println(value3)
	fmt.Println(value4)
}
