package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	sv := map[string]interface{}{"field": "value", "field2": 123, "field3": true, "arr": []string{"a", "b", "c"}}
	boolVar, _ := json.Marshal(sv)
	fmt.Println(string(boolVar))
}
