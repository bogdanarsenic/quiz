package service

import (
	db "quiz/quiz/database"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (t Quiz) ListUsers(ctx *gin.Context) (*[]models.User, error) {
	users := db.GlobalDB.ListUsers()

	return &users, nil
}
