package database

import "context"

type Profile interface {
	FindProfile(context.Context, FindProfileReq) (FindProfileResp, error)
	UpdateProfile(context.Context, UpdateProfileReq) (UpdateProfileResp, error)
	ListProfiles(context.Context, ListProfilesReq) (ListProfilesResp, error)
}
