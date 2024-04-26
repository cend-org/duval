package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/generated"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/pkg/address"
	"github.com/cend-org/duval/internal/pkg/education"
	"github.com/cend-org/duval/internal/pkg/mark"
	"github.com/cend-org/duval/internal/pkg/media"
	"github.com/cend-org/duval/internal/pkg/phone"
	"github.com/cend-org/duval/internal/pkg/planning"
	"github.com/cend-org/duval/internal/pkg/translator"
	usr "github.com/cend-org/duval/internal/pkg/user"
	"github.com/cend-org/duval/internal/pkg/user/link"
	"github.com/cend-org/duval/internal/school"
)

// RegisterWithEmail is the resolver for the RegisterWithEmail field.
func (r *mutationResolver) RegisterWithEmail(ctx context.Context, input string, as int) (*string, error) {
	return usr.RegisterWithEmail(ctx, input, as)
}

// Register is the resolver for the Register field.
func (r *mutationResolver) Register(ctx context.Context, input model.UserInput, as int) (*string, error) {
	return usr.Register(ctx, &input, as)
}

// UpdMyProfile is the resolver for the updMyProfile field.
func (r *mutationResolver) UpdMyProfile(ctx context.Context, input model.UserInput) (*string, error) {
	panic(fmt.Errorf("not implemented: UpdMyProfile - updMyProfile"))
}

// LogIn is the resolver for the LogIn field.
func (r *mutationResolver) LogIn(ctx context.Context, email string, password string) (*string, error) {
	return usr.LogIn(ctx, email, password)
}

// NewPassword is the resolver for the NewPassword field.
func (r *mutationResolver) NewPassword(ctx context.Context, password string) (*bool, error) {
	return usr.NewPassword(ctx, password)
}

// DelPassword is the resolver for the DelPassword field.
func (r *mutationResolver) DelPassword(ctx context.Context, id int) (*bool, error) {
	panic(fmt.Errorf("not implemented: DelPassword - DelPassword"))
}

// NewAsset is the resolver for the NewAsset field.
func (r *mutationResolver) NewAsset(ctx context.Context, asset model.AssetInput) (*model.Asset, error) {
	var (
		as  model.Asset
		err error
	)

	as = model.MapAssetInputToAsset(asset, as)
	_, err = database.InsertOne(as)
	if err != nil {
		return nil, err
	}

	return &as, err
}

// PopulateSchool is the resolver for the PopulateSchool field.
func (r *mutationResolver) PopulateSchool(ctx context.Context) (*bool, error) {
	err := school.Populate()
	if err != nil {
		return nil, err
	}

	done := true
	return &done, err
}

// NewMessage is the resolver for the newMessage field.
func (r *mutationResolver) NewMessage(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	return translator.NewMessage(ctx, &input)
}

// UpdMessage is the resolver for the updMessage field.
func (r *mutationResolver) UpdMessage(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	return translator.UpdMessage(ctx, &input)
}

// DelMessage is the resolver for the delMessage field.
func (r *mutationResolver) DelMessage(ctx context.Context, language int, messageNumber int) (*string, error) {
	return translator.DelMessage(ctx, language, messageNumber)
}

// NewMenu is the resolver for the newMenu field.
func (r *mutationResolver) NewMenu(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	return translator.NewMenu(ctx, &input)
}

// DelMenu is the resolver for the delMenu field.
func (r *mutationResolver) DelMenu(ctx context.Context, menuNumber int) (*string, error) {
	return translator.DelMenu(ctx, menuNumber)
}

// NewMenuItem is the resolver for the newMenuItem field.
func (r *mutationResolver) NewMenuItem(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	return translator.NewMenuItem(ctx, &input)
}

// DelMenuItem is the resolver for the delMenuItem field.
func (r *mutationResolver) DelMenuItem(ctx context.Context, input model.MessageInput) (*string, error) {
	return translator.DelMenuItem(ctx, &input)
}

// NewAddress is the resolver for the newAddress field.
func (r *mutationResolver) NewAddress(ctx context.Context, input model.AddressInput) (*model.Address, error) {
	return address.NewAddress(ctx, &input)
}

// UpdateUserAddress is the resolver for the updateUserAddress field.
func (r *mutationResolver) UpdateUserAddress(ctx context.Context, input model.AddressInput) (*model.Address, error) {
	return address.UpdateUserAddress(ctx, &input)
}

