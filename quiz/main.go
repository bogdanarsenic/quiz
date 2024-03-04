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

	user2 := &models.User{}
	user2.HashPassword("User123!")
	user2.Email = "user2@gmail.com"
	user2.RoleID = 2
	user2.TookQuiz = false
	db.AddUser(user2)

	user3 := &models.User{}
	user3.HashPassword("User123!")
	user3.Email = "user3@gmail.com"
	user3.RoleID = 2
	user3.TookQuiz = true
	user3.Score = 3
	db.AddUser(user3)

	user4 := &models.User{}
	user4.HashPassword("User123!")
	user4.Email = "user4@gmail.com"
	user4.RoleID = 2
	user4.TookQuiz = true
	user4.Score = 5
	db.AddUser(user4)

	user5 := &models.User{}
	user5.HashPassword("User123!")
	user5.Email = "user5@gmail.com"
	user5.RoleID = 2
	user5.TookQuiz = true
	user5.Score = 1
	db.AddUser(user5)

	question := &models.Question{}
	question.ID = 1
	question.Question = "Which is a capital city of Malta?"
	question.Answers = []string{"Valletta", "Birgu", "Mdina"}
	question.Answer = "Valletta"
	db.AddQuestion(question)

	question2 := &models.Question{}
	question2.ID = 2
	question2.Question = "Which planet is the closest to sun?"
	question2.Answers = []string{"Saturn", "Mercury", "Venus"}
	question2.Answer = "Mercury"
	db.AddQuestion(question2)

	question3 := &models.Question{}
	question3.ID = 3
	question3.Question = "How many timezones Russia have?"
	question3.Answers = []string{"10", "12", "15"}
	question3.Answer = "12"
	db.AddQuestion(question3)

	question4 := &models.Question{}
	question4.ID = 4
	question4.Question = "Who won Wimbledon 2011?"
	question4.Answers = []string{"Rafael Nadal", "Roger Federer", "Novak Djokovic"}
	question4.Answer = "Novak Djokovic"
	db.AddQuestion(question4)

	question5 := &models.Question{}
	question5.ID = 5
	question5.Question = "What is the largest continent?"
	question5.Answers = []string{"Asia", "Europe", "North America"}
	question5.Answer = "Asia"
	db.AddQuestion(question5)
}

func serveApp(db *db.Database) {
	quiz := service.New(db)

	restEngine := server.New(quiz)
	restEngine.Run(":8080")
}
