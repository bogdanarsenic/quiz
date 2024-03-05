package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Config client config type
type Config struct {
	BasePath     string
	QuestionPath string
	UserPath     string
	LoginPath    string
	RegisterPath string
	MethodGet    string
	MethodPost   string
	MethodPatch  string
	MethodDelete string
	Token        string
}

// QuizClient provides Quiz operations over HTTP
type QuizClient struct {
	HTTPClient HTTPClient
	Cfg        *Config
}

var clientReadAll = ioutil.ReadAll

// NewQuizClient creates an Quiz client
func NewQuizClient(cfg func() *Config) (*QuizClient, error) {
	config := cfg()

	qc := QuizClient{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 1000,
				MaxConnsPerHost:     1000,
			},
			Timeout: 30 * time.Second,
		},
		Cfg: config,
	}

	return &qc, nil
}

func (q *QuizClient) getEndpointURL(path string) (*url.URL, error) {

	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("error building URL, %s is not a valid URL: %s", path, err)
	}

	endpointURL, err := url.Parse(q.Cfg.BasePath)
	if err != nil {
		return nil, fmt.Errorf("error in config, (%s) is not a valid URL: %s", q.Cfg.BasePath, err)
	}
	quizURL := endpointURL.ResolveReference(pathURL)

	return quizURL, nil
}

func (q *QuizClient) getQuestionURL(id string) (*url.URL, error) {
	endPointURL, err := q.getEndpointURL(q.Cfg.QuestionPath)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	questionURL, err := endPointURL.Parse(fmt.Sprint(id))

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return questionURL, nil
}

func (q *QuizClient) getUserURL(email string) (*url.URL, error) {
	endPointURL, err := q.getEndpointURL(q.Cfg.UserPath)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	userURL, err := endPointURL.Parse(fmt.Sprint(email))
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return userURL, nil
}

func (f *QuizClient) sendRequest(method string, requestURL *url.URL, body io.Reader) ([]byte, error) {

	requestURL.Path = requestURL.EscapedPath()
	request, err := http.NewRequest(method, requestURL.String(), body)

	if err != nil {
		return nil, fmt.Errorf("error preparing request to %s: %s", requestURL, err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE3MDk1Mjk3MjEsImlhdCI6MTcwOTUyNzkyMSwiaWQiOiJhZG1pbkBnbWFpbC5jb20iLCJyb2xlIjoxfQ.sOw0DyOucK2XTPh1JW6XFr9TUiXMnmERUpOfcU_pjt4"))

	response, err := f.HTTPClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error from HTTPClient request: %s", err)
	}

	var respBytes []byte
	if response.Body != nil {
		defer func() { _ = response.Body.Close() }()

		respBytes, err = clientReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body (status was %s): %s", response.Status, err)
		}
	}

	if !successful(response.StatusCode) {
		return nil, fmt.Errorf("response status was %s, response: %s", response.Status, respBytes)
	}
	return respBytes, nil
}

func successful(statusCode int) bool {
	return http.StatusOK <= statusCode && statusCode < http.StatusMultipleChoices
}
