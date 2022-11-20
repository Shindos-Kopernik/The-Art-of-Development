package book

import (
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Architecture_Part1/ca-librare-app/internal/domain/book"
)

type bookStorage struct {
}

func NewStorage() book.Storage {
	return &bookStorage{}
}

func (bs *bookStorage) GetOne(uuid string) *book.Book {
	return nil
}
func (bs *bookStorage) GetAll(limit, offset int) []*book.Book {
	return nil
}
func (bs *bookStorage) Create(book *book.Book) *book.Book {
	return nil
}
func (bs *bookStorage) Delete(book *book.Book) error {
	return nil
}
