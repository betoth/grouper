package middleware

import (
	"grouper/adapter/input/controller/response"
	util "grouper/application/util/secutiry"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := getFormatedToken(r)

		if token == "" {
			err := rest_errors.NewBadRequestError("Missing token")
			logger.Error("Error trying get token", err, zap.String("journey", "Auth"))
			response.JSON(w, err.Code, err)
			return
		}

		auth := util.NewJwtToken().ValidateToken(token)
		if !auth {
			err := rest_errors.NewUnauthorizedError("Invalid token")
			logger.Error("Error trying validate token", err, zap.String("journey", "Auth"))
			response.JSON(w, err.Code, err)
			return
		}
		next(w, r)
	}
}

func getFormatedToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	parts := strings.Split(token, " ")
	if len(parts) == 2 {
		return parts[1]
	}

	return ""
}
