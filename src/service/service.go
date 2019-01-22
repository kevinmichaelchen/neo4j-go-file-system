package service

type Error struct {
	HttpCode     int
	ErrorMessage string
	Error        error
}

func NewError(httpCode int, errorMessage string, err error) *Error {
	return &Error{
		HttpCode:     httpCode,
		ErrorMessage: errorMessage,
		Error:        err,
	}
}
