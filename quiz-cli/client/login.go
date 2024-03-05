package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"quiz/quiz/models"
)

func (q *QuizClient) Login(user models.CreateUserRequest) (string, error) {

	// questionURL, err := q.getQuestionURL(id)
	// if err != nil {
	// 	return nil, fmt.Errorf("error preparing request: %s", err)
	// }

	questionURL, _ := url.Parse("http://localhost:8080/login")

	resp, err := q.sendRequest(q.Cfg.MethodPost, questionURL, nil)
	if err != nil {
		return "", fmt.Errorf("error in response: %s", err)
	}

	var result string
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return "", fmt.Errorf("eror reading response body: %s", err)
	}

	return result, nil
}
