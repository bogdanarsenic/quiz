package rest

import (
	"net/http"
	"quiz/quiz/models"
	validators "quiz/quiz/models/validators"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) UpdateQuestion(ctx *gin.Context) {
	question := &models.Question{}

	questionID := ctx.Param("id")
	id, err := validators.IsQuestionIDValid(questionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	updatedQuestion, err := c.service.UpdateQuestion(ctx, id, question)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, updatedQuestion)
}
