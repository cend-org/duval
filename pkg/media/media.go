package mediafile

import (
	"context"
	"fmt"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/authentication"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/joinverse/xid"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type MediaQuery struct{}

func Upload(ctx *gin.Context) {
	var (
		media        model.Media
		documentType int
		tok          *token.Token
		err          error
		file         *multipart.FileHeader
	)

	tok, err = authentication.GinContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "UnAuthorized",
		})
		return
	}

	if tok.UserId == state.ZERO {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "UnAuthorized",
		})
		return
	}

	file, err = ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "cannot parse the form file upload given",
		})
		return
	}

	mType, err := DetectMimeType(file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "invalid type of file",
		})
		return
	}

	documentType, err = strconv.Atoi(ctx.PostForm("documentType"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: errx.DummyError,
		})
		return
	}

	media, err = GetMedia(tok.UserId, documentType)
	if err == nil && media.Xid != state.EMPTY {
		var userMediaDetail model.UserMediaDetail
		var mediaThumb model.MediaThumb

		mediaThumb, err = GetMediaThumb(tok.UserId, documentType)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "error while trying to get data from database",
			})
			return
		}

		userMediaDetail, err = GetUserMediaDetail(tok.UserId, documentType)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "error while trying to get data from database",
			})
			return
		}

		err = RemoveMedia(media)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "error while trying to delete data from database",
			})
			return
		}

		if !utils.IsValidVideo(mType.String()) {
			err = RemoveMediaThumb(mediaThumb)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
					Message: "error while trying to delete data from database",
				})
				return
			}
		} else {
			err = RemoveVideoThumb(mediaThumb)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
					Message: "error while trying to delete data from database",
				})
				return
			}
		}

		err = RemoveUserMediaDetail(userMediaDetail)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "error while trying to delete data from database ",
			})
			return
		}

	}

	media.Extension = filepath.Ext(file.Filename)
	media.Xid = xid.New().String()
	media.FileName = media.Xid + media.Extension

	if !utils.IsValidFile(mType.String()) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "invalid type of file",
		})
		return
	}

	if utils.IsValidDocument(mType.String()) {
		if documentType != utils.CV && documentType != utils.Letter {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "invalid type of file",
			})
			return
		}

		err = utils.CreateDocumentThumb(media.Xid, media.Extension, file)
		if err != nil {
			fmt.Println("error here: ", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "failed to create thumb",
			})
			return
		}
	}

	if utils.IsValidVideo(mType.String()) {
		if documentType != utils.VideoPresentation {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "invalid type of file",
			})
			return
		}

		err = utils.CreateVideoThumb(media.Xid, file)
		if err != nil {
			fmt.Println("error here: ", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "failed to create thumb",
			})
			return
		}
	}

	if utils.IsValidImage(mType.String()) {
		if documentType != utils.UserProfileImage {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "invalid type of file",
			})
			return
		}

		err = utils.CreateThumb(media.Xid, mType.String(), file)
		if err != nil {
			fmt.Println("error here: ", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
				Message: "failed to create thumb",
			})
			return
		}
	}

	err = ctx.SaveUploadedFile(file, utils.FILE_UPLOAD_DIR+media.Xid+media.Extension)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "failed to save file into server",
		})
		return
	}

	media.Id, err = database.InsertOne(media)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "error while trying to insert data into database",
		})
		return
	}

	err = SetUserMediaDetail(documentType, tok.UserId, media.Xid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "error while trying to insert data into database",
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, media)
}

func (r *MediaQuery) ClearAllMedia(ctx context.Context) (*bool, error) {
	status := true
	paths := []string{
		utils.FILE_UPLOAD_DIR + "*.*",
		utils.FILE_UPLOAD_DIR + utils.THUMB_FILE_UPLOAD_DIR + "*.*",
		utils.FILE_UPLOAD_DIR + utils.QR_CODE_UPLOAD_DIR + "*.*",
	}

	for _, pattern := range paths {
		matchedPaths, err := filepath.Glob(pattern)
		if err != nil {
			return &status, err
		}
		for _, path := range matchedPaths {
			err = ClearMediaFile(path)
			if err != nil {
				status = false
				return &status, err
			}
		}
	}

	return &status, nil
}

/*

	Utils

*/

func DetectMimeType(file *multipart.FileHeader) (mType *mimetype.MIME, err error) {
	readFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	mType, err = mimetype.DetectReader(readFile)
	if err != nil {
		return nil, err
	}
	return mType, nil
}

func SetUserMediaDetail(documentType int, userId int, xId string) (err error) {
	var (
		userMediaDetail model.UserMediaDetail
	)

	userMediaDetail.OwnerId = userId
	userMediaDetail.DocumentType = documentType
	userMediaDetail.DocumentXid = xId

	_, err = database.InsertOne(userMediaDetail)
	if err != nil {
		return err
	}
	return
}

func GetUserMediaDetail(userId, documentType int) (userMediaDetail model.UserMediaDetail, err error) {
	err = database.Get(&userMediaDetail, `SELECT * FROM user_media_detail WHERE owner_id = ?   AND document_type = ?`, userId, documentType)
	if err != nil {
		return userMediaDetail, err
	}
	return userMediaDetail, err
}

func GetMedia(userId, documentType int) (media model.Media, err error) {
	err = database.Get(&media,
		`SELECT media.*
			FROM media
					 JOIN user_media_detail umd ON media.xid = umd.document_xid
					 JOIN user ON user.id = umd.owner_id
			WHERE umd.owner_id = ? AND umd.document_type = ?`, userId, documentType)
	if err != nil {
		return media, err
	}
	return media, err
}

func GetMediaThumb(userId, documentType int) (media model.MediaThumb, err error) {
	err = database.Get(&media,
		`SELECT mt.*
			FROM media_thumb mt
					 JOIN media ON mt.media_xid = media.xid
					 JOIN user_media_detail umd ON umd.document_xid = media.xid
			WHERE umd.owner_id = ?
			  AND umd.document_type = ?`, userId, documentType)
	if err != nil {
		return media, err
	}
	return media, nil
}

func RemoveUserMediaDetail(userMediaDetail model.UserMediaDetail) (err error) {
	err = database.Delete(userMediaDetail)
	if err != nil {
		return err
	}
	return
}

func RemoveMedia(media model.Media) (err error) {
	err = database.Delete(media)
	if err != nil {
		return err
	}

	filePath := utils.FILE_UPLOAD_DIR + media.Xid + media.Extension

	err = ClearMediaFile(filePath)
	if err != nil {
		return err
	}

	return
}

func RemoveMediaThumb(mediaThumb model.MediaThumb) (err error) {
	err = database.Delete(mediaThumb)
	if err != nil {
		return err
	}

	filePath := utils.FILE_UPLOAD_DIR + utils.THUMB_FILE_UPLOAD_DIR + mediaThumb.Xid + mediaThumb.Extension

	err = ClearMediaFile(filePath)
	if err != nil {
		return err
	}
	return nil
}

func RemoveVideoThumb(mediaThumb model.MediaThumb) (err error) {
	err = database.Delete(mediaThumb)
	if err != nil {
		return err
	}
	return
}

func ClearMediaFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
