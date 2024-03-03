package service

import (
	"errors"
	db "quiz/quiz/database"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (t Quiz) DeleteUser(ctx *gin.Context, userID string) (*models.User, error) {
	user, found := db.GlobalDB.GetUser(userID)

	if !found {
		return nil, errors.New("There is no user with this ID!")
	}
	db.GlobalDB.DeleteUser(userID)
	return &user, nil
}
