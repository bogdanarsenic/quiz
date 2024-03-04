package rest

import (
	definition "quiz/quiz/services/definition"

	validators "quiz/quiz/models/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type QuizController struct {
	service definition.Quiz

	validator binding.StructValidator
}

func NewQuizController(
	service definition.Quiz,
	router *gin.RouterGroup,
) *QuizController {
	c := &QuizController{
		service:   service,
		validator: validators.NewValidator(),
	}

	// // Questions
	router.GET("/questions", c.ListQuestions)
	router.GET("/questions/:id", c.GetQuestion)
	router.GET("/users", c.ListUsers)
	router.GET("/users/:id", c.GetUser)

	router.POST("/logout", c.Logout)
	return c
}

func NewAdminController(
	service definition.Quiz,
	router *gin.RouterGroup,
) *QuizController {
	c := &QuizController{
		service:   service,
		validator: validators.NewValidator(),
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

	router.POST("/logout", c.Logout)
	return c
}

func NewRegisterController(
	service definition.Quiz,
	router *gin.RouterGroup,
) *QuizController {
	c := &QuizController{
		service:   service,
		validator: validators.NewValidator(),
	}

	router.POST("/login", c.Login)
	router.POST("/register", c.Register)

	return c
}
