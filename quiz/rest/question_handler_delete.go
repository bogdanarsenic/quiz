package rest

import (
	"net/http"

	validators "quiz/quiz/models/validators"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) DeleteQuestion(ctx *gin.Context) {
	questionID := ctx.Param("id")

	id, err := validators.IsQuestionIDValid(questionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err = c.service.DeleteQuestion(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}
