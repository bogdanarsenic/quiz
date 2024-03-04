package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	_, err := c.service.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
