package security

import (
	"fmt"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type jwtToken struct {
	secretKey string
	issuer    string
}

func NewJwtToken() *jwtToken {
	return &jwtToken{
		secretKey: "secret-key",
		issuer:    "grouper",
	}
}

type Claim struct {
	Sum string `json:"sum"`
	jwt.StandardClaims
}

func (t *jwtToken) GenerateToken(userID string) (string, *rest_errors.RestErr) {
	claim := &Claim{
		Sum: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    t.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenS, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		logger.Error("Error signing token", err, zap.String("journey", "GenerateToken"))
		return "", rest_errors.NewInternalServerError("Error signing token: " + err.Error())
	}
	logger.Info("Token successfully generated", zap.String("journey", "Login"))
	return tokenS, nil
}

func (t *jwtToken) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, isValid := tk.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", tk.Method)
		}
		return []byte(t.secretKey), nil
	})

	if err != nil {
		logger.Error("Error validating token", err, zap.String("journey", "ValidateToken"))
		return false
	}
	return true
}

func (t *jwtToken) extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	parts := strings.Split(token, " ")

	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		return parts[1]
	}

	logger.Error("Malformed token header", nil, zap.String("journey", "ExtractToken"), zap.String("tokenHeader", token))
	return ""
}

func (t *jwtToken) returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(t.secretKey), nil
}

func (t *jwtToken) ExtractUserID(r *http.Request) (string, *rest_errors.RestErr) {
	tokenString := t.extractToken(r)
	token, err := jwt.Parse(tokenString, t.returnVerificationKey)
	if err != nil {
		logger.Error("Error parsing token", err, zap.String("journey", "ExtractUserID"))
		return "", rest_errors.NewInternalServerError("Error parsing token: " + err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sum"].(string)
		if !ok {
			logger.Error("Invalid claim format", nil, zap.String("journey", "ExtractUserID"))
			return "", rest_errors.NewInternalServerError("Invalid claim format")
		}
		return userID, nil
	}

	logger.Error("Invalid token", nil, zap.String("journey", "ExtractUserID"))
	return "", rest_errors.NewUnauthorizedRequestError("Invalid token")
}
