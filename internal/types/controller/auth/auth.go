package controller

import (
	"context"
)

type Auth interface {
	SignUp(context.Context, SignUpReq) (SignUpResp, error)
	SignIn(context.Context, SignInReq) (SignInResp, error)
	GetNewAccessToken(context.Context, GetNewAccessTokenReq) (GetNewAccessTokenResp, error)
}
