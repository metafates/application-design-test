package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

var _ http.Handler = (*_LoggingHandler)(nil)

func Logging(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return _LoggingHandler{
			logger:  logger,
			handler: next,
		}
	}
}

type _LoggingHandler struct {
	logger  *slog.Logger
	handler http.Handler
}

func (l _LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.handler.ServeHTTP(w, r)
	if r.MultipartForm != nil {
		err := r.MultipartForm.RemoveAll()
		if err != nil {
			return
		}
	}

	l.logger.Info("http", "method", r.Method, "path", r.URL.Path)
}
