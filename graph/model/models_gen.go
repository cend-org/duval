// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Asset struct {
	ID          int        `json:"Id"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	UpdatedAt   time.Time  `json:"UpdatedAt"`
	DeletedAt   *time.Time `json:"DeletedAt,omitempty"`
	Description string     `json:"Description"`
}

type AssetInput struct {
	Description *string `json:"Description,omitempty"`
}

type Authorization struct {
	ID          int        `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
	UserID      int        `json:"userId"`
	AccessLevel int        `json:"AccessLevel"`
}

type Mutation struct {
}

type Password struct {
	ID        int        `json:"Id"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt,omitempty"`
	UserID    int        `json:"UserId"`
	Hash      string     `json:"Hash"`
}

type PasswordInput struct {
	Hash *string `json:"Hash,omitempty"`
}

type Query struct {
}

type School struct {
	ID        int        `json:"Id"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt,omitempty"`
	Name      string     `json:"Name"`
}

type SchoolInput struct {
	Name *string `json:"Name,omitempty"`
}

type SchoolSubject struct {
	ID           int        `json:"Id"`
	CreatedAt    time.Time  `json:"CreatedAt"`
	UpdatedAt    time.Time  `json:"UpdatedAt"`
	DeletedAt    *time.Time `json:"DeletedAt,omitempty"`
	SchoolNumber int        `json:"SchoolNumber"`
	Name         string     `json:"Name"`
}

type SchoolSubjectInput struct {
	SchoolNumber *int    `json:"SchoolNumber,omitempty"`
	Name         *string `json:"Name,omitempty"`
}

type User struct {
	ID        int        `json:"Id"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt,omitempty"`
	Name      string     `json:"Name"`
	LastName  string     `json:"LastName"`
	Email     string     `json:"Email"`
}

type UserInput struct {
	Name     *string `json:"Name,omitempty"`
	LastName *string `json:"LastName,omitempty"`
	Email    *string `json:"Email,omitempty"`
}
