package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"duval/internal/authentication"
	"duval/internal/graph/model"
	"duval/internal/pkg/address"
	"duval/internal/pkg/education"
	"duval/internal/pkg/mark"
	"duval/internal/pkg/phone"
	"duval/internal/pkg/planning"
	"duval/internal/pkg/translator"
	"duval/internal/pkg/user"
	"duval/internal/pkg/user/link"
	"fmt"
)

// ID is the resolver for the id field.
func (r *addressResolver) ID(ctx context.Context, obj *model.Address) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

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
func (r *calendarPlanningResolver) ID(ctx context.Context, obj *model.CalendarPlanning) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// AuthorizationID is the resolver for the authorizationId field.
func (r *calendarPlanningResolver) AuthorizationID(ctx context.Context, obj *model.CalendarPlanning) (int, error) {
	panic(fmt.Errorf("not implemented: AuthorizationID - authorizationId"))
}

// ID is the resolver for the id field.
func (r *calendarPlanningActorResolver) ID(ctx context.Context, obj *model.CalendarPlanningActor) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// AuthorizationID is the resolver for the authorizationId field.
func (r *calendarPlanningActorResolver) AuthorizationID(ctx context.Context, obj *model.CalendarPlanningActor) (int, error) {
	panic(fmt.Errorf("not implemented: AuthorizationID - authorizationId"))
}

// CalendarPlanningID is the resolver for the calendarPlanningId field.
func (r *calendarPlanningActorResolver) CalendarPlanningID(ctx context.Context, obj *model.CalendarPlanningActor) (*int, error) {
	panic(fmt.Errorf("not implemented: CalendarPlanningID - calendarPlanningId"))
}

// ID is the resolver for the id field.
func (r *codeResolver) ID(ctx context.Context, obj *model.Code) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *codeResolver) UserID(ctx context.Context, obj *model.Code) (int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// ID is the resolver for the id field.
func (r *educationResolver) ID(ctx context.Context, obj *model.Education) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *messageResolver) ID(ctx context.Context, obj *model.Message) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUserInput, typeArg int) (string, error) {
	return user.Register(&input, &typeArg)
}

