package rest

import (
	"net/http"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

// Function for logging in
func (c *QuizController) Login(ctx *gin.Context) {
	user := &models.CreateUserRequest{}

	// Check user credentials and generate a JWT token
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	token, err := c.service.Login(ctx, user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "username": user.Email, "message": "Successfully logged in"})
}
