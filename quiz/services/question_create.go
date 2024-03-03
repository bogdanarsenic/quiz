package service

import (
	"errors"
	"net/http"
	err "quiz/quiz/errors"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) CreateQuestion(ctx *gin.Context, request *models.Question) (*models.Question, error) {

	question := &models.Question{}

	if _, found := q.database.GetQuestion(request.ID); found {
		err := err.NewErrorWrapper(http.StatusUnauthorized, errors.New("There is already a question with this id!"), "Can't perform the action!")
		return nil, err
	}

	q.database.AddQuestion(question)
	return question, nil
}
