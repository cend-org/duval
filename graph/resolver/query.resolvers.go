package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"time"

	"github.com/cend-org/duval/graph/generated"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/authentication"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/pkg/address"
	"github.com/cend-org/duval/internal/pkg/code"
	"github.com/cend-org/duval/internal/pkg/contract"
	"github.com/cend-org/duval/internal/pkg/education"
	"github.com/cend-org/duval/internal/pkg/mark"
	"github.com/cend-org/duval/internal/pkg/media/cover"
	"github.com/cend-org/duval/internal/pkg/media/cv"
	"github.com/cend-org/duval/internal/pkg/media/profile"
	"github.com/cend-org/duval/internal/pkg/media/video"
	"github.com/cend-org/duval/internal/pkg/phone"
	"github.com/cend-org/duval/internal/pkg/planning"
	"github.com/cend-org/duval/internal/pkg/post"
	"github.com/cend-org/duval/internal/pkg/translator"
	"github.com/cend-org/duval/internal/pkg/user"
	"github.com/cend-org/duval/internal/pkg/user/link"
)

// Passwords is the resolver for the Passwords field.
func (r *queryResolver) Passwords(ctx context.Context) ([]model.Password, error) {
	panic(fmt.Errorf("not implemented: Passwords - Passwords"))
}

