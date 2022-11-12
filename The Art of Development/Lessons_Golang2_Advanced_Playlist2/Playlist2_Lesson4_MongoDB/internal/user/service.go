package user

import (
	"Lesson1_Rest_API/pkg/logging"
	"context"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	// TODO for next one
	return
}
