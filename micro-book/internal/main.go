package main

import (
	"micro-book/internal/web"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	user := &web.UserHandler{}
	user.RegisterRoutes(server.Group("/user"))

	server.Run(":8080")
}
