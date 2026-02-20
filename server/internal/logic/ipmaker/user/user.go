// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package user

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/token"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
	"hotgo/internal/service/ipmaker"
	"hotgo/utility/validate"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type sAppUser struct{}

func NewAppUser() *sAppUser {
	return &sAppUser{}
}

func init() {
	ipmaker.RegisterAppUser(NewAppUser())
}

// Register 用户注册
func (s *sAppUser) Register(ctx context.Context, in *appin.UserRegisterInp) (userId int64, err error) {
	// 验证用户名唯一性
	cols := dao.AppUser.Columns()
	if err = hgorm.IsUnique(ctx, &dao.AppUser, g.Map{cols.Username: in.Username}, "用户名已存在", 0); err != nil {
		return
	}

	// 验证手机号唯一性
	if in.Mobile != "" {
		if !validate.IsMobile(in.Mobile) {
			err = gerror.New("手机号格式不正确")
			return
		}
		if err = hgorm.IsUnique(ctx, &dao.AppUser, g.Map{cols.Mobile: in.Mobile}, "手机号已被注册", 0); err != nil {
			return
		}
	}

	// 验证邮箱唯一性
	if in.Email != "" {
		if !validate.IsEmail(in.Email) {
			err = gerror.New("邮箱格式不正确")
			return
		}
		if err = hgorm.IsUnique(ctx, &dao.AppUser, g.Map{cols.Email: in.Email}, "邮箱已被注册", 0); err != nil {
			return
		}
	}

	// 生成密码盐和密码哈希
	salt := grand.S(6)
	passwordHash := gmd5.MustEncryptString(in.Password + salt)

	// 插入用户数据
	data := g.Map{
		cols.Username:     in.Username,
		cols.PasswordHash: passwordHash,
		cols.Salt:         salt,
		cols.Mobile:       in.Mobile,
		cols.Email:        in.Email,
		cols.Status:       consts.StatusEnabled,
		cols.CreatedAt:    gtime.Now(),
		cols.UpdatedAt:    gtime.Now(),
	}

	userId, err = dao.AppUser.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		err = gerror.Wrap(err, "用户注册失败")
		return
	}

	return
}

// Logout 退出登录
func (s *sAppUser) Logout(ctx context.Context) error {
	return token.Logout(g.RequestFromCtx(ctx))
}

// Login 用户登录
func (s *sAppUser) Login(ctx context.Context, in *appin.UserLoginInp) (tokenString string, user *entity.AppUser, err error) {
	// 查询用户
	err = dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().Username, in.Username).Scan(&user)
	if err != nil {
		err = gerror.Wrap(err, "查询用户失败")
		return
	}

	if user == nil {
		err = gerror.New("用户不存在")
		return
	}

	// 验证用户状态
	if user.Status != consts.StatusEnabled {
		err = gerror.New("账号已被禁用")
		return
	}

	// 验证密码
	inputPasswordHash := gmd5.MustEncryptString(in.Password + user.Salt)
	if inputPasswordHash != user.PasswordHash {
		err = gerror.New("用户名或密码错误")
		return
	}

	// 生成JWT Token
	identity := &model.Identity{
		Id:       user.Id,
		Username: user.Username,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Avatar:   user.Avatar,
		App:      consts.AppApi,
		LoginAt:  gtime.Now(),
	}
	tokenString, _, err = token.Login(ctx, identity)
	if err != nil {
		err = gerror.Wrap(err, "生成Token失败")
		return
	}

	// 更新最后登录信息
	var clientIp string
	if r := g.RequestFromCtx(ctx); r != nil {
		clientIp = r.GetClientIp()
	}
	_, err = dao.AppUser.Ctx(ctx).
		Where(dao.AppUser.Columns().Id, user.Id).
		Data(g.Map{
			dao.AppUser.Columns().LastLoginAt: gtime.Now(),
			dao.AppUser.Columns().LastLoginIp: clientIp,
			dao.AppUser.Columns().UpdatedAt:   gtime.Now(),
		}).
		Update()

	if err != nil {
		g.Log().Warningf(ctx, "更新用户登录信息失败: %v", err)
		// 登录信息更新失败不影响登录流程
		err = nil
	}

	return
}

