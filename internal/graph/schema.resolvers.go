package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"duval/internal/graph/model"
	"duval/internal/pkg/user"
	"fmt"
)

// ID is the resolver for the id field.
func (r *authorizationResolver) ID(ctx context.Context, obj *model.Authorization) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *authorizationResolver) UserID(ctx context.Context, obj *model.Authorization) (int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// Level is the resolver for the level field.
func (r *authorizationResolver) Level(ctx context.Context, obj *model.Authorization) (int, error) {
	panic(fmt.Errorf("not implemented: Level - level"))
}

// ID is the resolver for the id field.
func (r *codeResolver) ID(ctx context.Context, obj *model.Code) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *codeResolver) UserID(ctx context.Context, obj *model.Code) (int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUserInput, typeArg int) (string, error) {
	return user.Register(&input, &typeArg)
}

// UpdMyProfile is the resolver for the UpdMyProfile field.
func (r *mutationResolver) UpdMyProfile(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	return user.UpdMyProfile(&ctx, &input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.UserLogin) (string, error) {
	return user.Login(&input)
}

// NewPassword is the resolver for the newPassword field.
func (r *mutationResolver) NewPassword(ctx context.Context, input model.NewPassword) (*string, error) {
	return user.NewPassword(&ctx, &input)
}

// RegisterByEmail is the resolver for the registerByEmail field.
func (r *mutationResolver) RegisterByEmail(ctx context.Context, authorizationLevel int, email string) (*string, error) {
	return user.RegisterByEmail(&authorizationLevel, &email)
}

// ID is the resolver for the id field.
func (r *passwordResolver) ID(ctx context.Context, obj *model.Password) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *passwordResolver) UserID(ctx context.Context, obj *model.Password) (int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// UserAuthorizationLink is the resolver for the userAuthorizationLink field.
func (r *queryResolver) UserAuthorizationLink(ctx context.Context, id int) (*model.UserAuthorizationLink, error) {
	panic(fmt.Errorf("not implemented: UserAuthorizationLink - userAuthorizationLink"))
}

// UserAuthorizationLinks is the resolver for the userAuthorizationLinks field.
func (r *queryResolver) UserAuthorizationLinks(ctx context.Context) ([]*model.UserAuthorizationLink, error) {
	panic(fmt.Errorf("not implemented: UserAuthorizationLinks - userAuthorizationLinks"))
}

// GetCode is the resolver for the getCode field.
func (r *queryResolver) GetCode(ctx context.Context) (*model.Code, error) {
	return user.GetCode(&ctx)
}

// VerifyUserEmailValidationCode is the resolver for the VerifyUserEmailValidationCode field.
func (r *queryResolver) VerifyUserEmailValidationCode(ctx context.Context, code int) (int, error) {
	return user.VerifyUserEmailValidationCode(&ctx, code)
}

// SendUserEmailValidationCode is the resolver for the SendUserEmailValidationCode field.
func (r *queryResolver) SendUserEmailValidationCode(ctx context.Context) (*model.User, error) {
	return user.SendUserEmailValidationCode(&ctx)
}

// MyProfile is the resolver for the myProfile field.
func (r *queryResolver) MyProfile(ctx context.Context) (*model.User, error) {
	return user.MyProfile(&ctx)
}

// GetPasswordHistory is the resolver for the getPasswordHistory field.
func (r *queryResolver) GetPasswordHistory(ctx context.Context) ([]*model.Password, error) {
	return user.GetPasswordHistory(&ctx)
}

// ActivateUser is the resolver for the activateUser field.
func (r *queryResolver) ActivateUser(ctx context.Context) (*model.User, error) {
	return user.ActivateUser(&ctx)
}

// ID is the resolver for the Id field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (int, error) {
	panic(fmt.Errorf("not implemented: ID - Id"))
}

// Age is the resolver for the Age field.
func (r *userResolver) Age(ctx context.Context, obj *model.User) (int, error) {
	panic(fmt.Errorf("not implemented: Age - Age"))
}

// ID is the resolver for the id field.
func (r *userAuthorizationLinkResolver) ID(ctx context.Context, obj *model.UserAuthorizationLink) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// LinkType is the resolver for the linkType field.
func (r *userAuthorizationLinkResolver) LinkType(ctx context.Context, obj *model.UserAuthorizationLink) (int, error) {
	panic(fmt.Errorf("not implemented: LinkType - linkType"))
}

// Actors is the resolver for the actors field.
func (r *userAuthorizationLinkResolver) Actors(ctx context.Context, obj *model.UserAuthorizationLink) ([]*model.UserAuthorizationLinkActor, error) {
	panic(fmt.Errorf("not implemented: Actors - actors"))
}

// ID is the resolver for the id field.
func (r *userAuthorizationLinkActorResolver) ID(ctx context.Context, obj *model.UserAuthorizationLinkActor) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserAuthorizationLink is the resolver for the userAuthorizationLink field.
func (r *userAuthorizationLinkActorResolver) UserAuthorizationLink(ctx context.Context, obj *model.UserAuthorizationLinkActor) (*model.UserAuthorizationLink, error) {
	panic(fmt.Errorf("not implemented: UserAuthorizationLink - userAuthorizationLink"))
}

// AuthorizationID is the resolver for the authorizationId field.
func (r *userAuthorizationLinkActorResolver) AuthorizationID(ctx context.Context, obj *model.UserAuthorizationLinkActor) (int, error) {
	panic(fmt.Errorf("not implemented: AuthorizationID - authorizationId"))
}

// Age is the resolver for the age field.
func (r *newUserInputResolver) Age(ctx context.Context, obj *model.NewUserInput, data *int) error {
	panic(fmt.Errorf("not implemented: Age - age"))
}

// Authorization returns AuthorizationResolver implementation.
func (r *Resolver) Authorization() AuthorizationResolver { return &authorizationResolver{r} }

// Code returns CodeResolver implementation.
func (r *Resolver) Code() CodeResolver { return &codeResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Password returns PasswordResolver implementation.
func (r *Resolver) Password() PasswordResolver { return &passwordResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// UserAuthorizationLink returns UserAuthorizationLinkResolver implementation.
func (r *Resolver) UserAuthorizationLink() UserAuthorizationLinkResolver {
	return &userAuthorizationLinkResolver{r}
}

// UserAuthorizationLinkActor returns UserAuthorizationLinkActorResolver implementation.
func (r *Resolver) UserAuthorizationLinkActor() UserAuthorizationLinkActorResolver {
	return &userAuthorizationLinkActorResolver{r}
}

// NewUserInput returns NewUserInputResolver implementation.
func (r *Resolver) NewUserInput() NewUserInputResolver { return &newUserInputResolver{r} }

type authorizationResolver struct{ *Resolver }
type codeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type passwordResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userAuthorizationLinkResolver struct{ *Resolver }
type userAuthorizationLinkActorResolver struct{ *Resolver }
type newUserInputResolver struct{ *Resolver }
