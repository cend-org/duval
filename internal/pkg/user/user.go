package user

import (
	"database/sql"
	"duval/internal/authentication"
	"duval/internal/utils"
	"duval/internal/utils/state"
	"duval/pkg/database"
	"errors"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id         uint       `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	FirstName  string     `json:"first_name"`
	MiddleName string     `json:"middle_name"`
	LastName   string     `json:"last_name"`
	NickName   string     `json:"nick_name"`
	Email      string     `json:"email"`
	Matricule  string     `json:"matricule"`
	Age        uint       `json:"age"`
	BirthDate  time.Time  `json:"birth_date"`
	Sex        int        `json:"sex"`
	Lang       int        `json:"language"`
	Status     int        `json:"status"`
}

type Password struct {
	Id          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	UserId      uint       `json:"user_id"`
	Psw         string     `json:"psw"`
	ContentHash string     `json:"hash"`
}

type Authorization struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UserId    uint       `json:"user_id"`
	Level     uint       `json:"level"`
}

/*

	ROUTES

*/

/*
NewUser creates a new record of user in the system
*/
func NewUser(ctx *gin.Context) {
	var (
		user User
		err  error
	)

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Bad format of user",
		})
		return
	}

	if !utils.IsValidEmail(user.Email) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "error parsing email address",
		})
		return
	}

	_, err = GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if user.Id > 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse{
			Message: "user already exists, abort !",
		})
		return
	}
	user.Matricule , _ = utils.GenerateMatricule()
	user.Id, err = database.InsertOne(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "failed to insert user into database",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

/*
UpdateUser updates the designed user by the id field.
user.id should be provided and user.email should not be empty.
*/
func UpdateUser(ctx *gin.Context) {
	var (
		user User
		err  error
	)

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Invalid mail format",
		})
		return
	}

	if user.Id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "user id is required for the operation",
		})
		return
	}

	existing, err := GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if existing.Id > 0 && existing.Id != user.Id {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "an other user with the same email already exists !",
		})
		return
	}
	err = database.Update(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

/*
GetUser returns the data of the user designed by the id provided in the url params
*/
func GetUser(ctx *gin.Context) {
	var (
		user User
		err  error
		id   int
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = database.Client.Get(&user, `SELECT * FROM user WHERE id = ?`, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

/*
MyProfile get all the  user's data connected
*/
func MyProfile(ctx *gin.Context) {
	var (
		tok  *authentication.Token
		err  error
		user User
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "authentiaction failed",
		})
		return
	}

	user, err = GetUserWithId(tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

/*
Login takes auth.email and auth.password as parameters.
*/
func Login(ctx *gin.Context) {
	var (
		err      error
		usr      User
		tok      authentication.Token
		auth     authentication.Auth
		password Password
	)

	err = ctx.ShouldBindJSON(&auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	// GET USER DATA
	usr, err = GetUserByEmail(auth.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	err = database.Get(&password, `SELECT * FROM password WHERE user_id = ? ORDER BY created_at desc  LIMIT 1`, usr.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if !utils.CheckPasswordHash(auth.Password, password.Psw) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "password error",
		})
		return
	}

	tok.UserId = usr.Id

	tokStr, err := authentication.NewAccessToken(tok)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokStr,
	})
	return
}

func NewPassword(ctx *gin.Context) {
	var password Password
	var err error

	err = ctx.ShouldBindJSON(&password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	if password.UserId == state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "password should be bound to user",
		})
		return
	}

	if strings.TrimSpace(password.Psw) == state.EMPTY {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "new password value cannot be empty",
		})
		return
	}

	if utils.PasswordHasValidLength(password.Psw) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "maximum value of password is 18",
		})
		return
	}

	password.ContentHash, err = utils.CreateContentHash(password.Psw)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "password should be bound to user",
		})
		return
	}

	password.Psw, err = utils.HashPassword(password.Psw)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "password should be bound to user",
		})
		return
	}

	_, err = database.Insert(password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
	return
}

func GetUserPasswordHistory(ctx *gin.Context) {
	var (
		passwords []Password
		err       error
		tok       *authentication.Token
	)

	tok, err = authentication.GetTokenDataFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	err = database.Select(&passwords, `SELECT * FROM password WHERE user_id = ? ORDER BY  created_at desc `, tok.UserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, passwords)
	return
}

/*

	UTILITIES

*/

func GetUserByEmail(email string) (user User, err error) {
	err = database.Client.Get(&user, `SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return user, err
	}

	return user, err
}

func GetUserWithId(id uint) (user User, err error) {
	err = database.Client.Get(&user, `SELECT * FROM USER WHERE id = ?`, id)
	if err != nil {
		return user, err
	}

	return user, err
}

/*

	User Level

*/

func SetUserAuthorization(ctx *gin.Context) {
	var auth Authorization
	var err error

	err = ctx.ShouldBindJSON(&auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "unknown format",
		})
		return
	}

	_, err = database.Insert(auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to insert authorization into database ",
		})
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
	return
}

func GetUserAuthorization(ctx *gin.Context) {

	var (
		auth Authorization
		err  error
		id   int
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = database.Client.Get(&auth, `SELECT * FROM authorization WHERE user_id = ?`, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to fetch authorization  user_id unknown",
		})
		return
	}

	ctx.JSON(http.StatusOK, auth)
	return
}

func RemoveUserAuthorization(ctx *gin.Context) {
	var (
		id   int
		auth Authorization
		err  error
	)
	currentTime := time.Now()

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err,
		})
	}

	err = database.Client.Get(&auth, `SELECT * FROM authorization WHERE user_id = ?`, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to fetch authorization  user_id unknown",
		})
		return
	}

	if auth.DeletedAt != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Authorization already deleted",
		})
	}

	new_auth, err := database.Client.Exec(`UPDATE authorization SET deleted_at = ? WHERE user_id = ?`, currentTime, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Failed to delete authorization",
		})
		return
	}

	ctx.JSON(http.StatusOK, new_auth)
	return
}
