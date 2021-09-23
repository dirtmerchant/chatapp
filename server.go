package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	log.Print("Server starting at localhost:" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
