package client

import (
	"encoding/json"
	"fmt"
	"quiz/quiz-cli/structs"
)

// GetQuestion Retrieve a question
func (q *QuizClient) ListQuestions() (*[]structs.Question, error) {

	questionURL, err := q.getQuestionURL("")
	resp, err := q.sendRequest(q.Cfg.MethodGet, questionURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error in response: %s", err)
	}

	result := []structs.Question{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("eror reading response body: %s", err)
	}

	return &result, nil
}

func (q *QuizClient) ListUsers() (*[]structs.User, error) {

	userURL, err := q.getUserURL("")
	if err != nil {
		return nil, fmt.Errorf("error preparing request: %s", err)
	}

	resp, err := q.sendRequest(q.Cfg.MethodGet, userURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error in response: %s", err)
	}

	result := []structs.User{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("eror reading response body: %s", err)
	}

	return &result, nil
}
