package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ne241130/atcoder-learning-hub/backend/internal/mock"
)

func GetProblems(c *gin.Context) {
	c.JSON(http.StatusOK, mock.Problems)
}