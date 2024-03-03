package database

import (
	models "quiz/quiz/models"
	"sync"
)

// Database struct to hold key-value pairs
type Database struct {
	users     map[string]models.User
	questions map[int]models.Question
	mu        sync.RWMutex
}

var GlobalDB *Database

// NewDatabase initializes a new instance of the in-memory database
func NewDatabase() *Database {
	return &Database{
		users:     make(map[string]models.User),
		questions: make(map[int]models.Question),
	}
}

// Set inserts or updates a key-value pair in the database
func (db *Database) AddUser(user models.User) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.ID] = user
}

// Get retrieves the value for a given key from the database
func (db *Database) GetUser(userID string) (models.User, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	user, ok := db.users[userID]
	return user, ok
}

// Update retrieves the value for a given key from the database
func (db *Database) UpdateUser(user models.User, id string) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	db.users[id] = user
}

func (db *Database) ListUsers() []models.User {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var users []models.User
	for _, value := range db.users {
		users = append(users, value)
	}
	return users
}

// Delete removes a key-value pair from the database
func (db *Database) DeleteUser(userID string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.users, userID)
}

// Set inserts or updates a key-value pair in the database
func (db *Database) AddQuestion(question models.Question) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.questions[question.ID] = question
}

// Get retrieves the value for a given key from the database
func (db *Database) GetQuestion(questionID int) (models.Question, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	question, ok := db.questions[questionID]
	return question, ok
}

func (db *Database) ListQuestions() []models.Question {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var questions []models.Question
	for _, value := range db.questions {
		questions = append(questions, value)
	}
	return questions
}

// Get retrieves the value for a given key from the database
func (db *Database) UpdateQuestion(question models.Question, id int) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	db.questions[id] = question
}

// Delete removes a key-value pair from the database
func (db *Database) DeleteQuestion(questionID int) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.questions, questionID)
}
