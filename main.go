package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trading-ace/config"
	"trading-ace/db"
)

func main() {
	config.Init()

	db.InitDB()
	db.CreateTables()

	// go utils.GetSwapTransactions()

	r := gin.Default()

	r.GET("/users/:address", getUserByAddress)
	r.GET("/users/:address/points-history", getUserPointsByAddress)

	err := r.Run(":8080")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}

func getUserByAddress(c *gin.Context) {
	address := c.Param("address")

	user := db.GetUser(address)
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func getUserPointsByAddress(c *gin.Context) {
	address := c.Param("address")

	existed := db.UserExisted(address)
	if !existed {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	pointsHistory := db.GetUserPointsByAddress(address)
	c.JSON(http.StatusOK, pointsHistory)
}
