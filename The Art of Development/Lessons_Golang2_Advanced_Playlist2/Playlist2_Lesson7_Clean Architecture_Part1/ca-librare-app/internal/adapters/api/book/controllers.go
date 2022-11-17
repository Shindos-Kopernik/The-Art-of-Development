package user

import (
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Architecture_Part1/ca-librare-app/internal/adapters/api"
	"awesomeProject2/My_training_Go/My_training_Go/The Art of Development/Playlist2_Lesson7_Clean Architecture_Part1/ca-librare-app/internal/domain/book"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "books"
)

type handler struct {
	bookService book.Service
}

func NewHandler(service book.Service) api.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// books := h.bookService.GetAllBooks(context.Background(), 0, 0))
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}
