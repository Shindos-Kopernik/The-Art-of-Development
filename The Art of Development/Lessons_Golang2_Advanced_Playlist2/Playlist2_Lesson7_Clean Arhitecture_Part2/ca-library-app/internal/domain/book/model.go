package book

import (
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/domain/author"
	"fmt"
)

type Book struct {
	UUID   string        `json:"uuid,omitempty"`
	Name   string        `json:"name,omitempty"`
	Year   string        `json:"year,omitempty"`
	Author author.Author `json:"author,omitempty"`
	Genre  genre.Genre   `json:"genre"`
	Busy   bool          `json:"busy,omitempty"`
	Owner  string        `json:"owner,omitempty"`
}

func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book is busy")
	}
	b.Owner = owner
	b.Busy = true
	return nil
}
