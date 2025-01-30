package controller

import "time"

type GetProfileOfUserReq interface {
}

type GetProfileOfUserResp interface {
	GetFullName() string
	GetEmail() string
	GetBio() *string
	GetLocation() *string
	GetGithubUrl() *string
	GetFacebookUrl() *string
	GetTwitterUrl() *string
	GetInstagramUrl() *string
	GetLinkedinUrl() *string
	GetBirthDate() *time.Time
	GetProfilePicture() *string
	GetCountOfPosts() int
	GetCountOfComments() int
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
