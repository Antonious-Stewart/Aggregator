package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Antonious-Stewart/Aggregator/internal/config"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

type Database struct {
	Pool *sql.DB
}

type Pinger interface {
	Ping() error
}

var instance *Database

func (d *Database) Ping() error {
	if err := d.Pool.PingContext(context.Background()); err != nil {
		log.Println("Connection Unsuccessful")
		return err
	}

	log.Println("Connection Successful")
	return nil
}

func GetInstance() *Database {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating DB instance now.")
			instance = newInstance()
			if err := instance.Ping(); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Database instance already exists.")
		}
	} else {
		fmt.Println("Database instance already exists.")
	}

	return instance
}

func newInstance() *Database {
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

	return &Database{
		Pool: db,
	}
}

func logDBErrorValues(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
