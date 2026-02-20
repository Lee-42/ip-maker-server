package appin

// UserLoginOrRegisterByEmailInp defines the input for logging in or registering by email.
type UserLoginOrRegisterByEmailInp struct {
	Email    string `json:"email" v:"required|email#Email is required|Invalid email format"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
