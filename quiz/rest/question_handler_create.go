package rest

import (
	"net/http"
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) CreateQuestion(ctx *gin.Context) {
	req := &models.Question{}

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	question, err := c.service.CreateQuestion(ctx, req)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Question created successfully", "question": question})
}
