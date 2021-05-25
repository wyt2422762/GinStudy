package model

type User struct{
	ID string `form:"id"`
	Username string `form:"username"`
	Password string `form:"password"`
	Name string `form:"name"`
	Email string `form:"email"`
	Telephone string `form:"telephone"`
}

type LoginUser struct{
	CacheKey string
	User
}

const LOGIN_TOKEN_KEY string = "login_token:"
