package service

import (
	"errors"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) DeleteQuestion(ctx *gin.Context, questionID int) (*models.Question, error) {

	question, found := q.database.GetQuestion(questionID)

	if !found {
		return nil, errors.New("There is no user with this ID!")
	}

	q.database.DeleteQuestion(questionID)
	return question, nil
}
