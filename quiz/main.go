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
	loadDatabase()
	serveApp()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() {
	db.GlobalDB = db.NewDatabase()
	initializeValues()
}

func initializeValues() {
	user := models.User{}
	user.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	user.ID = os.Getenv("ADMIN_EMAIL")
	user.RoleID = 1
	db.GlobalDB.AddUser(user)

	question := models.Question{}
	question.ID = 1
	question.Question = "Which is a capital city of quiz/quiz?"
	question.Answers = []string{"Valletta", "Birgu", "Mdina"}
	question.Answer = 1
	db.GlobalDB.AddQuestion(question)
}

func serveApp() {
	quiz := service.New()

	restEngine := server.New(quiz)
	restEngine.Run(":8080")
}
