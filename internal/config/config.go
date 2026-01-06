package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func New() Config {

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

		Host: host,
		Port: port, 
		User: user, 
		Password: password, 
		Dbname: dbname,
	}

	return config 

}
