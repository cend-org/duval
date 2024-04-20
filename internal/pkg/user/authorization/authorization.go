package authorization

import (
	"duval/internal/graph/model"
	"duval/internal/utils/state"
	"duval/pkg/database"
)

const (
	StudentAuthorizationLevel   = 0
	ParentAuthorizationLevel    = 1
	TutorAuthorizationLevel     = 2
	ProfessorAuthorizationLevel = 3
)

func NewUserAuthorization(userId, authorizationLevel uint) (err error) {
	var (
		auth model.Authorization
	)

	auth.UserId = userId
	auth.Level = authorizationLevel

	_, err = database.InsertOne(auth)
	if err != nil {
		return err
	}
	return err
}

func GetUserAuthorizations(userId uint) (auth []model.Authorization, err error) {
	query := `SELECT a.* FROM authorization a WHERE a.user_id = ?`
	err = database.Select(&auth, query, userId)
	if err != nil {
		return nil, err
	}
	return auth, err
}

func GetUserAuthorization(userId, level uint) (auth model.Authorization, err error) {
	err = database.Get(&auth, `SELECT * FROM authorization WHERE user_id = ? AND level = ?`, userId, level)
	if err != nil {
		return auth, err
	}

	return auth, err
}

func DeleteUserAuthorization(userId, level uint) (err error) {
	var (
		auth model.Authorization
	)

	auth, err = GetUserAuthorization(userId, level)
	if err != nil {
		return err
	}

	err = database.Delete(auth)
	if err != nil {
		return err
	}

	return err
}

func DeleteUserAuthorizations(userId uint) (err error) {
	var (
		authorization model.Authorization
	)
	err = database.Get(&authorization, `SELECT  * FROM authorization WHERE  authorization.user_id= ?`, userId)
	if err != nil {
		return err
	}

	err = database.Delete(authorization)
	if err != nil {
		return err
	}
	return err
}

func isUserHasAuthorizationLevel(userId, authorizationLevel uint) (ret bool) {
	var (
		err  error
		auth model.Authorization
	)

	auth, err = GetUserAuthorization(userId, authorizationLevel)
	if err != nil {
		return false
	}

	return auth.Id > state.ZERO
}

func IsUserStudent(userId uint) (ret bool) {
	return isUserHasAuthorizationLevel(userId, StudentAuthorizationLevel)
}

func IsUserParent(userId uint) (ret bool) {
	return isUserHasAuthorizationLevel(userId, ParentAuthorizationLevel)
}

func IsUserTutor(userId uint) (ret bool) {
	return isUserHasAuthorizationLevel(userId, TutorAuthorizationLevel)
}

func IsUserProfessor(userId uint) (ret bool) {
	return isUserHasAuthorizationLevel(userId, ProfessorAuthorizationLevel)
}
