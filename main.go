package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "active",
		})
	})

	router.GET("/get-internal-rate", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helloo",
		})
	})

	router.GET("/get-external-rate", func(c *gin.Context) {
		c.Header("X-Billing-Operation", "*")
		c.Header("X-Billing-External-Rate", "2")
		c.JSON(200, gin.H{
			"message": "helloo",
		})
	})

	router.Run(":8080")
}
