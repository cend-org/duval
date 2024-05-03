package education

import (
	"context"
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/user/authorization"
)

func SetUserEducationLevel(ctx context.Context, subjectId int) (*model.Education, error) {
	var (
		tok                       *token.Token
		userEducationLevelSubject model.UserEducationLevelSubject
		err                       error
		userLevel                 model.Education
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &userLevel, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &userLevel, errx.UnAuthorizedError
	}

	if !authorization.IsUserStudent(tok.UserId) {
		if !authorization.IsUserProfessor(tok.UserId) {
			return &userLevel, errx.Lambda(errors.New("not a user or a professor"))
		}
	}

	userEducationLevelSubject.SubjectId = subjectId
	userEducationLevelSubject.UserId = tok.UserId

	_, err = database.InsertOne(userEducationLevelSubject)
	if err != nil {
		return &userLevel, errx.DbInsertError
	}

	userLevel, err = GetUserLevel(tok.UserId)
	if err != nil {
		return &userLevel, errx.DbGetError
	}

	return &userLevel, nil
}

func UpdateUserEducationLevel(ctx context.Context, subjectId int) (*model.Education, error) {
	var (
		tok                              *token.Token
		currentUserEducationLevelSubject model.UserEducationLevelSubject
		userEducationLevelSubject        model.UserEducationLevelSubject
		err                              error
		userLevel                        model.Education
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &userLevel, errx.UnAuthorizedError
	}
	if tok.UserId == state.ZERO {
		return &userLevel, errx.UnAuthorizedError
	}

	if !authorization.IsUserStudent(tok.UserId) {
		return &userLevel, errx.Lambda(errors.New("user is not a student"))
	}

	//err = ctx.ShouldBindJSON(&subject)
	//if err != nil {
	//	return &userLevel, errx.ParseError
	//}

	err = database.Get(&currentUserEducationLevelSubject, `SELECT user_education_level_subject.* FROM user_education_level_subject
			WHERE user_education_level_subject.user_id = ?`, tok.UserId)
	if err != nil {
		return &userLevel, errx.DbGetError
	}

	err = RemoveUserEducationLevelSubject(currentUserEducationLevelSubject)
	if err != nil {
		return &userLevel, errx.DbDeleteError
	}

	userEducationLevelSubject.SubjectId = subjectId
	userEducationLevelSubject.UserId = tok.UserId
	_, err = database.InsertOne(userEducationLevelSubject)
	if err != nil {
		return &userLevel, errx.DbInsertError
	}

	userLevel, err = GetUserLevel(tok.UserId)
	if err != nil {
		return &userLevel, errx.DbGetError
	}

	return &userLevel, nil
}

func GetUserSubjects(ctx context.Context) ([]model.Subject, error) {
	var (
		subjects []model.Subject
		tok      *token.Token
		err      error
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return subjects, errx.UnAuthorizedError
	}

	if authorization.IsUserParent(tok.UserId) || authorization.IsUserTutor(tok.UserId) {
		return subjects, errx.UnAuthorizedError
	}

	if authorization.IsUserStudent(tok.UserId) {
		err = database.GetMany(&subjects, `SELECT subject.* 
											FROM subject
											WHERE subject.education_level_id = (SELECT education.id FROM education  JOIN subject ON education.id  =  subject.education_level_id JOIN user_education_level_subject ON subject.id = user_education_level_subject.subject_id
                                   			WHERE user_education_level_subject.user_id = ?)`, tok.UserId)
		if err != nil {
			return subjects, errx.DbGetError
		}
	}

	if authorization.IsUserProfessor(tok.UserId) {
		err = database.GetMany(&subjects, `SELECT subject.* FROM subject
			JOIN user_education_level_subject  ON subject.id = user_education_level_subject.subject_id
			WHERE user_education_level_subject.user_id = ?`, tok.UserId)
		if err != nil {
			return subjects, errx.DbGetError
		}
	}

	return subjects, nil
}

func GetEducation(ctx context.Context) ([]model.Education, error) {
	var (
		err  error
		edus []model.Education
	)

	err = database.Select(&edus, `SELECT * FROM education WHERE id > 0 ORDER BY  created_at`)
	if err != nil {
		return edus, errx.Lambda(err)
	}

	return edus, nil
}

func GetUserEducationLevel(ctx context.Context) (*model.Education, error) {
	var (
		err       error
		tok       *token.Token
		userLevel model.Education
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &userLevel, errx.UnAuthorizedError
	}

	userLevel, err = GetUserLevel(tok.UserId)
	if err != nil {
		return &userLevel, errx.DbGetError

	}
	return &userLevel, nil
}

func GetSchools(ctx context.Context) ([]model.School, error) {
	var schools []model.School
	var err error

	err = database.Select(&schools, `SELECT * FROM school ORDER BY created_at`)
	if err != nil {
		return nil, err
	}

	return schools, err
}

func GetSubjects(ctx context.Context, id int) ([]model.SchoolSubject, error) {
	var subjects []model.SchoolSubject
	var err error

	err = database.Select(&subjects, `SELECT * FROM school_subject WHERE school_number = ?`, id)
	if err != nil {
		return nil, err
	}

	return subjects, err
}

func GetSchool(ctx context.Context, id int) (*model.School, error) {
	var (
		err    error
		school model.School
	)

	err = database.Get(&school, `SELECT * FROM school WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return &school, err
}

/*
	UTILS
*/

func GetUserLevel(UserId int) (educationLevel model.Education, err error) {

	err = database.Get(&educationLevel,
		`SELECT education.* FROM education
				JOIN subject ON education.id  =  subject.education_level_id
				JOIN user_education_level_subject ON subject.id = user_education_level_subject.subject_id
			WHERE user_education_level_subject.user_id = ?`, UserId)
	if err != nil {
		return educationLevel, err
	}

	return educationLevel, err
}

func RemoveUserEducationLevelSubject(userEducationLevelSubject model.UserEducationLevelSubject) (err error) {
	err = database.Delete(userEducationLevelSubject)
	if err != nil {
		return err
	}
	return err
}