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

type Pinger interface {
	Ping() error
}

type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Querier interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type Runner interface {
	Executor
	Querier
	Pinger
}

var lock = &sync.Mutex{}

type Database struct {
	Pool *sql.DB
}

func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Pool.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (db *Database) ExecContext(context context.Context, query string, args ...interface{}) (sql.Result, error) {

	return nil, nil
}

func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (db *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func (db *Database) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (db *Database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return nil
}

var instance *Database

func (db *Database) Ping() error {
	if err := db.Pool.PingContext(context.Background()); err != nil {
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
