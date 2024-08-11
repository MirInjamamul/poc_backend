package main

import (
	"log"
	"net/http"

	"github.com/Mirinjamamul/go-poc-api/database"
	"github.com/Mirinjamamul/go-poc-api/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load .env file

	database.Connect()
	defer database.Close()

	r := router.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
