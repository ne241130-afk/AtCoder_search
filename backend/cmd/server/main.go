package main

import (
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "github.com/ne241130/atcoder-learning-hub/backend/internal/handler"
)

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},

		MaxAge: 12 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend is running!",
		})
	})

	r.GET("/api/problems", handler.GetProblems)

	r.Run(":8080")
}