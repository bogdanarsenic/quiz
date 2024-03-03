package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *QuizController) ListQuestions(ctx *gin.Context) {
	questions, err := c.service.ListQuestions(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, questions)
}
