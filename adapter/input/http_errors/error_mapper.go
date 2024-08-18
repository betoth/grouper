package http_errors

import (
	"errors"
	"grouper/adapter/input/response"
	appErrors "grouper/application/errors"
	bsnErrors "grouper/application/errors"
	"grouper/config/logger"
	"net/http"

	"go.uber.org/zap"
)

// MapAndRespond centraliza o mapeamento de erros para respostas HTTP
func MapAndRespond(w http.ResponseWriter, err error, journey string) {
	logger.Error("Handling error", err, zap.String("journey", journey))

	switch {
	// Tratamento de Business Errors
	case errors.Is(err, appErrors.ErrGroupAlreadyExists):
		response.JSON(w, http.StatusConflict, NewConflictError(err.Error()))
	case errors.Is(err, appErrors.ErrInvalidGroupData):
		response.JSON(w, http.StatusBadRequest, NewBadRequestError(err.Error()))
	case errors.Is(err, bsnErrors.ErrGroupNotFound):
		response.JSON(w, http.StatusBadRequest, NewNotFoundError(err.Error()))

	// Todos os outros erros serão tratados como Internal Server Error (500)
	default:
		response.JSON(w, http.StatusInternalServerError, NewInternalServerError("Internal server error", err))
	}
}
