package service

import (
	"errors"
	db "quiz/quiz/database"
	"quiz/quiz/models"
	"quiz/quiz/utils"

	"github.com/gin-gonic/gin"
)

func (t Quiz) Login(ctx *gin.Context, request *models.CreateUserRequest) (string, error) {

	userDb, found := db.GlobalDB.GetUser(request.Email)

	if !found {
		return "", errors.New("There is no user with this username registered")
	}

	err := userDb.CheckPassword(request.Password)
	if err != nil {
		return "", errors.New("Password incorrect")
	}

	// Generate a JWT token
	token, err := utils.GenerateJWT(userDb)
	if err != nil {
		return "", errors.New("Error generating token")
	}

	return token, nil
}
