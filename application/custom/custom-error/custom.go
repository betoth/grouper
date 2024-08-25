package customerror

type ErrorType string

type CustomError interface {
	Error() string
}
