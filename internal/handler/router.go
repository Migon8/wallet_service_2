package handler

import (
	"net/http"
	"fmt"
)
func handleRequests() (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodPut:
		handlePut(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.RequestWriter, r *http.Request) {
	fmt.Fprintln(w, "GET request received")
}

func handlePost(w http.RequestWriter, r *http.Request) {
	fmt.Fprintln(w, "POST request received")
}

func handlePut(w http.RequestWriter, r *http.Request) {
	fmt.Fprintln(w, "PUT request received")

}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DELETE request received")
}
