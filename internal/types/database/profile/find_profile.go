package database

import "time"

type FindProfileReq interface {
	GetUserID() string
}

type FindProfileResp interface {
	GetUserID() string
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
