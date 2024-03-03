package service

import (
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) ListUsers(ctx *gin.Context) (*[]models.User, error) {
	users := q.database.ListUsers()

	return users, nil
}
