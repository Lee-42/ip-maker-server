// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package ipmaker

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
)

type (
	IAppUser interface {
		// Register 用户注册
		Register(ctx context.Context, in *appin.UserRegisterInp) (userId int64, err error)
		// Login 用户登录
		Login(ctx context.Context, in *appin.UserLoginInp) (tokenString string, user *entity.AppUser, err error)
		// Logout 退出登录
		Logout(ctx context.Context) error
		// LoginOrRegisterByEmail 邮箱登录或注册
		LoginOrRegisterByEmail(ctx context.Context, in *appin.UserEmailLoginInp) (tokenString string, user *entity.AppUser, err error)
		// GetUserInfo 获取用户信息
		GetUserInfo(ctx context.Context, userId int64) (*entity.AppUser, error)
		// UpdateProfile 更新用户资料
		UpdateProfile(ctx context.Context, userId int64, in *appin.UserUpdateProfileInp) error
		// LoginOrRegisterByEmailForGoogle 通过邮箱登录或注册
		LoginOrRegisterByEmailForGoogle(ctx context.Context, in *appin.UserLoginOrRegisterByEmailInp) (tokenString string, user *entity.AppUser, err error)
		// List 用户列表（后台管理使用）
		List(ctx context.Context, in *appin.UserListInp) (list []*entity.AppUser, totalCount int, err error)
		// Edit 编辑用户（后台管理使用）
		Edit(ctx context.Context, in *appin.UserEditInp) error
		// Delete 删除用户（后台管理使用）
		Delete(ctx context.Context, in *appin.UserDeleteInp) error
		// UpdateStatus 更新用户状态（后台管理使用）
		UpdateStatus(ctx context.Context, in *appin.UserStatusInp) error
		// View 查看用户详情（后台管理使用）
		View(ctx context.Context, in *appin.UserViewInp) (*entity.AppUser, error)
	}
)

var (
	localAppUser IAppUser
)

func AppUser() IAppUser {
	if localAppUser == nil {
		panic("implement not found for interface IAppUser, forgot register?")
	}
	return localAppUser
}

func RegisterAppUser(i IAppUser) {
	localAppUser = i
}
