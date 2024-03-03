package rest

import (
	"fmt"
	"net/http"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

// Function for registering a new user (for demonstration purposes)
func (c *QuizController) Register(ctx *gin.Context) {

	user := &models.CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	token, err := c.service.Register(ctx, user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"token": token, "username": user.Email, "message": "Successfully registered"})
}
