package mark

import (
	"duval/internal/authentication"
	"duval/internal/pkg/user"
	"duval/internal/pkg/user/authorization"
	"duval/internal/utils"
	"duval/internal/utils/errx"
	"duval/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserMark struct {
	Id            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	UserId        uint       `json:"user_id"`
	AuthorId      uint       `json:"author_id"`
	AuthorComment string     `json:"author_comment"`
	AuthorMark    uint       `json:"author_mark"`
}

func RateUser(ctx *gin.Context) {
	var (
		tok         *authentication.Token
		studentMark UserMark
		err         error
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	if authorization.IsUserStudent(tok.UserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	if authorization.IsUserParent(tok.UserId) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}

	err = ctx.ShouldBindJSON(&studentMark)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.ParseError,
		})
		return
	}

	studentMark.AuthorId = tok.UserId
	err = SetUserMark(studentMark)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbInsertError,
		})
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}

func GetUserAverageMark(ctx *gin.Context) {

	ctx.AbortWithStatus(http.StatusOK)
}

func GetUserMarkComment(ctx *gin.Context) {
	var (
		tok         *authentication.Token
		currentUser user.User
		err         error
		mark        UserMark
		userId      uint
		authorId    uint
	)
	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.UnAuthorizedError,
		})
		return
	}
	err = ctx.ShouldBindJSON(&currentUser)

	if authorization.IsUserStudent(tok.UserId) {
		userId = tok.UserId
		authorId = currentUser.Id
	}

	if !authorization.IsUserStudent(tok.UserId) {
		userId = currentUser.Id
		authorId = tok.UserId
	}

	err = database.Get(&mark, `SELECT user_mark.author_comment as 'author_comment' FROM user_mark WHERE user_mark.user_id = ? AND user_mark.author_id`, userId, authorId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DbGetError,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, mark)
}

/*
	UTILS
*/

func SetUserMark(userMark UserMark) (err error) {
	_, err = database.InsertOne(userMark)
	if err != nil {
		return err
	}
	return nil
}
