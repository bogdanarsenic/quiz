package service

import (
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) ListQuestions(ctx *gin.Context) (*[]models.Question, error) {
	questions := q.database.ListQuestions()
	return questions, nil
}
