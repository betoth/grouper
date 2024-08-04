package response

import (
	"encoding/json"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"net/http"

	"go.uber.org/zap"
)

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		response, err := json.Marshal(data)
		if err != nil {
			logger.Error("Error trying marshal data", err, zap.String("journey", "response"))
			rest_errors.NewInternalServerError("")
			return
		}

		if _, err := w.Write(response); err != nil {
			logger.Error("Error trying write response", err, zap.String("journey", "response"))
			rest_errors.NewInternalServerError("")
			return
		}
	}
}
