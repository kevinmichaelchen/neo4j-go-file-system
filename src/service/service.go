package service

type Error struct {
	HttpCode     int
	ErrorMessage string
	Error        error
}
