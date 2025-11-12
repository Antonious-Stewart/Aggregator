package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Antonious-Stewart/Aggregator/internal/config"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var Pool *sql.DB

func logDBErrorValues(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Setup(ctx context.Context) {
	cx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	host, err := config.GetVar("DB_HOST")
	logDBErrorValues(err)

	port, err := config.GetVar("DB_PORT")
	logDBErrorValues(err)

	user, err := config.GetVar("DB_USER")
	logDBErrorValues(err)

	logDBErrorValues(err)

	name, err := config.GetVar("DB_NAME")
	logDBErrorValues(err)

	sslmode, err := config.GetVar("SSLMODE")
	logDBErrorValues(err)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", host, port, user, name, sslmode)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.PingContext(cx); err != nil {
		log.Fatal("Connection to DB Failed")
	}

	log.Println("Connected to DB...")
	Pool = db
}
