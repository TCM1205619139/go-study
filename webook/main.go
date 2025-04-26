package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/incr", func(c *gin.Context) {
		c.JSON(200, gin.H{"count": 1111})
	})
	r.Run(":8080")
}
