package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(sampleMiddleware())
	r.GET("/:id", func(c *gin.Context) {
		log.Println("main logic")
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	r.Run(":8080")
}

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id == 0 {
			c.JSON(400, gin.H{"message": "invalid id"})
			c.Abort()
		}
	}
}