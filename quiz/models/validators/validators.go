package validators

import (
	"errors"
	"quiz/quiz/models"
	"reflect"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	validate "github.com/go-playground/validator"
)

type structValidator struct {
	*validate.Validate
}

func (v *structValidator) Engine() interface{} {
	// TODO: customized messaging?. Need to be figured out.
	// The gin-gonic interface requires this method
	return nil
}

func (v *structValidator) ValidateStruct(i interface{}) error {
	return v.Struct(i)
}

func NewValidator() binding.StructValidator {
	v := validate.New()
	v.SetTagName("binding")

	v.RegisterStructValidation(validateUpdateQuestion, models.UpdateQuestionRequest{})
	v.RegisterStructValidation(validateUpdateUser, models.UpdateUserRequest{})

	return &structValidator{v}
}

func checkAtLeastOneNotNil(sl validate.StructLevel, fields ...string) {
	var c = 0
	for _, name := range fields {
		val := sl.Current().FieldByName(name)
		if !val.IsNil() {
			c++
		}
	}
	if c == 0 {
		sl.ReportError(reflect.ValueOf(c), fields[0], sl.Top().Type().Name(), "atleastone", "")
	}
}

func validateUpdateQuestion(sl validator.StructLevel) {
	checkAtLeastOneNotNil(sl, "Question", "Answer", "Answers")
}

func validateUpdateUser(sl validator.StructLevel) {
	checkAtLeastOneNotNil(sl, "Email", "Password", "Score")
}

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
