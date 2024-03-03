package service

import (
	db "quiz/quiz/database"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (t Quiz) ListQuestions(ctx *gin.Context) (*[]models.Question, error) {
	questions := db.GlobalDB.ListQuestions()
	return &questions, nil
}
