package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	appendToFile()
}
func appendToFile() {
	f, err := os.OpenFile("test2.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(" Artur"); err != nil {
		panic(err)
	}
}

func writeToFile() {
	data := []byte("My name is Artur")
	err := ioutil.WriteFile("test2.txt", data, 0600) // chmod u+x test2.t
	if err != nil {
		fmt.Println(err)
	}

}
func readFile() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
