// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"
)

type IUserV1 interface {
	Register(ctx context.Context, req *RegisterReq) (res *RegisterRes, err error)
	Login(ctx context.Context, req *LoginReq) (res *LoginRes, err error)
	Info(ctx context.Context, req *InfoReq) (res *InfoRes, err error)
	Update(ctx context.Context, req *UpdateReq) (res *UpdateRes, err error)
}
