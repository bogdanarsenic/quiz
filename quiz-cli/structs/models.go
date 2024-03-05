package models

type User struct {
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

type GetQuestion struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
	Answer   string   `json:"-"`
}

type PromptContent struct {
	ErrorMsg string
	Label    string
}
