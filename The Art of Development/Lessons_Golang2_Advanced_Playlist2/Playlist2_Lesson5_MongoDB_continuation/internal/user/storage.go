package user

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, user User) (string, error) // Метод создания пользователя
	FindAll(ctx context.Context) (u []User, err error)
	FindOne(ctx context.Context, id string) (User, error) //Поиск пользователя по индификатору
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
