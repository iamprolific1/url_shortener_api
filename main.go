package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	
	"github.com/iamprolific1/db"
)

func main() {
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading up .env files")
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "4000"
	}

	db.Connect()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "server entry point")
	})

	fmt.Println("application is running on server http://localhost:4000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}