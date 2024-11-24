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

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	log.Println("Successfully connected to the database")
}

func CreateTables() {
	// Create the users table
	migrationQueries := []string{
		`CREATE TABLE IF NOT EXISTS users (
    		address TEXT PRIMARY KEY,
	    	onboarding_completed BOOLEAN DEFAULT FALSE,
	    	weekly_amount INTEGER DEFAULT 0,
	    	total_amount INTEGER DEFAULT 0,
	    	total_points INTEGER DEFAULT 0
		);`,
		`CREATE TABLE IF NOT EXISTS user_points (
    		address TEXT PRIMARY KEY,
	    	week_start DATE NOT NULL,
	    	week_end  DATE NOT NULL,
	    	shared_points INTEGER NOT NULL,
	    	total_points INTEGER NOT NULL
		);`,
	}

	for _, query := range migrationQueries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Error create table: ", err)
		}
	}
	fmt.Println("Tables created.")
}
