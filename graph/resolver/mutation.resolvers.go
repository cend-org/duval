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
	"github.com/cend-org/duval/pkg/academic"
	usr "github.com/cend-org/duval/pkg/user"
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
		return nil, errors.New("unAuthorized")
	}

	return academic.NewUserAcademicCourses(tok.UserId, courses)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
