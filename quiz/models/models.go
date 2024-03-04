package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Score    int    `json:"score"`
	TookQuiz bool   `json:"tookQuiz"`
	RoleID   int    `json:"-"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
	Score    int    `json:"score"`
	TookQuiz bool   `json:"tookQuiz"`
}

type UpdateQuestionRequest struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
	Answer   string   `json:"answer"`
}

type CreateQuestionRequest struct {
	Question string   `json:"question" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
	Answer   string   `json:"answer" binding:"required"`
}

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
	Answer   string   `json:"answer"`
}

// HashPassword encrypts user password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword checks user password
// CheckPassword takes a string as a parameter and compares it to the user's encrypted password
// It returns an error if there is an issue comparing the passwords
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
