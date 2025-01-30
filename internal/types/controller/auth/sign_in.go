package controller

type SignInReq interface {
	GetEmail() string
	GetPassword() string
}

type SignInResp interface {
	GetAccessToken() string
	GetRefreshToken() string
}
