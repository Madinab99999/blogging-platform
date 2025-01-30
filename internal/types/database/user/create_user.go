package database

import "time"

type CreateUserReq interface {
	GetEmail() string
	GetPasswordHash() string
	GetSalt() string
	GetFullName() string
}

type CreateUserResp interface {
	GetID() string
	GetEmail() string
	GetFullName() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
