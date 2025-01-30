package controller

type SignUpReq interface {
	GetEmail() string
	GetFullName() string
	GetPassword() string
}

type SignUpResp interface {
	GetAccessToken() string
	GetRefreshToken() string
}
