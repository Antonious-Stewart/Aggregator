package db

import (
	"database/sql"
	"fmt"
	"github.com/Antonious-Stewart/Aggregator/internal/config"
	_ "github.com/lib/pq"
	"log"
)

var Pool *sql.DB

func logDBErrorValues(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	host, err := config.GetVar("DB_HOST")
	logDBErrorValues(err)

	port, err := config.GetVar("DB_PORT")
	logDBErrorValues(err)

	user, err := config.GetVar("DB_USER")
	logDBErrorValues(err)

	password, err := config.GetVar("DB_PASSWORD")
	logDBErrorValues(err)

	name, err := config.GetVar("DB_NAME")
	logDBErrorValues(err)

	sslmode, err := config.GetVar("SSLMODE")
	logDBErrorValues(err)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, name, sslmode)
	log.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	Pool = db
}
