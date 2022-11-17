package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	json := `{
 "name": {"first": "Tom", "last": "Anderson"},
  "age": 37,
  "children": ["Sara", "Alex", "Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb","tw"]},
    {"first": "Roger", "last": "Craig", "age": 47, "nets": ["fb","tw"]},
    {"first": "Jane", "last": "Murphy", "age": 42, "nets": ["ig", "tw"]
  }]
}`
	//if !gjson.Valid(json) {
	//	panic("JSON IS NOT VALID")
	//}
	//gjson.AddModifier("case", func(json, arg string) string {
	//if arg == "upper" {
	//return strings.ToUpper(json)
	//} else if arg == "lower" {
	//	return strings.ToLower(json)
	//}
	//return json
	//})
	value := gjson.Get(json, `friends.#(age>=42)#.first`)
	// fav.movie если напишем с точкой, то не получим не чего fav\\.movie -good!
	// value := gjson.Get(json, "children.#") - получим кол-во детей(размерность массива)
	// child*.2 -> Jack; friends.1.last -> Craig; `friends.#(last=="Murphy").first` -> Dale;
	//`friends.#(last== age).first`(первый попавшийся эл-т), чтобы всех нашел (json, `friends.#(age==47)#.first`)
	fmt.Println(value.String())
	//fmt.Println(json.Parse(json).Get("name"))
	//resault, ok := gjson.Parse(json).Value().(map[string]interface{})
	//if !ok {
	//	panic("NOT OKAY PARSING TO MAP")
	//}
	//fmt.Println(resault)
}
