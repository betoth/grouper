package middleware

import (
	"fmt"
	"grouper/adapter/input/controller/response"
	util "grouper/application/util/secutiry"
	"grouper/config/rest_errors"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		const bearer_schema = "Bearer "
		token := header[len(bearer_schema):]
		fmt.Println(token)
		auth := util.NewJwtToken().ValidateToken(token)
		if !auth {
			rest_errors.NewInternalServerError("error token")
			response.JSON(w, http.StatusUnauthorized, "error token")
			return

		}
		next(w, r)
	}

}
