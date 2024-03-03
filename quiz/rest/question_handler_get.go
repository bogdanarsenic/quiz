package rest

import (
	"net/http"
	validators "quiz/quiz/models/validators"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) GetQuestion(ctx *gin.Context) {
	questionID := ctx.Param("id")

	id, err := validators.IsQuestionIDValid(questionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	question, err := c.service.GetQuestion(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, question)
}
