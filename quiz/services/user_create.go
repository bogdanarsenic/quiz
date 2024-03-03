package service

import (
	"errors"
	db "quiz/quiz/database"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (t Quiz) CreateUser(ctx *gin.Context, request *models.CreateUserRequest) (*models.User, error) {

	user := &models.User{}
	if _, found := db.GlobalDB.GetUser(request.Email); found {
		return nil, errors.New("There is already an user with this email!")
	}

	if err := user.HashPassword(request.Password); err != nil {
		return nil, errors.New("Error Hashing Password")
	}
	user.ID = request.Email
	user.Score = 0
	user.TookQuiz = false
	db.GlobalDB.AddUser(*user)
	return user, nil
}
