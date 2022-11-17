package main

import "context"

func main() {
	context.Background() // С него все начинается
	context.TODO()       // Тут возможно будет логика, но она еще не продумана
	context.WithCancel() // Требует родительского контекста и возвращает его копию.
}
