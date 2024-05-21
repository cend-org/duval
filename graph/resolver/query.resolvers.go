package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"errors"
	"github.com/cend-org/duval/graph/generated"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/academic"
	"github.com/cend-org/duval/pkg/media/cover"
	"github.com/cend-org/duval/pkg/media/cv"
	"github.com/cend-org/duval/pkg/media/profile"
	"github.com/cend-org/duval/pkg/media/video"
	usr "github.com/cend-org/duval/pkg/user"
	"github.com/cend-org/duval/pkg/user/authorization"
	"github.com/cend-org/duval/pkg/user/link"
)

// MyProfile is the resolver for the MyProfile field.
func (r *queryResolver) MyProfile(ctx context.Context) (*model.User, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errors.New("unAuthorized")
	}

	return usr.MyProfile(tok.UserId)
}

// AcademicLevels is the resolver for the AcademicLevels field.
func (r *queryResolver) AcademicLevels(ctx context.Context) ([]model.AcademicLevel, error) {
	return academic.GetAcademicLevels()
}

// AcademicCourses is the resolver for the AcademicCourses field.
func (r *queryResolver) AcademicCourses(ctx context.Context, academicLevelID int) ([]model.AcademicCourse, error) {
	return academic.GetAcademicCourses(academicLevelID)
}

// CoverLetter is the resolver for the CoverLetter field.
func (r *queryResolver) CoverLetter(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := cover.GetProfileLetter(tok.UserId)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// Cv is the resolver for the Cv field.
func (r *queryResolver) Cv(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := cv.GetProfileCv(tok.UserId)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// ProfileImage is the resolver for the ProfileImage field.
func (r *queryResolver) ProfileImage(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := profile.GetProfileImage(tok.UserId)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// VideoPresentation is the resolver for the VideoPresentation field.
func (r *queryResolver) VideoPresentation(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := video.GetProfileVideo(tok.UserId)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// CoverLetterThumb is the resolver for the CoverLetterThumb field.
func (r *queryResolver) CoverLetterThumb(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := cover.GetProfileLetterThumb(tok.UserId)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// CvThumb is the resolver for the CvThumb field.
func (r *queryResolver) CvThumb(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := cv.GetProfileCvThumb(tok.UserId)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// ProfileImageThumb is the resolver for the ProfileImageThumb field.
func (r *queryResolver) ProfileImageThumb(ctx context.Context) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	networkLink, err := profile.GetProfileImageThumb(tok.UserId)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// UserCoverLetter is the resolver for the UserCoverLetter field.
func (r *queryResolver) UserCoverLetter(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := cover.GetProfileLetter(userID)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// UserCv is the resolver for the UserCv field.
func (r *queryResolver) UserCv(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := cv.GetProfileCv(userID)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// UserProfileImage is the resolver for the UserProfileImage field.
func (r *queryResolver) UserProfileImage(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := profile.GetProfileImage(userID)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// UserVideoPresentation is the resolver for the UserVideoPresentation field.
func (r *queryResolver) UserVideoPresentation(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := video.GetProfileVideo(userID)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// UserCoverLetterThumb is the resolver for the UserCoverLetterThumb field.
func (r *queryResolver) UserCoverLetterThumb(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := cover.GetProfileLetterThumb(userID)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// UserCvThumb is the resolver for the UserCvThumb field.
func (r *queryResolver) UserCvThumb(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := cv.GetProfileCvThumb(userID)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// UserProfileImageThumb is the resolver for the UserProfileImageThumb field.
func (r *queryResolver) UserProfileImageThumb(ctx context.Context, userID int) (*string, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if tok.UserId == state.ZERO {
		return nil, errx.UnAuthorizedError
	}
	networkLink, err := profile.GetProfileImageThumb(userID)
	if err != nil {
		return nil, err
	}
	return &networkLink, nil
}

// UserAcademicLevels is the resolver for the UserAcademicLevels field.
func (r *queryResolver) UserAcademicLevels(ctx context.Context) ([]model.AcademicLevel, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	academicLevel, err := academic.GetUserAcademicLevels(tok.UserId)
	if err != nil {
		return nil, err
	}
	return academicLevel, nil
}

// StudentAcademicLevel is the resolver for the StudentAcademicLevel field.
func (r *queryResolver) StudentAcademicLevel(ctx context.Context, studentID int) ([]model.AcademicLevel, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !link.IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	academicLevel, err := academic.GetUserAcademicLevels(studentID)
	if err != nil {
		return nil, err
	}
	return academicLevel, nil
}

// UserPreferences is the resolver for the UserPreferences field.
func (r *queryResolver) UserPreferences(ctx context.Context, studentID int) (*model.UserAcademicCoursePreference, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	if !link.IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	course, err := academic.GetPreferences(studentID)
	if err != nil {
		return nil, errx.MissingPreferenceError
	}
	return &course, nil
}

// Preferences is the resolver for the Preferences field.
func (r *queryResolver) Preferences(ctx context.Context) (*model.UserAcademicCoursePreference, error) {
	var (
		tok *token.Token
		err error
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	course, err := academic.GetPreferences(tok.UserId)
	if err != nil {
		return nil, errx.MissingPreferenceError
	}
	return &course, nil
}

// UserProfile is the resolver for the UserProfile field.
func (r *queryResolver) UserProfile(ctx context.Context, userID int) (*model.User, error) {
	var tok *token.Token
	var err error

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}
	if tok.UserId == state.ZERO {
		return nil, errx.UnknownUserError
	}

	return usr.MyProfile(userID)
}

// SuggestTutor is the resolver for the SuggestTutor field.
func (r *queryResolver) SuggestTutor(ctx context.Context, studentID int) (*model.User, error) {
	var (
		tok  *token.Token
		err  error
		user model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !link.IsStudentParentLinked(tok.UserId, studentID) {
		return nil, errx.UlError
	}

	user, err = academic.GetTutorWithPreferredCourse(studentID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// SuggestTutorToUser is the resolver for the SuggestTutorToUser field.
func (r *queryResolver) SuggestTutorToUser(ctx context.Context) (*model.User, error) {
	var (
		tok  *token.Token
		err  error
		user model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	user, err = academic.GetTutorWithPreferredCourse(tok.UserId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ProfessorStudent is the resolver for the ProfessorStudent field.
func (r *queryResolver) ProfessorStudent(ctx context.Context, keyWord string) ([]model.User, error) {
	var (
		tok  *token.Token
		err  error
		user []model.User
	)

	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return nil, errx.UnAuthorizedError
	}

	if !authorization.IsUserProfessor(tok.UserId) {
		return nil, errx.UnAuthorizedError
	}

	user, err = link.GetProfessorStudent(keyWord)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
