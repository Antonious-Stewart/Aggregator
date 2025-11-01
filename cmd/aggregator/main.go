package main

import (
	"github.com/Antonious-Stewart/Aggregator/internal/db"
	"log"
)

func main() {
	log.Println("Running app...")
	log.Fatal(db.Pool.Ping())
}
