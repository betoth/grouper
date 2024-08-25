package customerror

type (
	ApplicationErrorType string

	ApplicationError struct {
		AppErrorType ApplicationErrorType
	}
)

var (
	_ CustomError = (*ApplicationError)(nil)
)

const (
	APP_ERROR_TYPE ApplicationErrorType = "APP_ERROR"
)

func NewApplicationError(errType ApplicationErrorType) CustomError {
	return &ApplicationError{
		AppErrorType: errType,
	}
}

func (err ApplicationError) Error() string {
	return string(err.AppErrorType)
}

func (err ApplicationError) isErrorType(t ApplicationErrorType) bool {
	return err.AppErrorType == t
}

func ContaisAppErrorType(list []ApplicationErrorType, t ApplicationErrorType) bool {
	typeMap := make(map[ApplicationErrorType]bool)
	for _, errType := range list {
		_, ok := typeMap[errType]
		if !ok {
			typeMap[errType] = true
		}
	}

	_, exists := typeMap[t]
	return exists
}
