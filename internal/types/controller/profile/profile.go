package controller

import (
	"context"
)

type Profile interface {
	UpdateProfile(context.Context, UpdateProfileReq) (UpdateProfileResp, error)
	GetProfileOfUser(context.Context, GetProfileOfUserReq) (GetProfileOfUserResp, error)
	ListProfiles(context.Context, ListProfilesReq) (ListProfilesResp, error)
}
