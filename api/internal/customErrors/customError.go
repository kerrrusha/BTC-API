package customErrors

type CustomError struct {
	errorMessage string
}

func (err *CustomError) GetMessage() string {
	return err.errorMessage
}

func CreateCustomError(message string) *CustomError {
	return &CustomError{errorMessage: message}
}