// LoginOrRegisterByEmail 邮箱登录或注册
func (s *sAppUser) LoginOrRegisterByEmail(ctx context.Context, in *appin.UserEmailLoginInp) (tokenString string, user *entity.AppUser, err error) {
	// 1. 根据邮箱查询用户
	cols := dao.AppUser.Columns()
	err = dao.AppUser.Ctx(ctx).Where(cols.Email, in.Email).Scan(&user)
	if err != nil {
		err = gerror.Wrap(err, "查询用户失败")
		return
	}

	// 2. 如果用户存在，验证密码和状态
	if user != nil {
		if user.Status != consts.StatusEnabled {
			err = gerror.New("账号已被禁用")
			return
		}

		inputPasswordHash := gmd5.MustEncryptString(in.Password + user.Salt)
		if inputPasswordHash != user.PasswordHash {
			err = gerror.New("邮箱或密码错误")
			return
		}
	} else {
		// 3. 如果用户不存在，自动注册
		salt := grand.S(6)
		passwordHash := gmd5.MustEncryptString(in.Password + salt)
		
		// 从邮箱地址中提取一个默认的用户名和昵称
		emailParts := strings.Split(in.Email, "@")
		username := emailParts[0] + "_" + grand.S(4)
		nickname := emailParts[0]

		// 插入新用户数据
		newUser := &entity.AppUser{
			Username:     username,
			Nickname:     nickname,
			PasswordHash: passwordHash,
			Salt:         salt,
			Email:        in.Email,
			Status:       consts.StatusEnabled,
			CreatedAt:    gtime.Now(),
			UpdatedAt:    gtime.Now(),
		}

		userId, err_ := dao.AppUser.Ctx(ctx).Data(newUser).InsertAndGetId()
		if err_ != nil {
			err = gerror.Wrap(err_, "用户注册失败")
			return
		}
		newUser.Id = userId
		user = newUser
	}

	// 4. 生成JWT Token
	identity := &model.Identity{
		Id:       user.Id,
		Username: user.Username,
		RealName: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		App:      consts.AppApi,
		LoginAt:  gtime.Now(),
	}
	tokenString, _, err = token.Login(ctx, identity)
	if err != nil {
		err = gerror.Wrap(err, "生成Token失败")
		return
	}

	// 5. 更新最后登录信息
	var clientIp string
	if r := g.RequestFromCtx(ctx); r != nil {
		clientIp = r.GetClientIp()
	}
	_, err = dao.AppUser.Ctx(ctx).
		Where(cols.Id, user.Id).
		Data(g.Map{
			cols.LastLoginAt: gtime.Now(),
			cols.LastLoginIp: clientIp,
			cols.UpdatedAt:   gtime.Now(),
		}).
		Update()

	if err != nil {
		g.Log().Warningf(ctx, "更新用户登录信息失败: %v", err)
		err = nil // 登录信息更新失败不影响登录流程
	}

	return
}

// GetUserInfo 获取用户信息
func (s *sAppUser) GetUserInfo(ctx context.Context, userId int64) (*entity.AppUser, error) {
	var user *entity.AppUser
	err := dao.AppUser.Ctx(ctx).
		Where(dao.AppUser.Columns().Id, userId).
		Scan(&user)

	if err != nil {
		return nil, gerror.Wrap(err, "获取用户信息失败")
	}

	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	// 清除敏感信息
	user.PasswordHash = ""
	user.Salt = ""

	return user, nil
}

