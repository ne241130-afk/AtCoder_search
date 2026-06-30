package handler

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"

	"github.com/ne241130/atcoder-learning-hub/backend/internal/mock"
)

func GetTags(c *gin.Context) {

	tagMap := make(map[string]struct{})

	for _, p := range mock.Problems {
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