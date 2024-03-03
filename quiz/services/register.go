package service

import (
	"errors"
	"net/http"
	"quiz/quiz/models"
	"quiz/quiz/utils"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) Register(ctx *gin.Context, request *models.CreateUserRequest) (string, error) {
	user := &models.User{}
	if _, found := q.database.GetUser(request.Email); found {
		return "", errors.New("There is already an user with this email!")
	}

	if err := user.HashPassword(request.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Hashing Password"})
		return "", errors.New("Error Hashing Password")
	}
	user.Email = request.Email
	user.Score = 0
	user.TookQuiz = false
	user.RoleID = 2
	q.database.AddUser(user)

	// Generate a JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", errors.New("Error generating token")
	}

	return token, nil
}
