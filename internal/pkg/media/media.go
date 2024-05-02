package media

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/token"
	"github.com/cend-org/duval/internal/utils"
	"github.com/cend-org/duval/internal/utils/errx"
	"github.com/cend-org/duval/internal/utils/state"
	"github.com/gabriel-vasile/mimetype"
	"github.com/joinverse/xid"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	CV                = 0
	CoverLetter       = 1
	PresentationVideo = 2
	UserProfileImage  = 3
)

func SingleUpload(ctx context.Context, file graphql.Upload) (*model.Media, error) {
	var (
		media        model.Media
		tok          *token.Token
		err          error
		documentType int
	)
	tok, err = token.GetFromContext(ctx)
	if err != nil {
		return &media, errx.UnAuthorizedError
	}

	if tok.UserId == state.ZERO {
		return &media, errx.UnAuthorizedError
	}

	mType, err := mimetype.DetectReader(file.File)
	if err != nil {
		return &media, errx.TypeError

	}
	if !utils.IsValidFile(mType.String()) {
		return &media, errx.TypeError
	}

	media.FileName = file.Filename
	media.Extension = filepath.Ext(file.Filename)
	media.Xid = xid.New().String()
	uploadPath := "./" + utils.FILE_UPLOAD_DIR + media.Xid + media.Extension

	f, err := os.Create(uploadPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, file.File)
	if err != nil {
		return &media, nil
	}

	err = f.Sync()
	if err != nil {
		return &media, nil
	}

	_, err = database.InsertOne(media)
	if err != nil {
		return &media, errx.DbInsertError
	}

	if utils.IsValidImage(mType.String()) {
		documentType = UserProfileImage
		err = utils.CreateThumb(media.Xid, media.Extension, file)
		if err != nil {
			return &media, errx.ThumbError
		}
	}

	if utils.IsValidDocument(mType.String()) {
		documentType = CV
		err = utils.CreateDocumentThumb(media.Xid, media.Extension, file)
		if err != nil {
			return &media, errx.ThumbError
		}
	}
	if utils.IsValidVideo(mType.String()) {
		documentType = PresentationVideo
	}

	err = SetUserMediaDetail(documentType, tok.UserId, media.Xid)
	if err != nil {
		return &media, errx.DbInsertError
	}

	return &media, nil
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