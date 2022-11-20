package main

import (
	author3 "awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/adapters/api/author"
	book3 "awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/adapters/api/book"
	author2 "awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/adapters/db/author"
	book2 "awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/adapters/db/book"
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/domain/author"
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/domain/book"
)

func main() {
	// entry point  точка входа в приложение
	bookStorage := book2.NewStorage()
	bookService := book.NewService(bookStorage)
	bookHandler := book3.NewHandler(bookService)

	authorStorage := author2.NewStorage()
	authorService := author.NewService(authorStorage)
	authorHandler := author3.NewHandler(authorService)
}
