package main

import (
	"github.com/Antonious-Stewart/Aggregator/internal/apis"
	"github.com/Antonious-Stewart/Aggregator/internal/config"
	"github.com/Antonious-Stewart/Aggregator/internal/db"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database := db.GetInstance()

	port, err := config.GetVar("PORT")

	handler := apis.Routes(database)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
