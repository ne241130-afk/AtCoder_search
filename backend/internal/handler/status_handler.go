package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/repository"
)

func GetProblemStatus(c *gin.Context) {
	problemID := c.Param("id")
	repo := &repository.StatusRepository{}
	status, err := repo.GetStatus(problemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to load status"})
		return
	}
	c.JSON(http.StatusOK, status)
}

func PatchProblemSolved(c *gin.Context) {
	problemID := c.Param("id")
	var payload struct {
		Solved bool `json:"solved"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	repo := &repository.StatusRepository{}
	if err := repo.UpdateSolved(problemID, payload.Solved); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update solved"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func PatchProblemFavorite(c *gin.Context) {
	problemID := c.Param("id")
	var payload struct {
		Favorite bool `json:"favorite"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	repo := &repository.StatusRepository{}
	if err := repo.UpdateFavorite(problemID, payload.Favorite); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update favorite"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func PatchProblemMemo(c *gin.Context) {
	problemID := c.Param("id")
	var payload struct {
		Memo string `json:"memo"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	repo := &repository.StatusRepository{}
	if err := repo.UpdateMemo(problemID, payload.Memo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update memo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
