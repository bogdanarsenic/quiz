package error

type CustomErrorWrapper struct {
	Message string `json:"message"` // Human readable message for clients
	Code    int    `json:"-"`       // HTTP Status code. We use `-` to skip json marshaling.
	Err     error  `json:"-"`       // The original error. Same reason as above.
}

func (err CustomErrorWrapper) Error() string {
	return err.Message
}

func NewErrorWrapper(code int, err error, message string) error {
	return CustomErrorWrapper{
		Message: message,
		Code:    code,
		Err:     err,
	}
}
