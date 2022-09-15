package customErrors

type RateNotInCacheError struct {
	*CustomError
}

func CreateRateNotInCacheError(message string) *RateNotInCacheError {
	return &RateNotInCacheError{&CustomError{errorMessage: message}}
}
