package validators

import (
	"errors"
	"regexp"
	"strconv"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func IsPasswordValid(e string) bool {
	passwordRegex := regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`)
	return passwordRegex.MatchString(e)
}

func IsQuestionIDValid(e string) (int, error) {
	value, err := strconv.Atoi(e)
	if err != nil {
		return 0, errors.New("ID must be number!")
	}
	return value, nil
}
