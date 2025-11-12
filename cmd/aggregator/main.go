package main

import (
	"context"
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

	ctx := context.Background()
	db.Setup(ctx)

	port, err := config.GetVar("PORT")

	handler := apis.Handler(db.Pool)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
