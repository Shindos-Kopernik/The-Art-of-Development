package apperror

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNotFound = NewAppError(nil, "not found", "", "US-000003")
)

type AppError struct {
	Err              error  `json:"-"` // игнорируем исходную ошибку
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
	Code             string `json:"code,omitempty"`
}

func (e *AppError) Error() string { // нужен для того, чтобы соответствовать итерфейсу Error
	return e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", err.Error(), "US-000000")
}
