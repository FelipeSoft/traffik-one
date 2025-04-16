package jsonwebtoken

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type JsonWebTokenManager struct {
	expiration time.Duration
}

type Claims struct {
	Payload any `json:"payload"`
	jwt.RegisteredClaims
}

func NewJsonWebTokenManager() *JsonWebTokenManager {
	return &JsonWebTokenManager{
		expiration: time.Hour * 1,
	}
}

func (j *JsonWebTokenManager) Sign(payload any) (string, error) {
	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (j *JsonWebTokenManager) Verify(tokenStr string) (bool, error) {
	_, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (j *JsonWebTokenManager) Decode(tokenStr string) (any, error) {
    tokenStr = strings.TrimSpace(tokenStr)
    tokenStr = strings.Trim(tokenStr, `"`)

    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims.Payload, nil
    }

    return nil, errors.New("invalid token claims")
}