package service

import (
	"errors"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) CreateUser(ctx *gin.Context, request *models.CreateUserRequest) (*models.User, error) {

	user := &models.User{}
	if _, found := q.database.GetUser(request.Email); found {
		return nil, errors.New("There is already an user with this email!")
	}

	if err := user.HashPassword(request.Password); err != nil {
		return nil, errors.New("Error Hashing Password")
	}
	user.Email = request.Email
	user.Score = 0
	user.TookQuiz = false
	q.database.AddUser(user)
	return user, nil
}
