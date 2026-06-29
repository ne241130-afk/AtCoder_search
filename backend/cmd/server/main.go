package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/handler"
)

func main() {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend is running!",
		})
	})

	r.GET("/api/problems", handler.GetProblems)

	r.Run(":8080")
}