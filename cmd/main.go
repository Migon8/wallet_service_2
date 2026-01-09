package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"wallet_service_2/internal/config"
	"wallet_service_2/internal/handler"
)

func main() {

	config := config.New()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname,
	)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Ошибка открытия подключения:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}

	log.Println("Подключение к БД успешно!")

	http.HandleFunc("/users", handler.GetUsersHandler)
	http.HandleFunc("/users/{id}", handler.GetUsersByIdHandler)
	http.HandleFunc("/users/create", handler.MakeUsersHandler)

	// нам нужна ручка чтобы получить пользователя по id
	// ...

	// нам нужна ручка чтобы создать пользователя
	// ...

	log.Println("Сервер слушает на порту 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка сервера:", err)
	}

}
