package controller

import "time"

type UpdateProfileReq interface {
	GetFullName() *string
	GetEmail() *string
	GetBio() *string
	GetLocation() *string
	GetGithubUrl() *string
	GetFacebookUrl() *string
	GetTwitterUrl() *string
	GetInstagramUrl() *string
	GetLinkedinUrl() *string
	GetBirthDate() *time.Time
	GetProfilePicture() *string
}

type UpdateProfileResp interface{}
