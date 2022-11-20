package author

import (
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Arhitecture_Part2/ca-library-app/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	authorService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	panic("implement me")
}