// NewPhoneNumber is the resolver for the newPhoneNumber field.
func (r *mutationResolver) NewPhoneNumber(ctx context.Context, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	return phone.NewPhoneNumber(ctx, &input)
}

// UpdateUserPhoneNumber is the resolver for the updateUserPhoneNumber field.
func (r *mutationResolver) UpdateUserPhoneNumber(ctx context.Context, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	return phone.UpdateUserPhoneNumber(ctx, &input)
}

// CreateUserPlannings is the resolver for the createUserPlannings field.
func (r *mutationResolver) CreateUserPlannings(ctx context.Context, input model.CalendarPlanningInput) (*model.CalendarPlanning, error) {
	return planning.CreateUserPlannings(ctx, &input)
}

// AddUserIntoPlanning is the resolver for the addUserIntoPlanning field.
func (r *mutationResolver) AddUserIntoPlanning(ctx context.Context, calendarID int, selectedUserID int) (*model.CalendarPlanningActor, error) {
	return planning.AddUserIntoPlanning(ctx, calendarID, selectedUserID)
}

// RemoveUserPlannings is the resolver for the removeUserPlannings field.
func (r *mutationResolver) RemoveUserPlannings(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: RemoveUserPlannings - removeUserPlannings"))
}

// RemoveUserFromPlanning is the resolver for the removeUserFromPlanning field.
func (r *mutationResolver) RemoveUserFromPlanning(ctx context.Context, calendarPlanningID int, selectedUserID int) (*string, error) {
	return planning.RemoveUserFromPlanning(ctx, calendarPlanningID, selectedUserID)
}

// SetUserEducationLevel is the resolver for the setUserEducationLevel field.
func (r *mutationResolver) SetUserEducationLevel(ctx context.Context, subjectID int) (*model.Education, error) {
	return education.SetUserEducationLevel(ctx, subjectID)
}

// UpdateUserEducationLevel is the resolver for the updateUserEducationLevel field.
func (r *mutationResolver) UpdateUserEducationLevel(ctx context.Context, subjectID int) (*model.Education, error) {
	return education.UpdateUserEducationLevel(ctx, subjectID)
}

// RateUser is the resolver for the rateUser field.
func (r *mutationResolver) RateUser(ctx context.Context, input model.MarkInput) (*model.Mark, error) {
	return mark.RateUser(ctx, &input)
}

// AddParentToUser is the resolver for the addParentToUser field.
func (r *mutationResolver) AddParentToUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	return link.AddParentToUser(ctx, &input)
}

// RemoveUserParent is the resolver for the removeUserParent field.
func (r *mutationResolver) RemoveUserParent(ctx context.Context, input model.UserInput) (*string, error) {
	return link.RemoveUserParent(ctx, &input)
}

// AddTutorToUser is the resolver for the addTutorToUser field.
func (r *mutationResolver) AddTutorToUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	return link.AddTutorToUser(ctx, &input)
}

// RemoveUserTutor is the resolver for the removeUserTutor field.
func (r *mutationResolver) RemoveUserTutor(ctx context.Context, input model.UserInput) (*string, error) {
	return link.RemoveUserTutor(ctx, &input)
}

// AddProfessorToUser is the resolver for the addProfessorToUser field.
func (r *mutationResolver) AddProfessorToUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	return link.AddProfessorToUser(ctx, &input)
}

// RemoveUserProfessor is the resolver for the removeUserProfessor field.
func (r *mutationResolver) RemoveUserProfessor(ctx context.Context, input model.UserInput) (*string, error) {
	return link.RemoveUserProfessor(ctx, &input)
}

// AddStudentToLink is the resolver for the addStudentToLink field.
func (r *mutationResolver) AddStudentToLink(ctx context.Context, input model.UserInput) (*model.User, error) {
	return link.AddStudentToLink(ctx, &input)
}

// RemoveStudent is the resolver for the removeStudent field.
func (r *mutationResolver) RemoveStudent(ctx context.Context, input model.UserInput) (*string, error) {
	return link.RemoveStudent(ctx, &input)
}

// LoginWithQR is the resolver for the loginWithQr field.
func (r *mutationResolver) LoginWithQR(ctx context.Context, xID string) (*string, error) {
	panic(fmt.Errorf("not implemented: LoginWithQR - loginWithQr"))
}

// SingleUpload is the resolver for the singleUpload field.
func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.Media, error) {
	return media.SingleUpload(ctx, file)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
