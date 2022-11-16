package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}

				//} else if errors.Is(err, NoAuthErr) {
				//	w.WriteHeader(http.StatusUnauthorized)
				//	w.Write(ErrNotFound.Marshal())
				//	return
				//}

				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest) // когда код выбрасывает error 418, то мы точно знаем, что накосячил наш КОД!
				w.Write(appErr.Marshal())
				return
			}

			w.WriteHeader(http.StatusTeapot) // не наша ошибка, отдаем как есть. А надо обернуть в систем. ошибку
			// w.Write([]byte(err.Error())) // так не красиво и "неправильно"
			w.Write(systemError(err).Marshal()) // Красиво обернутые ситем. ошибки
		}
	}
}
