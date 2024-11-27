package db

import (
	"database/sql"
	"errors"
	"log"
)

func CreateNewUser(address string) {
	insertQuery := `INSERT INTO users (address) VALUES ($1)`
	_, err := DB.Exec(insertQuery, address)
	if err != nil {
		log.Fatal("Failed to insert record:", err)
	}
}

func UserExisted(address string) bool {
	query := `SELECT 1 FROM users WHERE address = $1;`
	var result int
	err := DB.QueryRow(query, address).Scan(&result)
	if errors.Is(err, sql.ErrNoRows) {
		// No result means user doesn't exist
		return false
	} else if err != nil {
		log.Fatal("Error executing query:", err)
		return false
	} else {
		// User exists
		return true
	}
}

func GetUser(address string) map[string]interface{} {
	query := `SELECT address, onboarding_completed, weekly_amount, total_amount, total_points FROM users WHERE address = $1;`

	// Execute the query
	var addr string
	var onboarding_completed bool
	var weeklyAmount, totalAmount, totalPoints float32
	err := DB.QueryRow(query, address).Scan(&addr, &onboarding_completed, &weeklyAmount, &totalAmount, &totalPoints)
	if err == sql.ErrNoRows {
		log.Printf("User not found")
		return nil
	} else if err != nil {
		log.Fatalf("Error executing query: %s", err)
		return nil
	}

	user := map[string]interface{}{
		"address":              addr,
		"onboarding_completed": onboarding_completed,
		"weekly_amount":        weeklyAmount,
		"total_amount":         totalAmount,
		"total_points":         totalPoints,
	}
	return user
}
