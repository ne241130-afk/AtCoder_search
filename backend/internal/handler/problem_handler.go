package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/mock"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
)

func GetProblems(c *gin.Context) {

	tagsParam := strings.ToLower(c.Query("tags"))

	var selectedTags []string

	if tagsParam != "" {
		selectedTags = strings.Split(tagsParam, ",")
	}
	keyword := strings.ToLower(c.Query("keyword"))

	var result []model.Problem

	for _, p := range mock.Problems {

		// タグ検索
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

		// キーワード検索
		if keyword != "" {

			target := strings.ToLower(
				p.Title + " " + p.Contest,
			)

			if !strings.Contains(target, keyword) {
				continue
			}
		}

		result = append(result, p)
	}

	c.JSON(http.StatusOK, result)
}