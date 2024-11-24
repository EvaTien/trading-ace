package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"trading-ace/config"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)
	log.Println(connStr)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	log.Println("Successfully connected to the database")
}