// UpdMyProfile is the resolver for the UpdMyProfile field.
func (r *mutationResolver) UpdMyProfile(ctx context.Context, input map[string]interface{}) (*model.User, error) {
	return user.UpdMyProfile(&ctx, input)
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

// NewMessage is the resolver for the newMessage field.
func (r *mutationResolver) NewMessage(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	return translator.NewMessage(&ctx, &input)
}

// UpdMessage is the resolver for the updMessage field.
func (r *mutationResolver) UpdMessage(ctx context.Context, input model.MessageUpdateInput) (*model.Message, error) {
	return translator.UpdMessage(&ctx, &input)
}

// DelMessage is the resolver for the delMessage field.
func (r *mutationResolver) DelMessage(ctx context.Context, language int, messageNumber int) (*string, error) {
	return translator.DelMessage(&ctx, language, messageNumber)
}

// NewMenu is the resolver for the newMenu field.
func (r *mutationResolver) NewMenu(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	return translator.NewMenu(&ctx, &input)
}

// DelMenu is the resolver for the delMenu field.
func (r *mutationResolver) DelMenu(ctx context.Context, menuNumber int) (*string, error) {
	return translator.DelMenu(&ctx, menuNumber)
}

// NewMenuItem is the resolver for the newMenuItem field.
func (r *mutationResolver) NewMenuItem(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	return translator.NewMenuItem(&ctx, &input)
}

// DelMenuItem is the resolver for the delMenuItem field.
func (r *mutationResolver) DelMenuItem(ctx context.Context, input model.MessageUpdateInput) (*string, error) {
	return translator.DelMenuItem(&ctx, &input)
}

// NewAddress is the resolver for the newAddress field.
func (r *mutationResolver) NewAddress(ctx context.Context, input model.NewAddress) (*model.Address, error) {
	return address.NewAddress(&ctx, &input)
}

// UpdateUserAddress is the resolver for the UpdateUserAddress field.
func (r *mutationResolver) UpdateUserAddress(ctx context.Context, input model.NewAddress) (*model.Address, error) {
	return address.UpdateUserAddress(&ctx, &input)
}

// NewPhoneNumber is the resolver for the newPhoneNumber field.
func (r *mutationResolver) NewPhoneNumber(ctx context.Context, input model.NewPhoneNumber) (*model.PhoneNumber, error) {
	return phone.NewPhoneNumber(&ctx, &input)
}

// UpdateUserPhoneNumber is the resolver for the updateUserPhoneNumber field.
func (r *mutationResolver) UpdateUserPhoneNumber(ctx context.Context, input model.NewPhoneNumber) (*model.PhoneNumber, error) {
	return phone.UpdateUserPhoneNumber(&ctx, &input)
}

// CreateUserPlannings is the resolver for the createUserPlannings field.
func (r *mutationResolver) CreateUserPlannings(ctx context.Context, input model.NewCalendarPlanning) (*model.CalendarPlanning, error) {
	return planning.CreateUserPlannings(&ctx, &input)
}

// AddUserIntoPlanning is the resolver for the addUserIntoPlanning field.
func (r *mutationResolver) AddUserIntoPlanning(ctx context.Context, calendarID int, selectedUserID int) (*model.CalendarPlanningActor, error) {
	return planning.AddUserIntoPlanning(&ctx, calendarID, selectedUserID)
}

// SetUserEducationLevel is the resolver for the setUserEducationLevel field.
func (r *mutationResolver) SetUserEducationLevel(ctx context.Context, input model.SubjectInput) (*model.Education, error) {
	return education.SetUserEducationLevel(&ctx, &input)
}

// UpdateUserEducationLevel is the resolver for the updateUserEducationLevel field.
func (r *mutationResolver) UpdateUserEducationLevel(ctx context.Context, input model.SubjectInput) (*model.Education, error) {
	return education.UpdateUserEducationLevel(&ctx, &input)
}

// RateUser is the resolver for the rateUser field.
func (r *mutationResolver) RateUser(ctx context.Context, input model.UserMarkInput) (*model.UserMark, error) {
	return mark.RateUser(&ctx, &input)
}

// AddParentToUser is the resolver for the addParentToUser field.
func (r *mutationResolver) AddParentToUser(ctx context.Context, input map[string]interface{}) (*model.User, error) {
	return link.AddParentToUser(&ctx, input)
}

// RemoveUserParent is the resolver for the removeUserParent field.
func (r *mutationResolver) RemoveUserParent(ctx context.Context, input map[string]interface{}) (*string, error) {
	return link.RemoveUserParent(&ctx, input)
}

// AddTutorToUser is the resolver for the addTutorToUser field.
func (r *mutationResolver) AddTutorToUser(ctx context.Context, input map[string]interface{}) (*model.User, error) {
	return link.AddTutorToUser(&ctx, input)
}

// RemoveUserTutor is the resolver for the removeUserTutor field.
func (r *mutationResolver) RemoveUserTutor(ctx context.Context, input map[string]interface{}) (*string, error) {
	return link.RemoveUserTutor(&ctx, input)
}

// AddProfessorToUser is the resolver for the addProfessorToUser field.
func (r *mutationResolver) AddProfessorToUser(ctx context.Context, input map[string]interface{}) (*model.User, error) {
	return link.AddProfessorToUser(&ctx, input)
}

// RemoveUserProfessor is the resolver for the removeUserProfessor field.
func (r *mutationResolver) RemoveUserProfessor(ctx context.Context, input map[string]interface{}) (*string, error) {
	return link.RemoveUserProfessor(&ctx, input)
}

// AddStudentToLink is the resolver for the addStudentToLink field.
func (r *mutationResolver) AddStudentToLink(ctx context.Context, input map[string]interface{}) (*model.User, error) {
	return link.AddStudentToLink(&ctx, input)
}

// RemoveStudent is the resolver for the removeStudent field.
func (r *mutationResolver) RemoveStudent(ctx context.Context, input map[string]interface{}) (*string, error) {
	return link.RemoveStudent(&ctx, input)
}

// LoginWithQR is the resolver for the LoginWithQr field.
func (r *mutationResolver) LoginWithQR(ctx context.Context, xID string) (*string, error) {
	return authentication.LoginWithQr(&ctx, xID)
}

// ID is the resolver for the id field.
func (r *passwordResolver) ID(ctx context.Context, obj *model.Password) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *passwordResolver) UserID(ctx context.Context, obj *model.Password) (int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// ID is the resolver for the id field.
func (r *phoneNumberResolver) ID(ctx context.Context, obj *model.PhoneNumber) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ID is the resolver for the id field.
func (r *qrCodeRegistryResolver) ID(ctx context.Context, obj *model.QrCodeRegistry) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ResourceType is the resolver for the resourceType field.
func (r *qrCodeRegistryResolver) ResourceType(ctx context.Context, obj *model.QrCodeRegistry) (int, error) {
	panic(fmt.Errorf("not implemented: ResourceType - resourceType"))
}

// ResourceValue is the resolver for the resourceValue field.
func (r *qrCodeRegistryResolver) ResourceValue(ctx context.Context, obj *model.QrCodeRegistry) (int, error) {
	panic(fmt.Errorf("not implemented: ResourceValue - resourceValue"))
}

// UserID is the resolver for the userId field.
func (r *qrCodeRegistryResolver) UserID(ctx context.Context, obj *model.QrCodeRegistry) (int, error) {
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

// GetMessages is the resolver for the getMessages field.
func (r *queryResolver) GetMessages(ctx context.Context) ([]*model.Message, error) {
	return translator.GetMessages(&ctx)
}

// GetMessagesInLanguage is the resolver for the getMessagesInLanguage field.
func (r *queryResolver) GetMessagesInLanguage(ctx context.Context, language int) ([]*model.Message, error) {
	return translator.GetMessagesInLanguage(&ctx, language)
}

// GetMessage is the resolver for the getMessage field.
func (r *queryResolver) GetMessage(ctx context.Context, language int, resourceNumber int) (*model.Message, error) {
	return translator.GetMessage(&ctx, language, resourceNumber)
}

// GetMenuList is the resolver for the getMenuList field.
func (r *queryResolver) GetMenuList(ctx context.Context) ([]*model.Message, error) {
	return translator.GetMenuList(&ctx)
}

// GetMenuItems is the resolver for the getMenuItems field.
func (r *queryResolver) GetMenuItems(ctx context.Context, language int, menuNumber int) ([]*model.Message, error) {
	return translator.GetMenuItems(&ctx, language, menuNumber)
}

// GetUserAddress is the resolver for the getUserAddress field.
func (r *queryResolver) GetUserAddress(ctx context.Context) (*model.Address, error) {
	return address.GetUserAddress(&ctx)
}

// RemoveUserAddress is the resolver for the removeUserAddress field.
func (r *queryResolver) RemoveUserAddress(ctx context.Context) (string, error) {
	return address.RemoveUserAddress(&ctx)
}

// GetUserPhoneNumber is the resolver for the getUserPhoneNumber field.
func (r *queryResolver) GetUserPhoneNumber(ctx context.Context) (*model.PhoneNumber, error) {
	return phone.GetUserPhoneNumber(&ctx)
}

// GetUserPlannings is the resolver for the getUserPlannings field.
func (r *queryResolver) GetUserPlannings(ctx context.Context) (*model.CalendarPlanning, error) {
	return planning.GetUserPlannings(&ctx)
}

// RemoveUserPlannings is the resolver for the removeUserPlannings field.
func (r *queryResolver) RemoveUserPlannings(ctx context.Context) (*string, error) {
	return planning.RemoveUserPlannings(&ctx)
}

// GetPlanningActors is the resolver for the getPlanningActors field.
func (r *queryResolver) GetPlanningActors(ctx context.Context, calendarID int) ([]*model.User, error) {
	return planning.GetPlanningActors(&ctx, calendarID)
}

// RemoveUserFromPlanning is the resolver for the removeUserFromPlanning field.
func (r *queryResolver) RemoveUserFromPlanning(ctx context.Context, calendarPlanningID int, selectedUserID int) (*string, error) {
	return planning.RemoveUserFromPlanning(&ctx, calendarPlanningID, selectedUserID)
}

// GetUserSubjects is the resolver for the getUserSubjects field.
func (r *queryResolver) GetUserSubjects(ctx context.Context) ([]*model.Subject, error) {
	return education.GetUserSubjects(&ctx)
}

// GetSubjects is the resolver for the getSubjects field.
func (r *queryResolver) GetSubjects(ctx context.Context, eduID int) ([]*model.Subject, error) {
	return education.GetSubjects(eduID)
}

// GetEducation is the resolver for the getEducation field.
func (r *queryResolver) GetEducation(ctx context.Context) ([]*model.Education, error) {
	return education.GetEducation()
}

// GetUserEducationLevel is the resolver for the getUserEducationLevel field.
func (r *queryResolver) GetUserEducationLevel(ctx context.Context) (*model.Education, error) {
	return education.GetUserEducationLevel(&ctx)
}

// GetUserAverageMark is the resolver for the getUserAverageMark field.
func (r *queryResolver) GetUserAverageMark(ctx context.Context, userID int) (*int, error) {
	return mark.GetUserAverageMark(&ctx, userID)
}

// GetUserMarkComment is the resolver for the getUserMarkComment field.
func (r *queryResolver) GetUserMarkComment(ctx context.Context) ([]*model.UserMark, error) {
	return mark.GetUserMarkComment(&ctx)
}

// GetUserParent is the resolver for the getUserParent field.
func (r *queryResolver) GetUserParent(ctx context.Context) ([]*model.User, error) {
	return link.GetUserParent(&ctx)
}

// GetUserTutor is the resolver for the getUserTutor field.
func (r *queryResolver) GetUserTutor(ctx context.Context) ([]*model.User, error) {
	return link.GetUserTutor(&ctx)
}

// GetUserProfessor is the resolver for the getUserProfessor field.
func (r *queryResolver) GetUserProfessor(ctx context.Context) ([]*model.User, error) {
	return link.GetUserProfessor(&ctx)
}

// GetStudent is the resolver for the getStudent field.
func (r *queryResolver) GetStudent(ctx context.Context) ([]*model.User, error) {
	return link.GetStudent(&ctx)
}

// GenerateQRCode is the resolver for the generateQrCode field.
func (r *queryResolver) GenerateQRCode(ctx context.Context) (*string, error) {
	return authentication.GenerateQrCode(&ctx)
}

// ID is the resolver for the id field.
func (r *subjectResolver) ID(ctx context.Context, obj *model.Subject) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// EducationLevelID is the resolver for the educationLevelId field.
func (r *subjectResolver) EducationLevelID(ctx context.Context, obj *model.Subject) (int, error) {
	panic(fmt.Errorf("not implemented: EducationLevelID - educationLevelId"))
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
func (r *userAddressResolver) ID(ctx context.Context, obj *model.UserAddress) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *userAddressResolver) UserID(ctx context.Context, obj *model.UserAddress) (*int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// AddressID is the resolver for the addressId field.
func (r *userAddressResolver) AddressID(ctx context.Context, obj *model.UserAddress) (*int, error) {
	panic(fmt.Errorf("not implemented: AddressID - addressId"))
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

// ID is the resolver for the id field.
func (r *userEducationLevelSubjectResolver) ID(ctx context.Context, obj *model.UserEducationLevelSubject) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *userEducationLevelSubjectResolver) UserID(ctx context.Context, obj *model.UserEducationLevelSubject) (int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// SubjectID is the resolver for the SubjectId field.
func (r *userEducationLevelSubjectResolver) SubjectID(ctx context.Context, obj *model.UserEducationLevelSubject) (int, error) {
	panic(fmt.Errorf("not implemented: SubjectID - SubjectId"))
}

// ID is the resolver for the id field.
func (r *userMarkResolver) ID(ctx context.Context, obj *model.UserMark) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *userMarkResolver) UserID(ctx context.Context, obj *model.UserMark) (int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// AuthorID is the resolver for the authorId field.
func (r *userMarkResolver) AuthorID(ctx context.Context, obj *model.UserMark) (int, error) {
	panic(fmt.Errorf("not implemented: AuthorID - authorId"))
}

// AuthorMark is the resolver for the authorMark field.
func (r *userMarkResolver) AuthorMark(ctx context.Context, obj *model.UserMark) (string, error) {
	panic(fmt.Errorf("not implemented: AuthorMark - authorMark"))
}

// ID is the resolver for the id field.
func (r *userPhoneNumberResolver) ID(ctx context.Context, obj *model.UserPhoneNumber) (int, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// UserID is the resolver for the userId field.
func (r *userPhoneNumberResolver) UserID(ctx context.Context, obj *model.UserPhoneNumber) (int, error) {
	panic(fmt.Errorf("not implemented: UserID - userId"))
}

// PhoneNumberID is the resolver for the phoneNumberId field.
func (r *userPhoneNumberResolver) PhoneNumberID(ctx context.Context, obj *model.UserPhoneNumber) (int, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberID - phoneNumberId"))
}

// ID is the resolver for the Id field.
func (r *messageUpdateInputResolver) ID(ctx context.Context, obj *model.MessageUpdateInput, data int) error {
	panic(fmt.Errorf("not implemented: ID - Id"))
}

// Age is the resolver for the Age field.
func (r *newUserInputResolver) Age(ctx context.Context, obj *model.NewUserInput, data *int) error {
	panic(fmt.Errorf("not implemented: Age - Age"))
}

// ID is the resolver for the Id field.
func (r *subjectInputResolver) ID(ctx context.Context, obj *model.SubjectInput, data *int) error {
	panic(fmt.Errorf("not implemented: ID - Id"))
}

// EducationLevelID is the resolver for the EducationLevelId field.
func (r *subjectInputResolver) EducationLevelID(ctx context.Context, obj *model.SubjectInput, data int) error {
	panic(fmt.Errorf("not implemented: EducationLevelID - EducationLevelId"))
}

// UserID is the resolver for the UserId field.
func (r *userMarkInputResolver) UserID(ctx context.Context, obj *model.UserMarkInput, data int) error {
	panic(fmt.Errorf("not implemented: UserID - UserId"))
}

// AuthorID is the resolver for the AuthorId field.
func (r *userMarkInputResolver) AuthorID(ctx context.Context, obj *model.UserMarkInput, data int) error {
	panic(fmt.Errorf("not implemented: AuthorID - AuthorId"))
}

// AuthorMark is the resolver for the AuthorMark field.
func (r *userMarkInputResolver) AuthorMark(ctx context.Context, obj *model.UserMarkInput, data string) error {
	panic(fmt.Errorf("not implemented: AuthorMark - AuthorMark"))
}

// Address returns AddressResolver implementation.
func (r *Resolver) Address() AddressResolver { return &addressResolver{r} }

// Authorization returns AuthorizationResolver implementation.
func (r *Resolver) Authorization() AuthorizationResolver { return &authorizationResolver{r} }

// CalendarPlanning returns CalendarPlanningResolver implementation.
func (r *Resolver) CalendarPlanning() CalendarPlanningResolver { return &calendarPlanningResolver{r} }

// CalendarPlanningActor returns CalendarPlanningActorResolver implementation.
func (r *Resolver) CalendarPlanningActor() CalendarPlanningActorResolver {
	return &calendarPlanningActorResolver{r}
}

// Code returns CodeResolver implementation.
func (r *Resolver) Code() CodeResolver { return &codeResolver{r} }

// Education returns EducationResolver implementation.
func (r *Resolver) Education() EducationResolver { return &educationResolver{r} }

// Message returns MessageResolver implementation.
func (r *Resolver) Message() MessageResolver { return &messageResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Password returns PasswordResolver implementation.
func (r *Resolver) Password() PasswordResolver { return &passwordResolver{r} }

// PhoneNumber returns PhoneNumberResolver implementation.
func (r *Resolver) PhoneNumber() PhoneNumberResolver { return &phoneNumberResolver{r} }

// QrCodeRegistry returns QrCodeRegistryResolver implementation.
func (r *Resolver) QrCodeRegistry() QrCodeRegistryResolver { return &qrCodeRegistryResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subject returns SubjectResolver implementation.
func (r *Resolver) Subject() SubjectResolver { return &subjectResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// UserAddress returns UserAddressResolver implementation.
func (r *Resolver) UserAddress() UserAddressResolver { return &userAddressResolver{r} }

// UserAuthorizationLink returns UserAuthorizationLinkResolver implementation.
func (r *Resolver) UserAuthorizationLink() UserAuthorizationLinkResolver {
	return &userAuthorizationLinkResolver{r}
}

// UserAuthorizationLinkActor returns UserAuthorizationLinkActorResolver implementation.
func (r *Resolver) UserAuthorizationLinkActor() UserAuthorizationLinkActorResolver {
	return &userAuthorizationLinkActorResolver{r}
}

// UserEducationLevelSubject returns UserEducationLevelSubjectResolver implementation.
func (r *Resolver) UserEducationLevelSubject() UserEducationLevelSubjectResolver {
	return &userEducationLevelSubjectResolver{r}
}

// UserMark returns UserMarkResolver implementation.
func (r *Resolver) UserMark() UserMarkResolver { return &userMarkResolver{r} }

// UserPhoneNumber returns UserPhoneNumberResolver implementation.
func (r *Resolver) UserPhoneNumber() UserPhoneNumberResolver { return &userPhoneNumberResolver{r} }

// MessageUpdateInput returns MessageUpdateInputResolver implementation.
func (r *Resolver) MessageUpdateInput() MessageUpdateInputResolver {
	return &messageUpdateInputResolver{r}
}

// NewUserInput returns NewUserInputResolver implementation.
func (r *Resolver) NewUserInput() NewUserInputResolver { return &newUserInputResolver{r} }

// SubjectInput returns SubjectInputResolver implementation.
func (r *Resolver) SubjectInput() SubjectInputResolver { return &subjectInputResolver{r} }

// UserMarkInput returns UserMarkInputResolver implementation.
func (r *Resolver) UserMarkInput() UserMarkInputResolver { return &userMarkInputResolver{r} }

type addressResolver struct{ *Resolver }
type authorizationResolver struct{ *Resolver }
type calendarPlanningResolver struct{ *Resolver }
type calendarPlanningActorResolver struct{ *Resolver }
type codeResolver struct{ *Resolver }
type educationResolver struct{ *Resolver }
type messageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type passwordResolver struct{ *Resolver }
type phoneNumberResolver struct{ *Resolver }
type qrCodeRegistryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subjectResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userAddressResolver struct{ *Resolver }
type userAuthorizationLinkResolver struct{ *Resolver }
type userAuthorizationLinkActorResolver struct{ *Resolver }
type userEducationLevelSubjectResolver struct{ *Resolver }
type userMarkResolver struct{ *Resolver }
type userPhoneNumberResolver struct{ *Resolver }
type messageUpdateInputResolver struct{ *Resolver }
type newUserInputResolver struct{ *Resolver }
type subjectInputResolver struct{ *Resolver }
type userMarkInputResolver struct{ *Resolver }
