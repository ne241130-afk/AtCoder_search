package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/database"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/handler"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/mock"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := mock.SeedProblems(); err != nil {
		log.Fatalf("failed to seed problems: %v", err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
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
	r.GET("/api/problems/:id", handler.GetProblemByID)
	r.GET("/api/problems/:id/status", handler.GetProblemStatus)
	r.PATCH("/api/problems/:id/solved", handler.PatchProblemSolved)
	r.PATCH("/api/problems/:id/favorite", handler.PatchProblemFavorite)
	r.PATCH("/api/problems/:id/memo", handler.PatchProblemMemo)
	r.GET("/api/tags", handler.GetTags)
	r.POST("/api/tags", handler.CreateTag)
	r.DELETE("/api/tags/:id", handler.DeleteTag)
	r.POST("/api/problems/:id/tags", handler.AttachProblemTag)
	r.DELETE("/api/problems/:id/tags/:tagId", handler.DetachProblemTag)
	r.Run(":8080")
}
