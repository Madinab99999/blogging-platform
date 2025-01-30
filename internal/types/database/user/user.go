package database

import "context"

type User interface {
	CreateUser(context.Context, CreateUserReq) (CreateUserResp, error)
	FindUserByID(context.Context, FindUserByIDReq) (FindUserByIDResp, error)
	FindUserByEmail(context.Context, FindUserByEmailReq) (FindUserByEmailResp, error)
}
