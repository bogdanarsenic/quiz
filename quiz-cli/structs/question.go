package structs

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
