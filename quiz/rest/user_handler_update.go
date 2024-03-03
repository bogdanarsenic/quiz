package rest

import (
	"net/http"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) UpdateUser(ctx *gin.Context) {
	user := &models.User{}

	userID := ctx.Param("email")
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	updatedUser, err := c.service.UpdateUser(ctx, userID, user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}
