package service

import (
	db "quiz/quiz/database"
)

type QuizApp struct {
	database *db.Database
}

func New(database *db.Database) *QuizApp {
	return &QuizApp{
		database: database,
	}
}
