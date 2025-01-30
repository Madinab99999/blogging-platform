package controller

type GetNewAccessTokenReq interface {
	GetRefreshToken() string
}

type GetNewAccessTokenResp interface {
	GetAccessToken() string
}
