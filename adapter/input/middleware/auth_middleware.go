package middleware

import (
	"fmt"
	"grouper/adapter/input/response"
	"grouper/application/util/secutiry"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Init auth middleware", zap.String("journey", "Auth"))
		token := getFormatedToken(r)

		if token == "" {
			err := rest_errors.NewBadRequestError("Missing token")
			logger.Error("Error trying get token", err, zap.String("journey", "Auth"))
			response.JSON(w, err.Code, err)
			return
		}

		auth := secutiry.NewJwtToken().ValidateToken(token)
		if !auth {
			err := rest_errors.NewUnauthorizedError("Invalid token")
			logger.Error("Error trying validate token", err, zap.String("journey", "Auth"))
			response.JSON(w, err.Code, err)
			return
		}
		logger.Debug("Finish auth middleware", zap.String("journey", "Auth"))
		next(w, r)
	}
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest := fmt.Sprintf("%s %s", r.Method, r.RequestURI)
		logger.Debug("Request", zap.String("request_info", logRequest))

		next.ServeHTTP(w, r)
	})
}

func getFormatedToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	parts := strings.Split(token, " ")
	if len(parts) == 2 {
		return parts[1]
	}

	return ""
}
