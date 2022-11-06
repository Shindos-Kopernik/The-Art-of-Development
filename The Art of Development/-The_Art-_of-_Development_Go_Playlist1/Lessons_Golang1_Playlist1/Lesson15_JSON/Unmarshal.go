package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string `json:"name"` // Чтобы вывод был с маленькой буквы
	Age       int    `json:"age"`
	IsBlocked bool   `json:"isBlocked"`
	Books     []Book `json:"books"`
}

type Book struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func main() {
	byt := []byte(`{"name":"Artur","age":80,"isBlocked":true,"books":[{"name":"BN","year":1990},{"name":"BN2","year":2090}]}`)
	var dat map[string]interface{} // var dat User
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat["books"].([]interface{})[0].(map[string]interface{})["name"]) // fmt.Println(dat.Books[1].Name)
}

func serialize() {
	var books []Book
	book := Book{
		Name: "BN",
		Year: 1990,
	}
	book2 := Book{
		Name: "BN2",
		Year: 2090,
	}
	books = append(books, book, book2)
	sv := User{
		Name:      "Artur",
		Age:       80,
		IsBlocked: true,
		Books:     books,
	}
	boolVar, _ := json.Marshal(sv)
	fmt.Println(string(boolVar))
}
