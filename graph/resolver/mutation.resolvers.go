package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"errors"
	"fmt"

	"github.com/cend-org/duval/graph/generated"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/pkg/academic"
	"github.com/cend-org/duval/pkg/media/cover"
	"github.com/cend-org/duval/pkg/media/cv"
	"github.com/cend-org/duval/pkg/media/profile"
	"github.com/cend-org/duval/pkg/media/video"
	usr "github.com/cend-org/duval/pkg/user"
	"github.com/cend-org/duval/pkg/user/link"
)

// NewStudent is the resolver for the NewStudent field.
func (r *mutationResolver) NewStudent(ctx context.Context, email string) (*model.BearerToken, error) {
	return usr.NewStudent(email)
}

// NewParent is the resolver for the NewParent field.
func (r *mutationResolver) NewParent(ctx context.Context, email string) (*model.BearerToken, error) {
	return usr.NewParent(email)
}

// NewTutor is the resolver for the NewTutor field.
func (r *mutationResolver) NewTutor(ctx context.Context, email string) (*model.BearerToken, error) {
	return usr.NewTutor(email)
}

// NewProfessor is the resolver for the NewProfessor field.
func (r *mutationResolver) NewProfessor(ctx context.Context, email string) (*model.BearerToken, error) {
	return usr.NewProfessor(email)
}

// NewPassword is the resolver for the NewPassword field.
func (r *mutationResolver) NewPassword(ctx context.Context, password model.PasswordInput) (*bool, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errors.New("unAuthorized")
	}

	return usr.NewPassword(tok.UserId, password)
}

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.BearerToken, error) {
	return usr.Login(email, password)
}

// UpdateMyProfile is the resolver for the UpdateMyProfile field.
func (r *mutationResolver) UpdateMyProfile(ctx context.Context, profile model.UserInput) (*model.User, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errors.New("unAuthorized")
	}

	return usr.UpdMyProfile(tok.UserId, profile)
}

// UpdateProfileAndPassword is the resolver for the UpdateProfileAndPassword field.
func (r *mutationResolver) UpdateProfileAndPassword(ctx context.Context, profile model.UserInput, password model.PasswordInput) (*model.User, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errors.New("unAuthorized")
	}

	return usr.UpdateProfileAndPassword(tok.UserId, profile, password)
}

// NewUserAcademicCourses is the resolver for the NewUserAcademicCourses field.
func (r *mutationResolver) NewUserAcademicCourses(ctx context.Context, courses []*model.UserAcademicCourseInput) (*bool, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	return academic.NewUserAcademicCourses(tok.UserId, courses)
}

// SetUserAcademicLevel is the resolver for the SetUserAcademicLevel field.
func (r *mutationResolver) SetUserAcademicLevel(ctx context.Context, academicLevelID int) (*model.AcademicLevel, error) {
	panic("not implemented")
}

// NewAcademicCoursePreference is the resolver for the NewAcademicCoursePreference field.
func (r *mutationResolver) NewAcademicCoursePreference(ctx context.Context, preferences []model.UserAcademicCoursePreferenceInput) (*model.UserAcademicCoursePreference, error) {
	panic(fmt.Errorf("not implemented: NewAcademicCoursePreference - NewAcademicCoursePreference"))
}

// NewUserAcademicLevel is the resolver for the NewUserAcademicLevel field.
func (r *mutationResolver) NewUserAcademicLevel(ctx context.Context, levels []*int) (*bool, error) {
	panic(fmt.Errorf("not implemented: NewUserAcademicLevel - NewUserAcademicLevel"))
}

// RemoveCoverLetter is the resolver for the RemoveCoverLetter field.
func (r *mutationResolver) RemoveCoverLetter(ctx context.Context) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	status, err = cover.RemoveProfileLetter(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

// RemoveCv is the resolver for the RemoveCv field.
func (r *mutationResolver) RemoveCv(ctx context.Context) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	status, err = cv.RemoveProfileCv(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

// RemoveProfileImage is the resolver for the RemoveProfileImage field.
func (r *mutationResolver) RemoveProfileImage(ctx context.Context) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	status, err = profile.RemoveProfileImage(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

// RemoveVideoPresentation is the resolver for the RemoveVideoPresentation field.
func (r *mutationResolver) RemoveVideoPresentation(ctx context.Context) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	status, err = video.RemoveProfileVideo(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

// NewStudentByParent is the resolver for the NewStudentByParent field.
func (r *mutationResolver) NewStudentByParent(ctx context.Context, email string) (*bool, error) {
	var (
		tok    *token.Token
		err    error
		status bool
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	_, err = usr.NewStudent(email)
	if err != nil {
		return nil, errx.SupportError
	}

	err = link.AddStudent(tok.UserId, email)
	if err != nil {
		return nil, err
	}

	status = true
	return &status, nil
}

// RemoveStudentByParent is the resolver for the RemoveStudentByParent field.
func (r *mutationResolver) RemoveStudentByParent(ctx context.Context, studentID int) (*bool, error) {
	panic(fmt.Errorf("not implemented: RemoveStudentByParent - RemoveStudentByParent"))
}

// UpdateStudentProfileByParent is the resolver for the UpdateStudentProfileByParent field.
func (r *mutationResolver) UpdateStudentProfileByParent(ctx context.Context, profile model.UserInput, studentID int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &status, errx.UnAuthorizedError
	}

	if !link.IsStudentParentLinked(tok.UserId, studentID) {
		return &status, errx.ParentLinkUserError
	}

	err = link.UpdateStudent(studentID, profile)
	if err != nil {
		return &status, err
	}
	return &status, nil
}

// SetStudentAcademicLevelByParent is the resolver for the SetStudentAcademicLevelByParent field.
func (r *mutationResolver) SetStudentAcademicLevelByParent(ctx context.Context, academicLevelID int, studentID int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !link.IsStudentParentLinked(tok.UserId, studentID) {
		return &status, errx.ParentLinkUserError
	}

	err = academic.SetUserAcademicLevel(tok.UserId, studentID, academicLevelID)
	if err != nil {
		return nil, err
	}

	status = true
	return &status, nil
}

// NewStudentAcademicCoursesByParent is the resolver for the NewStudentAcademicCoursesByParent field.
func (r *mutationResolver) NewStudentAcademicCoursesByParent(ctx context.Context, courses []*model.UserAcademicCourseInput, studentID int) (*bool, error) {
	panic(fmt.Errorf("not implemented: NewStudentAcademicCoursesByParent - NewStudentAcademicCoursesByParent"))
}

// NewStudentAcademicCoursesPreferenceByParent is the resolver for the NewStudentAcademicCoursesPreferenceByParent field.
func (r *mutationResolver) NewStudentAcademicCoursesPreferenceByParent(ctx context.Context, coursesPreferences []*model.UserAcademicCoursePreferenceInput, studentID int) (*bool, error) {
	var tok *token.Token
	var err error
	var status bool

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !link.IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.ParentLinkUserError
	}
	_, err = link.NewStudentAcademicCoursesPreferenceByParent(studentID, coursesPreferences)
	if err != nil {
		return nil, errx.Lambda(err)
	}
	status = true

	return &status, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
