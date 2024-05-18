package link

import (
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/cend-org/duval/pkg/user/authorization"
	"github.com/xorcare/pointer"
)

const (
	// AuthorizationLevel

	StudentAuthorizationLevel   = 0
	ParentAuthorizationLevel    = 1
	TutorAuthorizationLevel     = 2
	ProfessorAuthorizationLevel = 3

	//Link_type

	StudentParent    = 0
	StudentTutor     = 1
	StudentProfessor = 2
)

func AddStudent(parentId int, email string) (studentId int, err error) {
	var (
		userAuthorizationLinkId int
		student                 model.User
	)

	student, err = GetUserWithEmail(email)
	if err != nil {
		return studentId, errx.DbGetError
	}

	studentId = student.Id

	auth, err := authorization.GetUserAuthorization(student.Id, StudentAuthorizationLevel)
	if err != nil {
		return studentId, errx.DbGetError
	}

	userAuthorizationLinkId, err = GetUserLink(StudentParent, auth.Id)
	if userAuthorizationLinkId != state.ZERO {
		currentParentAuth, err := authorization.GetUserAuthorization(parentId, ParentAuthorizationLevel)
		if err != nil {
			return studentId, errx.DbGetError
		}

		parents, err := GetLink(currentParentAuth.Id, ParentAuthorizationLevel, StudentParent)
		if len(parents) > 0 {
			return studentId, errx.DuplicateUserError
		}
	}

	if userAuthorizationLinkId == state.ZERO {

		userAuthorizationLinkId, err = SetUserAuthorizationLink(StudentParent, student.Id, StudentAuthorizationLevel)
		if err != nil {
			return studentId, errx.DbInsertError
		}
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, parentId, ParentAuthorizationLevel)
	if err != nil {
		return studentId, errx.DbInsertError
	}

	return studentId, nil
}

func UpdateStudent(studentId int, profile model.UserInput) (err error) {
	var (
		user model.User
	)

	user, err = GetUserWithId(studentId)
	if err != nil {
		return errx.DbGetError
	}

	user = model.MapUserInputToUser(profile, user)

	err = database.Update(user)
	if err != nil {
		return errx.DbUpdateError
	}

	return nil
}

func NewStudentAcademicCoursesPreferenceByParent(studentId int, new []*model.UserAcademicCoursePreferenceInput) (ret *bool, err error) {
	err = database.Exec(`DELETE FROM user_academic_course WHERE user_id = ?`, studentId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(new); i++ {
		courseInput := new[i]
		if courseInput != nil {
			course := model.MapUserAcademicCoursePreferenceInputToUserAcademicCoursePreference(*courseInput, model.UserAcademicCoursePreference{})

			_, err = database.Insert(course)
			if err != nil {
				return nil, err
			}
		}
	}

	return pointer.Bool(true), err
}

/*
 UTILS
*/

func GetUserWithEmail(email string) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserWithId(userId int) (user model.User, err error) {
	err = database.Get(&user, `SELECT * FROM user WHERE id = ?`, userId)
	if err != nil {
		return user, err
	}
	return user, nil
}

func SetUserAuthorizationLink(linkType int, userId int, userLevel int) (userAuthorizationLinkId int, err error) {
	var (
		userAuthorizationLink model.UserAuthorizationLink
	)

	userAuthorizationLink.LinkType = linkType

	userAuthorizationLinkId, err = database.InsertOne(userAuthorizationLink)
	if err != nil {
		return userAuthorizationLinkId, err
	}

	err = SetUserAuthorizationLinkActor(userAuthorizationLinkId, userId, userLevel)
	if err != nil {
		return userAuthorizationLinkId, err
	}

	return userAuthorizationLinkId, nil
}

func SetUserAuthorizationLinkActor(linkId int, userId int, level int) (err error) {
	var userAuthorizationLinkActor model.UserAuthorizationLinkActor

	auth, err := authorization.GetUserAuthorization(userId, level)
	if err != nil {
		return err
	}

	userAuthorizationLinkActor.AuthorizationId = auth.Id
	userAuthorizationLinkActor.UserAuthorizationLinkId = linkId

	_, err = database.InsertOne(userAuthorizationLinkActor)
	if err != nil {
		return err
	}

	return nil
}

func GetLink(authId int, authorizationLevel int, linkType int) (link []model.User, err error) {
	err = database.Select(&link,
		`SELECT u.*
			FROM user  u
					 JOIN authorization  a ON u.Id = a.user_id
					 JOIN user_authorization_link_actor  ua_la ON a.Id = ua_la.authorization_id
					 JOIN user_authorization_link  ua ON ua_la.user_authorization_link_id = ua.id
					 JOIN user_authorization_link  ua2 ON ua.id = ua2.id
					JOIN user_authorization_link_actor  ua_la2 ON ua2.id = ua_la2.user_authorization_link_id
			WHERE ua_la2.authorization_id = ?
			  AND ua2.link_type = ?
			  AND a.level = ?`, authId, linkType, authorizationLevel)
	if err != nil {
		return link, err
	}

	return link, nil
}

func GetUserLink(linkType int, authorizationId int) (linkId int, err error) {
	var userLink model.UserAuthorizationLink

	err = database.Get(&userLink,
		`SELECT ual.* FROM user_authorization_link ual
                                  JOIN user_authorization_link_actor ua_la ON ual.Id = ua_la.user_authorization_link_id
                                  WHERE ual.link_type = ? AND ua_la.authorization_id = ?`, linkType, authorizationId)
	if err != nil {
		return state.ZERO, err
	}

	return userLink.Id, nil
}

func IsStudentParentLinked(parentId, userId int) bool {
	return true
}
