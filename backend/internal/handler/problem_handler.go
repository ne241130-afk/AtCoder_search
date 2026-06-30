package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/repository"
)

func filterProblems(problems []model.Problem, keyword string, selectedTags []string, contestType string, minDifficulty *int, maxDifficulty *int) []model.Problem {
	var result []model.Problem

	for _, p := range problems {
		if len(selectedTags) > 0 {
			found := false
			for _, selectedTag := range selectedTags {
				for _, t := range p.Tags {
					if strings.ToLower(t) == selectedTag {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				continue
			}
		}

		if contestType != "" {
			if !strings.HasPrefix(strings.ToUpper(p.Contest), strings.ToUpper(contestType)) {
				continue
			}
		}

		if minDifficulty != nil && p.Difficulty < *minDifficulty {
			continue
		}

		if maxDifficulty != nil && p.Difficulty > *maxDifficulty {
			continue
		}

		if keyword != "" {
			target := strings.ToLower(p.Title + " " + p.Contest)
			if !strings.Contains(target, keyword) {
				continue
			}
		}

		result = append(result, p)
	}

	return result
}

func GetProblemByID(c *gin.Context) {
	problemID := c.Param("id")
	repo := &repository.ProblemRepository{}

	problem, err := repo.FindByID(problemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "problem not found"})
		return
	}

	c.JSON(http.StatusOK, problem)
}

func GetProblems(c *gin.Context) {
	tagsParam := strings.ToLower(c.Query("tags"))
	var selectedTags []string
	if tagsParam != "" {
		selectedTags = strings.Split(tagsParam, ",")
	}
	keyword := strings.ToLower(c.Query("keyword"))
	contestType := strings.TrimSpace(c.Query("contestType"))

	var minDifficulty *int
	if minParam := c.Query("minDifficulty"); minParam != "" {
		if parsed, err := strconv.Atoi(minParam); err == nil {
			minDifficulty = &parsed
		}
	}

	var maxDifficulty *int
	if maxParam := c.Query("maxDifficulty"); maxParam != "" {
		if parsed, err := strconv.Atoi(maxParam); err == nil {
			maxDifficulty = &parsed
		}
	}

	repo := &repository.ProblemRepository{}
	problems, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to load problems"})
		return
	}

	filtered := filterProblems(problems, keyword, selectedTags, contestType, minDifficulty, maxDifficulty)

	page := 1
	if pageParam := c.Query("page"); pageParam != "" {
		if parsed, err := strconv.Atoi(pageParam); err == nil && parsed > 0 {
			page = parsed
		}
	}

	limit := 20
	if limitParam := c.Query("limit"); limitParam != "" {
		if parsed, err := strconv.Atoi(limitParam); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	start := (page - 1) * limit
	end := start + limit
	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	response := gin.H{
		"items": filtered[start:end],
		"total": len(filtered),
		"page":  page,
		"limit": limit,
	}

	c.JSON(http.StatusOK, response)
}