package authentication

import (
	"duval/internal/pkg/user/authorization"
	"github.com/gin-gonic/gin"
)

func IsStudent(ctx *gin.Context) (ret bool) {
	tok, err := GetTokenDataFromContext(ctx)
	if err != nil {
		return false
	}

	return authorization.IsUserStudent(tok.UserId)
}

func IsParent(ctx *gin.Context) (ret bool) {
	tok, err := GetTokenDataFromContext(ctx)
	if err != nil {
		return false
	}

	return authorization.IsUserParent(tok.UserId)
}

func IsTutor(ctx *gin.Context) (ret bool) {
	tok, err := GetTokenDataFromContext(ctx)
	if err != nil {
		return false
	}

	return authorization.IsUserTutor(tok.UserId)
}

func IsProfessor(ctx *gin.Context) (ret bool) {
	tok, err := GetTokenDataFromContext(ctx)
	if err != nil {
		return false
	}

	return authorization.IsUserProfessor(tok.UserId)
}
