package user

import "context"

type Storage interface {
	Create(ctx context.Context, user User) (string, error) // Метод создания пользователя
	FindOne(ctx context.Context, id string) (User, error)  //Поиск пользователя по индификатору
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
