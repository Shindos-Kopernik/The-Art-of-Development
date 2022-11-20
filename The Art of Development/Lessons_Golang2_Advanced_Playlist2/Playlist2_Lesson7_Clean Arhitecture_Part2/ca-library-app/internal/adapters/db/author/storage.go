package author

import (
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/domain/author"
)

type authorStorage struct {
}

func NewStorage() author.Storage {
	return &authorStorage{}
}

func (bs *authorStorage) GetOne(uuid string) *author.Author {
	return nil
}
func (bs *authorStorage) GetAll(limit, offset int) []*author.Author {
	return nil
}
func (bs *authorStorage) Create(book *author.Author) *author.Author {
	return nil
}
func (bs *authorStorage) Delete(book *author.Author) error {
	return nil
}
