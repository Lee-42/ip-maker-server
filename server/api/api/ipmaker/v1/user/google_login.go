
package v1

import "github.com/gogf/gf/v2/frame/g"

// GoogleLoginReq defines the request for Google login.
type GoogleLoginReq struct {
	g.Meta      `path:"/user/google-login" method:"post" summary:"Google Login" tags:"User"`
	Credential  string `json:"credential" v:"required#Credential is required"`
}

// GoogleLoginRes defines the response for Google login.
type GoogleLoginRes struct {
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
	UserInfo  LoginUserInfo `json:"userInfo"`
}
