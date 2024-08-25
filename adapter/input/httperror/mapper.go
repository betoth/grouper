package httperror

import (
	"grouper/adapter/input/response"
	customerror "grouper/application/custom/custom-error"
	"grouper/config/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func ErrorToErrorResponse(w http.ResponseWriter, err error, journey string) {

	if err == nil {
		logger.Error("Error", err, zap.String("journey", journey))
		response.JSON(w, http.StatusInternalServerError, NewInternalServerError("Internal server error", err))
		return
	}

	if buErr, ok := err.(*customerror.BussinesError); ok {
		errCode, err := strconv.Atoi(buErr.Detais.BusinessErrorCode)
		if err != nil {
			logger.Info(buErr.Detais.BusinessErrorDescription, zap.String("journey", journey))
			return
		}
		response.JSON(w, errCode, NewNotFoundError(buErr.Detais.BusinessErrorDescription))
		return
	}

	logger.Error("Error", err, zap.String("journey", journey))
	response.JSON(w, http.StatusInternalServerError, NewInternalServerError("Internal server error", err))
}
