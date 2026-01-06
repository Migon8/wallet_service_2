package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func New() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")


	config := Config{

		host: host,
		port: port, 
		user: user, 
		password: password, 
		dbname: dbname,
	}

}
