package db

import (
	"database/sql"
	"os"
	"fmt"
	"log"
	_"github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s user=%s port=%s password=%s dbname=%s sslmode=disabled",
		host, user, port, password, dbname,
	)

	var err error

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Error pinging DB: ", err)
	}

	fmt.Println("Connected to postgres DB")
}