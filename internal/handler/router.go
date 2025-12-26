package handler

import (
	"io"
	"net/http"
)

type UserHandler struct {
	// Add necessary fields here
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteSring(w, "Hello, User!")

}
