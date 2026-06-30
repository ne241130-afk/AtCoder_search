package handler

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/repository"
)

func GetTags(c *gin.Context) {
	repo := &repository.ProblemRepository{}
	problems, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to load tags"})
		return
	}

	tagMap := make(map[string]struct{})
	for _, p := range problems {
		for _, t := range p.Tags {
			tagMap[t] = struct{}{}
		}
	}

	tags := make([]string, 0, len(tagMap))
	for tag := range tagMap {
		tags = append(tags, tag)
	}

	sort.Strings(tags)

	c.JSON(http.StatusOK, tags)
}