// LoginOrRegisterByEmailForGoogle 通过邮箱登录或注册
func (s *sAppUser) LoginOrRegisterByEmailForGoogle(ctx context.Context, in *appin.UserLoginOrRegisterByEmailInp) (tokenString string, user *entity.AppUser, err error) {
	cols := dao.AppUser.Columns()

	// 1. 根据邮箱查询用户
	err = dao.AppUser.Ctx(ctx).Where(cols.Email, in.Email).Scan(&user)
	if err != nil {
		return "", nil, gerror.Wrap(err, "查询用户失败")
	}

	// 2. 如果用户存在，直接登录
	if user != nil {
		if user.Status != consts.StatusEnabled {
			return "", nil, gerror.New("账号已被禁用")
		}
	} else {
		// 3. 如果用户不存在，自动注册
		// 构造一个唯一的用户名，防止重复
		username := "google_" + grand.S(10)

		// 插入新用户数据
		newUser := &entity.AppUser{
			Username: username,
			Nickname: in.Nickname,
			Email:    in.Email,
			Avatar:   in.Avatar,
			Status:   consts.StatusEnabled,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}

		userId, err := dao.AppUser.Ctx(ctx).Data(newUser).InsertAndGetId()
		if err != nil {
			return "", nil, gerror.Wrap(err, "用户注册失败")
		}
		newUser.Id = userId
		user = newUser
	}

	// 4. 生成JWT Token
	identity := &model.Identity{
		Id:       user.Id,
		Username: user.Username,
		RealName: user.Nickname, // Use RealName instead of Nickname
		Email:    user.Email,
		Avatar:   user.Avatar,
		App:      consts.AppApi,
		LoginAt:  gtime.Now(),
	}
	tokenString, _, err = token.Login(ctx, identity)
	if err != nil {
		return "", nil, gerror.Wrap(err, "生成Token失败")
	}

	// 5. 更新最后登录信息
	var clientIp string
	if r := g.RequestFromCtx(ctx); r != nil {
		clientIp = r.GetClientIp()
	}
	_, err = dao.AppUser.Ctx(ctx).
		Where(cols.Id, user.Id).
		Data(g.Map{
			cols.LastLoginAt: gtime.Now(),
			cols.LastLoginIp: clientIp,
			cols.UpdatedAt:   gtime.Now(),
		}).
		Update()

	if err != nil {
		g.Log().Warningf(ctx, "更新用户登录信息失败: %v", err)
		err = nil // 登录信息更新失败不影响登录流程
	}

	return
}

// UpdateProfile 更新用户资料
func (s *sAppUser) UpdateProfile(ctx context.Context, userId int64, in *appin.UserUpdateProfileInp) error {
	cols := dao.AppUser.Columns()
	data := g.Map{
		cols.UpdatedAt: gtime.Now(),
	}

	if in.Nickname != "" {
		data[cols.Nickname] = in.Nickname
	}
	if in.Avatar != "" {
		data[cols.Avatar] = in.Avatar
	}
	if in.Sex > 0 {
		data[cols.Sex] = in.Sex
	}
	if in.Birthday != "" {
		data[cols.Birthday] = in.Birthday
	}

	_, err := dao.AppUser.Ctx(ctx).
		Where(cols.Id, userId).
		Data(data).
		Update()

	if err != nil {
		return gerror.Wrap(err, "更新用户资料失败")
	}

	return nil
}

