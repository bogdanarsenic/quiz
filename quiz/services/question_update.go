package service

import (
	"errors"
	"quiz/quiz/models"
	"reflect"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) UpdateQuestion(ctx *gin.Context, questionID int, req *models.Question) (*models.Question, error) {

	question, found := q.database.GetQuestion(questionID)
	if !found {
		return nil, errors.New("There is no question with this ID!")
	}

	if req.Answer != question.Answer && req.Answer != "" {
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

	q.database.UpdateQuestion(question, questionID)

	return question, nil
}
