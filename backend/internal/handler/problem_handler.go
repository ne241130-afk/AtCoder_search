package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/mock"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/model"
)

func GetProblems(c *gin.Context) {

	tag := strings.ToLower(c.Query("tag"))
	keyword := strings.ToLower(c.Query("keyword"))

	var result []model.Problem

	for _, p := range mock.Problems {

		// タグ検索
		if tag != "" {
			found := false

			for _, t := range p.Tags {
				if strings.ToLower(t) == tag {
					found = true
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