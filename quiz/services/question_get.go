package service

import (
	"errors"
	db "quiz/quiz/database"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (t Quiz) GetQuestion(ctx *gin.Context, questionID int) (*models.Question, error) {
	question, found := db.GlobalDB.GetQuestion(questionID)

	if !found {
		return nil, errors.New("There is no question with this ID!")
	}

	return &question, nil
}
