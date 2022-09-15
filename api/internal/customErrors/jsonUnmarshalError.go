package customErrors

type JsonUnmarshalError struct {
	*CustomError
}

func CreateJsonUnmarshalError(message string) *JsonUnmarshalError {
	return &JsonUnmarshalError{&CustomError{errorMessage: message}}
}
