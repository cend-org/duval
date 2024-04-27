package utils

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cend-org/duval/internal/database"
	"github.com/disintegration/imaging"
	"github.com/joinverse/xid"
	"github.com/unidoc/unipdf/v3/model"
	"github.com/unidoc/unipdf/v3/render"
	"image"
	"image/color"
	"io"
	"time"
)

const (
	CV                = 0
	Letter            = 1
	VideoPresentation = 2
	UserProfileImage  = 3
)

type MediaThumb struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Extension string     `json:"extension"`
	MediaXid  string     `json:"media_xid"`
	Xid       string     `json:"xid"`
}

type UserMediaDetail struct {
	Id           uint       `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	OwnerId      uint       `json:"owner_id"`
	DocumentType uint       `json:"document_type"`
}

/*
CREATE THUMBNAIL FOR UPLOADED IMAGE
*/

func CreateThumb(mediaXid string, extension string, file graphql.Upload) (err error) {
	var (
		mediaThumb MediaThumb
		thumbnail  image.Image
	)

	if _, err := file.File.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek file: %w", err)
	}

	img, err := imaging.Decode(file.File)
	if err != nil {
		return err
	}

	thumbnail = imaging.Thumbnail(img, 200, 200, imaging.CatmullRom)

	dst := imaging.New(200, 200, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumbnail, image.Pt(0, 0))
	err = imaging.Save(dst, FILE_UPLOAD_DIR+THUMB_FILE_UPLOAD_DIR+mediaXid+extension)
	if err != nil {
		return
	}

	mediaThumb.Extension = extension
	mediaThumb.MediaXid = mediaXid
	mediaThumb.Xid = "T_" + xid.New().String()

	_, err = database.InsertOne(mediaThumb)
	if err != nil {
		return err
	}

	return
}

/*
CREATE THUMBNAIL FOR UPLOADED COVER LETTER
*/

func CreateDocumentThumb(mediaXid string, extension string, file graphql.Upload) (err error) {
	var (
		mediaThumb MediaThumb
		thumbnail  image.Image
	)

	if _, err := file.File.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek file: %w", err)
	}

	pdfReader, err := model.NewPdfReader(file.File)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil || numPages < 1 {
		return err
	}

	firstPage, err := pdfReader.GetPage(1)
	if err != nil {
		return err
	}
	device := render.NewImageDevice()

	img, err := device.Render(firstPage)
	if err != nil {
		return err
	}
	thumbnail = imaging.Thumbnail(img, 800, 1100, imaging.CatmullRom)

	dst := imaging.New(800, 1100, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumbnail, image.Pt(0, 0))
	err = imaging.Save(dst, FILE_UPLOAD_DIR+THUMB_FILE_UPLOAD_DIR+mediaXid+".jpg")
	if err != nil {
		return err
	}

	mediaThumb.Extension = ".jpg"
	mediaThumb.MediaXid = mediaXid
	mediaThumb.Xid = "T_" + xid.New().String()

	_, err = database.InsertOne(mediaThumb)
	if err != nil {
		return err
	}
	return
}

/*
CREATE THUMBNAIL FOR UPLOADED Video
*/
