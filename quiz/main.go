package main

import (
	"log"
	"os"
	db "quiz/quiz/database"
	"quiz/quiz/models"
	server "quiz/quiz/rest"
	service "quiz/quiz/services"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	db := loadDatabase()
	serveApp(db)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() *db.Database {
	globalDb := db.NewDatabase()
	initializeValues(globalDb)
	return globalDb
}

func initializeValues(db *db.Database) {
	user := &models.User{}
	user.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	user.Email = os.Getenv("ADMIN_EMAIL")
	user.RoleID = 1
	db.AddUser(user)

	question := &models.Question{}
	question.ID = 1
	question.Question = "Which is a capital city of Malta?"
	question.Answers = []string{"Valletta", "Birgu", "Mdina"}
	question.Answer = 1
	db.AddQuestion(question)
}

func serveApp(db *db.Database) {
	quiz := service.New(db)

	restEngine := server.New(quiz)
	restEngine.Run(":8080")
}
