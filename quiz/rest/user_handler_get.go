package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) GetUser(ctx *gin.Context) {
	userID := ctx.Param("email")

	user, err := c.service.GetUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
