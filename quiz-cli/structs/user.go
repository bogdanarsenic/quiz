package structs

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
