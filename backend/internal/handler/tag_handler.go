package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/repository"
)

var (
	tagRepo        = &repository.TagRepository{}
	problemTagRepo = &repository.ProblemTagRepository{}
)

func GetTags(c *gin.Context) {
	tags, err := tagRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func CreateTag(c *gin.Context) {
	var payload struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil || payload.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	tag, err := tagRepo.Create(payload.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tag)
}

func DeleteTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag id"})
		return
	}
	if err := tagRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func AttachProblemTag(c *gin.Context) {
	problemID := c.Param("id")
	var payload struct {
		TagID uint `json:"tag_id"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil || payload.TagID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tag_id is required"})
		return
	}
	if err := problemTagRepo.AttachTag(problemID, payload.TagID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func DetachProblemTag(c *gin.Context) {
	problemID := c.Param("id")
	tagID, err := strconv.ParseUint(c.Param("tagId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag id"})
		return
	}
	if err := problemTagRepo.DetachTag(problemID, uint(tagID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
