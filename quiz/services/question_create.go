package service

import (
	"errors"
	"net/http"
	db "quiz/quiz/database"
	err "quiz/quiz/errors"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (t Quiz) CreateQuestion(ctx *gin.Context, request *models.Question) (*models.Question, error) {

	question := &models.Question{}

	if _, found := db.GlobalDB.GetQuestion(request.ID); found {
		err := err.NewErrorWrapper(http.StatusUnauthorized, errors.New("There is already a question with this id!"), "Can't perform the action!")
		return nil, err
	}

	db.GlobalDB.AddQuestion(*question)
	return question, nil
}
