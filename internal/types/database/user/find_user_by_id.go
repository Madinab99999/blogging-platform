package database

import "time"

type FindUserByIDReq interface {
	GetID() string
}

type FindUserByIDResp interface {
	GetID() string
	GetEmail() string
	GetFullName() string
	GetPasswordHash() string
	GetSalt() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
