package main

import (
	"github.com/gin-gonic/gin"
	"trading-ace/config"
	"trading-ace/db"
)

func main() {
	config.LoadConfig()
	db.InitDB()

	r := gin.Default()
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
