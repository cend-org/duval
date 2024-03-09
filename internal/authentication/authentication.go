package authentication

import (
	"duval/internal/configuration"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	UserId    uint `json:"user_id"`
	UserLevel uint `json:"user_level"`
	jwt.RegisteredClaims
}

func GetTokenDataFromContext(ctx *gin.Context) (tok *Token, err error) {
	tokenString := ctx.GetHeader("Authorization")
	if len(strings.TrimSpace(tokenString)) == 0 {
		return nil, errors.New("bad header value given")
	}

	bearer := strings.Split(tokenString, " ")
	if len(bearer) != 2 {
		return nil, errors.New("incorrectly formatted authorization header")
	}

	return ParseAccessToken(bearer[1]), err
}

func NewAccessToken(claims Token) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(configuration.App.TokenSecret))
}

func NewRefreshToken(claims jwt.RegisteredClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(configuration.App.TokenSecret))
}

func ParseAccessToken(accessToken string) *Token {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &Token{}, func(token *jwt.Token) (any, error) {
		return []byte(configuration.App.TokenSecret), nil
	})

	return parsedAccessToken.Claims.(*Token)
}

func ParseRefreshToken(refreshToken string) *jwt.RegisteredClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(configuration.App.TokenSecret), nil
	})

	return parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
}
