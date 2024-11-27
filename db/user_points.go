package db

import (
	"log"
)

func GetUserPointsByAddress(address string) []map[string]interface{} {
	query := `SELECT address, week_start, week_end, shared_points, total_points FROM user_points WHERE address = $1;`
	rows, err := DB.Query(query, address)
	if err != nil {
		log.Fatalf("Error executing query: %s", err)
		return nil
	}
	defer rows.Close()

	var userPoints []map[string]interface{}
	for rows.Next() {
		var addr, weekStart, weekEnd string
		var sharedPoints, totalPoints float32
		if err := rows.Scan(&addr, &weekStart, &weekEnd, &sharedPoints, &totalPoints); err != nil {
			log.Fatalf("Error scanning row: %s", err)
			return nil
		}
		userPoints = append(userPoints, map[string]interface{}{
			"address":       addr,
			"week_start":    weekStart,
			"week_end":      weekEnd,
			"shared_points": sharedPoints,
			"total_points":  totalPoints,
		})
	}
	return userPoints
}
