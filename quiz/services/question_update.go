package service

import (
	"errors"
	db "quiz/quiz/database"
	"quiz/quiz/models"
	"reflect"

	"github.com/gin-gonic/gin"
)

func (t Quiz) UpdateQuestion(ctx *gin.Context, questionID int, req *models.Question) (*models.Question, error) {

	question, found := db.GlobalDB.GetQuestion(questionID)
	if !found {
		return nil, errors.New("There is no question with this ID!")
	}

	if req.Answer != question.Answer && req.Answer != 0 {
		question.Answer = req.Answer
	}
	if req.ID != question.ID && req.ID != 0 {
		question.ID = req.ID
	}

	if req.Question != question.Question && req.Question != "" {
		question.Question = req.Question
	}

	if !reflect.DeepEqual(req.Answers, question.Answers) && req.Answers != nil {
		question.Answers = req.Answers
	}

	db.GlobalDB.UpdateQuestion(question, questionID)

	return &question, nil
}
