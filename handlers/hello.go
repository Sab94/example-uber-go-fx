package handlers

import (
	"io"
	"net/http"

	"github.com/eg-fx/loggerfx"
)

func NewHandler(logger loggerfx.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Handler Called")
		io.WriteString(w, "Hello World \n")
	})
}
