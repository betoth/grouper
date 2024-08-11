package middleware

import (
	"fmt"
	"grouper/adapter/input/response"
	"grouper/application/util/security"
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
			restErr := rest_errors.NewBadRequestError("Missing token")
			logger.Error("Error trying get token", restErr, zap.String("journey", "Auth"))
			response.JSON(w, restErr.Code, restErr)
			return
		}

		auth := security.NewJwtToken().ValidateToken(token)
		if !auth {
			restErr := rest_errors.NewUnauthorizedError("Invalid token")
			logger.Error("Error trying validate token", restErr, zap.String("journey", "Auth"))
			response.JSON(w, restErr.Code, restErr)
			return
		}
		logger.Debug("Finish auth middleware", zap.String("journey", "Auth"))

		next(w, r)
	}
}

func RequestLogger(next http.Handler) http.Handler {
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
