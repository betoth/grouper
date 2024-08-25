package customerror

type (
	BusinessErrorDetais struct {
		BusinessErrorCode        string
		BusinessErrorDescription string
	}

	BussinesErrorType string
	BussinesError     struct {
		Detais BusinessErrorDetais
	}
)

var (
	_ CustomError = (*BussinesError)(nil)
)

const (
	ErrorType_Business ErrorType = "BUSSINESS_ERROR"
)

var (
	BUSSINES_ERROR_GROUP_NOT_FOUND BusinessErrorDetais = BusinessErrorDetais{
		BusinessErrorCode:        "404",
		BusinessErrorDescription: "Group not found",
	}
	BUSSINES_ERROR_USER_NOT_IN_GROUP BusinessErrorDetais = BusinessErrorDetais{
		BusinessErrorCode:        "404",
		BusinessErrorDescription: "User does not belong to the group",
	}
)

var (
	BUSSINES_ERROR_SUBTOPIC_NOT_FOUND BusinessErrorDetais = BusinessErrorDetais{
		BusinessErrorCode:        "404",
		BusinessErrorDescription: "Subtopic not found",
	}
)

var (
	BUSSINES_ERROR_TOPIC_NOT_FOUND BusinessErrorDetais = BusinessErrorDetais{
		BusinessErrorCode:        "404",
		BusinessErrorDescription: "Topic not found",
	}
)

func NewBusinessError(details BusinessErrorDetais) CustomError {
	return &BussinesError{
		Detais: details,
	}
}

func (err BussinesError) Error() string {

	return err.Detais.BusinessErrorDescription
}
