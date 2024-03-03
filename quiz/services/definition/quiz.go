package definition

import (
	"quiz/quiz/models"

	"github.com/gin-gonic/gin"
)

type Quiz interface {
	ListQuestions(c *gin.Context) (*[]models.Question, error)
	GetQuestion(c *gin.Context, questionID int) (*models.Question, error)
	CreateQuestion(c *gin.Context, req *models.Question) (*models.Question, error)
	UpdateQuestion(c *gin.Context, questionID int, req *models.Question) (*models.Question, error)
	DeleteQuestion(c *gin.Context, questionID int) (*models.Question, error)

	ListUsers(c *gin.Context) (*[]models.User, error)
	GetUser(c *gin.Context, questionID string) (*models.User, error)
	CreateUser(c *gin.Context, req *models.CreateUserRequest) (*models.User, error)
	UpdateUser(c *gin.Context, userID string, req *models.User) (*models.User, error)
	DeleteUser(c *gin.Context, userID string) (*models.User, error)

	Login(c *gin.Context, req *models.CreateUserRequest) (string, error)
	Register(c *gin.Context, req *models.CreateUserRequest) (string, error)
}
