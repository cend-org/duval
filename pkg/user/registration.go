package user

import (
	"database/sql"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/authentication"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/user/authorization"
)

func NewStudent(email string) (*model.BearerToken, error) {
	return register(email, authorization.STUDENT)
}

func NewParent(email string) (*model.BearerToken, error) {
	return register(email, authorization.PARENT)
}

func NewTutor(email string) (*model.BearerToken, error) {
	return register(email, authorization.TUTOR)
}

func NewProfessor(email string) (*model.BearerToken, error) {
	return register(email, authorization.PROF)
}

func register(email string, userType int) (bearer *model.BearerToken, err error) {
	var user model.User
	var T string

	if isValid := utils.IsValidEmail(email); !isValid {
		return nil, errx.InvalidEmailError
	}

	user, err = getUserByEmail(email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errx.SupportError
	}

	if user.Id > state.ZERO {
		return nil, errx.DuplicateUserError
	}

	user.Matricule, err = utils.GenerateMatricule()
	if err != nil {
		return nil, errx.SupportError
	}

	user.Email = email

	user.Id, err = database.InsertOne(user)
	if err != nil {
		return nil, errx.SupportError
	}

	err = authorization.NewUserAuthorization(user.Id, userType)
	if err != nil {
		return nil, errx.SupportError
	}

	T, err = authentication.NewAccessToken(user)
	if err != nil {
		return nil, errx.SupportError
	}

	bearer = &model.BearerToken{
		T: T,
	}

	return bearer, err
}