// GetAsset is the resolver for the GetAsset field.
func (r *queryResolver) GetAsset(ctx context.Context, id int) (*model.Asset, error) {
	var (
		as  model.Asset
		err error
	)

	err = database.Get(&as, `SELECT * FROM asset WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return &as, err
}

// Assets is the resolver for the Assets field.
func (r *queryResolver) Assets(ctx context.Context) ([]model.Asset, error) {
	var (
		assets []model.Asset
		err    error
	)

	err = database.Select(&assets, `SELECT * FROM asset WHERE id > 0`)
	if err != nil {
		return nil, err
	}

	return assets, err
}

// UserAuthorizationLink is the resolver for the userAuthorizationLink field.
func (r *queryResolver) UserAuthorizationLink(ctx context.Context, id int) (*model.UserAuthorizationLink, error) {
	return user.GetUserAuthorizationLink(ctx, id)
}

// UserAuthorizationLinks is the resolver for the userAuthorizationLinks field.
func (r *queryResolver) UserAuthorizationLinks(ctx context.Context) ([]model.UserAuthorizationLink, error) {
	return user.GetUserAuthorizationLinks(ctx)
}

// GetCode is the resolver for the getCode field.
func (r *queryResolver) GetCode(ctx context.Context) (*model.Code, error) {
	return code.GetCode(ctx)
}

// VerifyUserEmailValidationCode is the resolver for the verifyUserEmailValidationCode field.
func (r *queryResolver) VerifyUserEmailValidationCode(ctx context.Context, verifCode int) (int, error) {
	return code.VerifyUserEmailValidationCode(ctx, verifCode)
}

// SendUserEmailValidationCode is the resolver for the sendUserEmailValidationCode field.
func (r *queryResolver) SendUserEmailValidationCode(ctx context.Context) (*model.User, error) {
	return code.SendUserEmailValidationCode(ctx)
}

// GetPasswordHistory is the resolver for the getPasswordHistory field.
func (r *queryResolver) GetPasswordHistory(ctx context.Context) ([]model.Password, error) {
	return user.GetPasswordHistory(ctx)
}

// ActivateUser is the resolver for the activateUser field.
func (r *queryResolver) ActivateUser(ctx context.Context) (*model.User, error) {
	return user.ActivateUser(ctx)
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context) ([]model.User, error) {
	panic(fmt.Errorf("not implemented: Register - Register"))
}

// MyProfile is the resolver for the MyProfile field.
func (r *queryResolver) MyProfile(ctx context.Context) (*model.User, error) {
	return user.MyProfile(ctx)
}

// GetMessages is the resolver for the getMessages field.
func (r *queryResolver) GetMessages(ctx context.Context) ([]model.Message, error) {
	return translator.GetMessages(ctx)
}

// GetMessagesInLanguage is the resolver for the getMessagesInLanguage field.
func (r *queryResolver) GetMessagesInLanguage(ctx context.Context, language int) ([]model.Message, error) {
	return translator.GetMessagesInLanguage(ctx, language)
}

// GetMessage is the resolver for the getMessage field.
func (r *queryResolver) GetMessage(ctx context.Context, language int, resourceNumber int) (*model.Message, error) {
	return translator.GetMessage(ctx, language, resourceNumber)
}

// GetMenuList is the resolver for the getMenuList field.
func (r *queryResolver) GetMenuList(ctx context.Context) ([]model.Message, error) {
	return translator.GetMenuList(ctx)
}

// GetMenuItems is the resolver for the getMenuItems field.
func (r *queryResolver) GetMenuItems(ctx context.Context, language int, menuNumber int) ([]model.Message, error) {
	return translator.GetMenuItems(ctx, language, menuNumber)
}

// GetUserAddress is the resolver for the getUserAddress field.
func (r *queryResolver) GetUserAddress(ctx context.Context) (*model.Address, error) {
	return address.GetUserAddress(ctx)
}

// GetUserPhoneNumber is the resolver for the getUserPhoneNumber field.
func (r *queryResolver) GetUserPhoneNumber(ctx context.Context) (*model.PhoneNumber, error) {
	return phone.GetUserPhoneNumber(ctx)
}

// GetUserPlannings is the resolver for the getUserPlannings field.
func (r *queryResolver) GetUserPlannings(ctx context.Context) (*model.CalendarPlanning, error) {
	return planning.GetUserPlannings(ctx)
}

// GetPlanningActors is the resolver for the getPlanningActors field.
func (r *queryResolver) GetPlanningActors(ctx context.Context, calendarID int) ([]model.User, error) {
	return planning.GetPlanningActors(ctx, calendarID)
}

// GetUserSubjects is the resolver for the getUserSubjects field.
func (r *queryResolver) GetUserSubjects(ctx context.Context) ([]model.Subject, error) {
	return education.GetUserSubjects(ctx)
}

// GetEducation is the resolver for the getEducation field.
func (r *queryResolver) GetEducation(ctx context.Context) ([]model.Education, error) {
	return education.GetEducation(ctx)
}

// GetUserEducationLevel is the resolver for the getUserEducationLevel field.
func (r *queryResolver) GetUserEducationLevel(ctx context.Context) (*model.Education, error) {
	return education.GetUserEducationLevel(ctx)
}

// GetSchools is the resolver for the getSchools field.
func (r *queryResolver) GetSchools(ctx context.Context) ([]model.School, error) {
	return education.GetSchools(ctx)
}

// GetSubjects is the resolver for the getSubjects field.
func (r *queryResolver) GetSubjects(ctx context.Context, id int) ([]model.SchoolSubject, error) {
	return education.GetSubjects(ctx, id)
}

// GetSchool is the resolver for the getSchool field.
func (r *queryResolver) GetSchool(ctx context.Context, id int) (*model.School, error) {
	return education.GetSchool(ctx, id)
}

// GetUserAverageMark is the resolver for the getUserAverageMark field.
func (r *queryResolver) GetUserAverageMark(ctx context.Context, userID int) (*int, error) {
	return mark.GetUserAverageMark(ctx, userID)
}

// GetUserMarkComment is the resolver for the getUserMarkComment field.
func (r *queryResolver) GetUserMarkComment(ctx context.Context) ([]model.Mark, error) {
	return mark.GetUserMarkComment(ctx)
}

// GetUserParent is the resolver for the getUserParent field.
func (r *queryResolver) GetUserParent(ctx context.Context) ([]model.User, error) {
	return link.GetUserParent(ctx)
}

// GetUserTutor is the resolver for the getUserTutor field.
func (r *queryResolver) GetUserTutor(ctx context.Context) ([]model.User, error) {
	return link.GetUserTutor(ctx)
}

// GetUserProfessor is the resolver for the getUserProfessor field.
func (r *queryResolver) GetUserProfessor(ctx context.Context) ([]model.User, error) {
	return link.GetUserProfessor(ctx)
}

// GetStudent is the resolver for the getStudent field.
func (r *queryResolver) GetStudent(ctx context.Context) ([]model.User, error) {
	return link.GetStudent(ctx)
}

// GenerateQRCode is the resolver for the generateQrCode field.
func (r *queryResolver) GenerateQRCode(ctx context.Context) (*string, error) {
	return authentication.GenerateQrCode(ctx)
}

// GetProfileLetter is the resolver for the getProfileLetter field.
func (r *queryResolver) GetProfileLetter(ctx context.Context) (*model.Media, error) {
	return cover.GetProfileLetter(ctx)
}

// GetProfileLetterThumb is the resolver for the getProfileLetterThumb field.
func (r *queryResolver) GetProfileLetterThumb(ctx context.Context) (*model.MediaThumb, error) {
	return cover.GetProfileLetterThumb(ctx)
}

// GetProfileCv is the resolver for the getProfileCv field.
func (r *queryResolver) GetProfileCv(ctx context.Context) (*model.Media, error) {
	return cv.GetProfileCv(ctx)
}

// GetProfileCvThumb is the resolver for the getProfileCvThumb field.
func (r *queryResolver) GetProfileCvThumb(ctx context.Context) (*model.MediaThumb, error) {
	return cv.GetProfileCvThumb(ctx)
}

// GetProfileImage is the resolver for the getProfileImage field.
func (r *queryResolver) GetProfileImage(ctx context.Context) (*model.Media, error) {
	return profile.GetProfileImage(ctx)
}

// GetProfileImageThumb is the resolver for the getProfileImageThumb field.
func (r *queryResolver) GetProfileImageThumb(ctx context.Context) (*model.MediaThumb, error) {
	return profile.GetProfileImageThumb(ctx)
}

// GetProfileVideo is the resolver for the getProfileVideo field.
func (r *queryResolver) GetProfileVideo(ctx context.Context) (*model.Media, error) {
	return video.GetProfileVideo(ctx)
}

// GetProfileVideoThumb is the resolver for the getProfileVideoThumb field.
func (r *queryResolver) GetProfileVideoThumb(ctx context.Context) (*model.MediaThumb, error) {
	return video.GetProfileVideoThumb(ctx)
}

// GetContracts is the resolver for the getContracts field.
func (r *queryResolver) GetContracts(ctx context.Context) ([]model.Contract, error) {
	return contract.GetContracts(ctx)
}

// GetContract is the resolver for the getContract field.
func (r *queryResolver) GetContract(ctx context.Context, contractID int) (*model.Contract, error) {
	return contract.GetContract(ctx, contractID)
}

// GetContractTimesheetDetail is the resolver for the getContractTimesheetDetail field.
func (r *queryResolver) GetContractTimesheetDetail(ctx context.Context) ([]model.ContractTimesheetDetail, error) {
	return contract.GetContractTimesheetDetail(ctx)
}

// GetContractTimesheetDetailInfo is the resolver for the getContractTimesheetDetailInfo field.
func (r *queryResolver) GetContractTimesheetDetailInfo(ctx context.Context, contractTimesheetID int) (*model.ContractTimesheetDetail, error) {
	return contract.GetContractTimesheetDetailInfo(ctx, contractTimesheetID)
}

// GetTotalSalaryValue is the resolver for the getTotalSalaryValue field.
func (r *queryResolver) GetTotalSalaryValue(ctx context.Context, studentID int, startDate time.Time, endDate time.Time) (*float64, error) {
	return contract.GetTotalSalary(ctx, studentID, startDate, endDate)
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]model.Post, error) {
	return post.GetPosts(ctx)
}

// GetUserPosts is the resolver for the getUserPosts field.
func (r *queryResolver) GetUserPosts(ctx context.Context) ([]model.Post, error) {
	return post.GetUserPosts(ctx)
}

// ViewPost is the resolver for the viewPost field.
func (r *queryResolver) ViewPost(ctx context.Context, postID int) (*model.Post, error) {
	return post.ViewPost(ctx, postID)
}

// SearchPost is the resolver for the searchPost field.
func (r *queryResolver) SearchPost(ctx context.Context, keyword string) ([]model.Post, error) {
	return post.SearchPost(ctx, keyword)
}

// GetTaggedPost is the resolver for the getTaggedPost field.
func (r *queryResolver) GetTaggedPost(ctx context.Context, postID int) ([]model.PostTag, error) {
	return post.GetTaggedPost(ctx, postID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
