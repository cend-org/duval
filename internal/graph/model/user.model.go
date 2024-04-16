package model

import "time"

type User struct {
	Id                    uint       `json:"id"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeletedAt             *time.Time `json:"deleted_at"`
	Name                  string     `json:"name"`
	FamilyName            string     `json:"family_name"`
	NickName              string     `json:"nick_name"`
	Email                 string     `json:"email"`
	Matricule             string     `json:"matricule"`
	Age                   uint       `json:"age"`
	BirthDate             time.Time  `json:"birth_date"`
	Sex                   int        `json:"sex"`
	Lang                  int        `json:"language"`
	Status                int        `json:"status"`
	ProfileImageXid       string     `json:"profile_image_xid"`
	Description           string     `json:"description"`
	CoverText             string     `json:"cover_text"`
	Profile               string     `json:"profile"`
	ExperienceDetail      string     `json:"experience_detail"`
	AdditionalDescription string     `json:"additional_description"`
	AddOnTitle            string     `json:"add_on_title"`
}

type NewUserInput struct {
	Name       string    `json:"name"`
	FamilyName string    `json:"family_name"`
	NickName   string    `json:"nick_name"`
	Email      string    `json:"email"`
	Matricule  string    `json:"matricule,omitempty"`
	Age        uint      `json:"age,omitempty"`
	BirthDate  time.Time `json:"birth_date"`
	Sex        int       `json:"sex"`
	Lang       int       `json:"language"`
}
