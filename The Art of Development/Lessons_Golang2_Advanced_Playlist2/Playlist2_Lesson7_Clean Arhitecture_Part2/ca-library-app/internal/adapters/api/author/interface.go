package author

import (
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/domain/author"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *author.Author
	GetAll(ctx context.Context, limit, offset int) []*author.Author
	Create(ctx context.Context, dto *author.CreateAuthorDTO) *author.Author
}
