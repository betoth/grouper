package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
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

func (t *jwtToken) GenerateToken(userID string) (string, error) {
	claim := &Claim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    t.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenS, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", err
	}
	return tokenS, nil
}

func (t *jwtToken) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, isValid := tk.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(t.secretKey), nil
	})

	return err == nil
}
