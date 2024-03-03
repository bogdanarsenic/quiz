package rest

import (
	definition "quiz/quiz/services/definition"

	"github.com/gin-gonic/gin"
)

type QuizController struct {
	service definition.Quiz

	// validator binding.StructValidator
}

func NewQuizController(
	service definition.Quiz,
	router *gin.RouterGroup,
) *QuizController {
	c := &QuizController{
		service: service,
		// validator: validators.NewValidator(),
	}

	// // Questions
	router.GET("/questions", c.ListQuestions)
	router.GET("/questions/:id", c.GetQuestion)

	return c
}

func NewAdminController(
	service definition.Quiz,
	router *gin.RouterGroup,
) *QuizController {
	c := &QuizController{
		service: service,
	}

	//questions
	router.POST("/questions", c.CreateQuestion)
	router.GET("/questions", c.ListQuestions)
	router.GET("/questions/:id", c.GetQuestion)
	router.PATCH("/questions/:id", c.UpdateQuestion)
	router.DELETE("/questions/:id", c.DeleteQuestion)

	// Users
	router.POST("/users", c.CreateUser)
	router.GET("/users", c.ListUsers)
	router.GET("/users/:id", c.GetUser)
	router.PATCH("/users/:id", c.UpdateUser)
	router.DELETE("/users/:id", c.DeleteUser)
	return c
}

func NewRegisterController(
	service definition.Quiz,
	router *gin.RouterGroup,
) *QuizController {
	c := &QuizController{
		service: service,
	}

	router.POST("/login", c.Login)
	router.POST("/register", c.Register)

	return c
}
