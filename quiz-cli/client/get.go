package client

import (
	"encoding/json"
	"fmt"
	models "quiz/quiz-cli/structs"
)

// GetQuestion Retrieve a question
func (q *QuizClient) GetQuestion(id string) (*models.GetQuestion, error) {

	questionURL, err := q.getQuestionURL(id)
	if err != nil {
		return nil, fmt.Errorf("error preparing request: %s", err)
	}

	resp, err := q.sendRequest(q.Cfg.MethodGet, questionURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error in response: %s", err)
	}

	result := &models.GetQuestion{}

	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, fmt.Errorf("eror reading response body: %s", err)
	}

	return result, nil
}

func (f *QuizClient) GetUser(shareID string) (*models.User, error) {

	userURL, err := f.getUserURL(shareID)
	if err != nil {
		return nil, fmt.Errorf("error preparing request: %s", err)
	}

	resp, err := f.sendRequest(f.Cfg.MethodGet, userURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error in response: %s", err)
	}

	result := &models.User{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		return nil, fmt.Errorf("eror reading response body: %s", err)
	}

	return result, nil
}
