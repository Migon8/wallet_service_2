package handler

import (
	"fmt"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Write("Список пользователей")
	} else {
		w.Write("Не правильный метод")

	}

}

func GetUsersByIdHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Write("User id")
	} else {
		w.Write("Пользователя нет")
	}

}

func MakeUsersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.Write("Создали пользователя")
	} else {
		w.Write("Попробуйте заново")
	}

}
