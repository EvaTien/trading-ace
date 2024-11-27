package main

import (
	"github.com/gin-gonic/gin"
	"trading-ace/config"
	"trading-ace/db"
)

func main() {
	config.Init()

	db.InitDB()
	db.CreateTables()

	r := gin.Default()
	err := r.Run(":8080")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
