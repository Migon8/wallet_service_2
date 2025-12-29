package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	
)

func main() {


	config := config.New()



	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
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

	// 👇 вот этот маршрут ДОЛЖЕН быть
	http.HandleFunc("/users", usersHandler)

	log.Println("Сервер слушает на порту 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка сервера:", err)
	}

}
