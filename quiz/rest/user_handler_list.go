package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) ListUsers(ctx *gin.Context) {
	users, err := c.service.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
