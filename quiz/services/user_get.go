package service

import (
	"errors"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) GetUser(ctx *gin.Context, userID string) (*models.User, error) {
	user, found := q.database.GetUser(userID)

	if !found {
		return nil, errors.New("There is no user with this ID!")
	}

	return user, nil
}
