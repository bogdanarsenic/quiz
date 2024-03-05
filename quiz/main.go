package main

import (
	"encoding/json"
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

func loadFiles(db *db.Database) {
	usersFile, _ := os.ReadFile("users.json")
	users := []models.User{}
	_ = json.Unmarshal([]byte(usersFile), &users)

	for _, user := range users {
		user.HashPassword(user.Password)
		user.RoleID = 2
		db.AddUser(&user)
	}

	questionsFile, _ := os.ReadFile("questions.json")
	questions := []models.Question{}
	_ = json.Unmarshal([]byte(questionsFile), &questions)

	for _, question := range questions {
		db.AddQuestion(&question)
	}
}

func initializeValues(db *db.Database) {
	loadFiles(db)
	user := &models.User{}
	user.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	user.Email = os.Getenv("ADMIN_EMAIL")
	user.RoleID = 1
	db.AddUser(user)
}

func serveApp(db *db.Database) {
	quiz := service.New(db)

	restEngine := server.New(quiz)
	restEngine.Run(":8080")
}
