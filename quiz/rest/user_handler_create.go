package rest

import (
	"net/http"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) CreateUser(ctx *gin.Context) {
	req := &models.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	user, err := c.service.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
