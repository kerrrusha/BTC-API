package customErrors

type RequestFailureError struct {
	*CustomError
}

func CreateRequestFailureError(message string) *RequestFailureError {
	return &RequestFailureError{&CustomError{errorMessage: message}}
}
