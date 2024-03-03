package service

import (
	"errors"
	"net/http"
	db "quiz/quiz/database"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (t Quiz) UpdateUser(ctx *gin.Context, userID string, req *models.User) (*models.User, error) {

	user, found := db.GlobalDB.GetUser(userID)
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
	if req.ID != user.ID && req.ID != "" {
		user.ID = req.ID
	}

	if req.Score != user.Score {
		user.Score = req.Score
	}

	if req.TookQuiz != user.TookQuiz {
		user.TookQuiz = req.TookQuiz
	}

	db.GlobalDB.UpdateUser(user, userID)
	return &user, nil
}
