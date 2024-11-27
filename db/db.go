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
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Name)

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
	    	weekly_amount FLOAT DEFAULT 0.0,
	    	total_amount FLOAT DEFAULT 0.0,
	    	total_points FLOAT DEFAULT 0.0
		);`,
		`CREATE TABLE IF NOT EXISTS user_points (
    		address TEXT NOT NULL,
	    	week_start DATE NOT NULL,
	    	week_end  DATE NOT NULL,
	    	shared_points FLOAT NOT NULL,
	    	total_points FLOAT NOT NULL,
	    	CONSTRAINT pk_points PRIMARY KEY (address, week_start, week_end)
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