// List 用户列表（后台管理）
func (s *sAppUser) List(ctx context.Context, in *appin.UserListInp) (list []*entity.AppUser, totalCount int, err error) {
	mod := dao.AppUser.Ctx(ctx)
	cols := dao.AppUser.Columns()

	// 条件筛选
	if in.Username != "" {
		mod = mod.WhereLike(cols.Username, "%"+in.Username+"%")
	}

	if in.Mobile != "" {
		mod = mod.WhereLike(cols.Mobile, "%"+in.Mobile+"%")
	}

	if in.Email != "" {
		mod = mod.WhereLike(cols.Email, "%"+in.Email+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(cols.Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	// 分页查询
	totalCount, err = mod.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取用户总数失败")
		return
	}

	err = mod.Page(in.Page, in.PerPage).
		OrderDesc(cols.Id).
		Scan(&list)

	if err != nil {
		err = gerror.Wrap(err, "查询用户列表失败")
		return
	}

	// 清除敏感信息
	for _, user := range list {
		user.PasswordHash = ""
		user.Salt = ""
	}

	return
}

// Edit 编辑用户（后台管理）
func (s *sAppUser) Edit(ctx context.Context, in *appin.UserEditInp) error {
	cols := dao.AppUser.Columns()

	// 验证用户名唯一性
	if err := hgorm.IsUnique(ctx, &dao.AppUser, g.Map{cols.Username: in.Username}, "用户名已存在", in.Id); err != nil {
		return err
	}

	// 验证手机号唯一性
	if in.Mobile != "" {
		if !validate.IsMobile(in.Mobile) {
			return gerror.New("手机号格式不正确")
		}
		if err := hgorm.IsUnique(ctx, &dao.AppUser, g.Map{cols.Mobile: in.Mobile}, "手机号已存在", in.Id); err != nil {
			return err
		}
	}

	// 验证邮箱唯一性
	if in.Email != "" {
		if !validate.IsEmail(in.Email) {
			return gerror.New("邮箱格式不正确")
		}
		if err := hgorm.IsUnique(ctx, &dao.AppUser, g.Map{cols.Email: in.Email}, "邮箱已存在", in.Id); err != nil {
			return err
		}
	}

	// 修改
	if in.Id > 0 {
		data := g.Map{
			cols.Username:  in.Username,
			cols.Nickname:  in.Nickname,
			cols.Mobile:    in.Mobile,
			cols.Email:     in.Email,
			cols.Avatar:    in.Avatar,
			cols.Sex:       in.Sex,
			cols.Status:    in.Status,
			cols.Remark:    in.Remark,
			cols.UpdatedAt: gtime.Now(),
		}

		// 如果修改密码
		if in.Password != "" {
			// 获取用户的salt
			var user *entity.AppUser
			err := dao.AppUser.Ctx(ctx).Where(cols.Id, in.Id).Scan(&user)
			if err != nil || user == nil {
				return gerror.New("用户不存在")
			}
			data[cols.PasswordHash] = gmd5.MustEncryptString(in.Password + user.Salt)
		}

		_, err := dao.AppUser.Ctx(ctx).Where(cols.Id, in.Id).Data(data).Update()
		if err != nil {
			return gerror.Wrap(err, "修改用户失败")
		}
		return nil
	}

	// 新增
	salt := grand.S(6)
	passwordHash := gmd5.MustEncryptString(in.Password + salt)

	data := g.Map{
		cols.Username:     in.Username,
		cols.Nickname:     in.Nickname,
		cols.PasswordHash: passwordHash,
		cols.Salt:         salt,
		cols.Mobile:       in.Mobile,
		cols.Email:        in.Email,
		cols.Avatar:       in.Avatar,
		cols.Sex:          in.Sex,
		cols.Status:       in.Status,
		cols.Remark:       in.Remark,
		cols.CreatedAt:    gtime.Now(),
		cols.UpdatedAt:    gtime.Now(),
	}

	_, err := dao.AppUser.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return gerror.Wrap(err, "新增用户失败")
	}

	return nil
}

// Delete 删除用户（后台管理）
func (s *sAppUser) Delete(ctx context.Context, in *appin.UserDeleteInp) error {
	_, err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.Wrap(err, "删除用户失败")
	}
	return nil
}

// UpdateStatus 更新用户状态（后台管理）
func (s *sAppUser) UpdateStatus(ctx context.Context, in *appin.UserStatusInp) error {
	_, err := dao.AppUser.Ctx(ctx).
		Where(dao.AppUser.Columns().Id, in.Id).
		Data(g.Map{
			dao.AppUser.Columns().Status:    in.Status,
			dao.AppUser.Columns().UpdatedAt: gtime.Now(),
		}).
		Update()

	if err != nil {
		return gerror.Wrap(err, "更新用户状态失败")
	}

	return nil
}

// View 查看用户详情（后台管理）
func (s *sAppUser) View(ctx context.Context, in *appin.UserViewInp) (*entity.AppUser, error) {
	var user *entity.AppUser
	err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().Id, in.Id).Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "获取用户详情失败")
	}

	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	// 清除敏感信息
	user.PasswordHash = ""
	user.Salt = ""

	return user, nil
}
