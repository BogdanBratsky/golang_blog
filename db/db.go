package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

const (
	dsn = "user=postgres password=2048 host=localhost port=5432 dbname=goblogdb sslmode=disable"
)

// функция подключения к бд

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

// функция закрытия подключения к бд

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
