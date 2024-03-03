package service

import (
	"errors"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) GetQuestion(ctx *gin.Context, questionID int) (*models.Question, error) {
	question, found := q.database.GetQuestion(questionID)

	if !found {
		return nil, errors.New("There is no question with this ID!")
	}

	return question, nil
}
