package appin

// UserEmailLoginInp 邮箱登录
type UserEmailLoginInp struct {
	Email    string `json:"email"    v:"required|email#请输入邮箱|邮箱格式不正确"`
	Password string `json:"password" v:"required#请输入密码"`
}
