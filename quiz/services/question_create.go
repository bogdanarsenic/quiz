package service

import (
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) CreateQuestion(ctx *gin.Context, request *models.CreateQuestionRequest) (*models.Question, error) {

	question := &models.Question{}
	questions := q.database.ListQuestions()
	question.ID = len(*questions) + 1
	q.database.AddQuestion(question)
	return question, nil
}
