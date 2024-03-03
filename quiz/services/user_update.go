package service

import (
	"errors"
	"net/http"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (q *QuizApp) UpdateUser(ctx *gin.Context, userID string, req *models.User) (*models.User, error) {

	user, found := q.database.GetUser(userID)
	if !found {
		return nil, errors.New("There is no user with this ID!")
	}

	if req.Password != user.Password && req.Password != "" {
		user.Password = req.Password
		if err := user.HashPassword(user.Password); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Hashing Password"})
			return nil, err
		}
	}
	if req.Email != user.Email && req.Email != "" {
		user.Email = req.Email
	}

	if req.Score != user.Score {
		user.Score = req.Score
	}

	if req.TookQuiz != user.TookQuiz {
		user.TookQuiz = req.TookQuiz
	}

	q.database.UpdateUser(user, userID)
	return user, nil
}
