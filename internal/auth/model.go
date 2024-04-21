package auth

type User struct {
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
}

type UserLoginPassword struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

type LogInRequest struct {
	NickName string `json:"login"`
	Password string `json:"password"`
}

type LogInResponse struct {
	User          User   `json:"user_data"`
	Authenticated bool   `json:"auth"`
	Access        string `json:"access_token"`
}

type CheckLogInRequest struct {
	NickName string `json:"login" db:"login"`
	Password string `json:"password"`
}

type CheckLogInResponse struct {
	Authenticated bool `json:"is_authenticated"`
}

type GetUserByLoginRequest struct {
	NickName string `json:"nickname"`
}

type GetUserByLoginResponse struct {
	UserResp User `json:"user"`
}
