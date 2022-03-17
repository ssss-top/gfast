package model

// LoginForOauthReq 登陆参数
type LoginForOauthReq struct {
	AccessToken string `p:"access_token" v:"required#AccessToken不能为空"`
}

type GraphqlQueryUserProfileRes struct {
	Message string `json:"message"`
	Data    struct {
		OauthUser *OauthUser `json:"profile"`
	} `json:"data"`
}

type OauthUser struct {
	ID                  string   `json:"id"`
	Email               string   `json:"email"`
	EmailVerified       bool     `json:"email_verified"`
	SignupMethods       string   `json:"signup_methods"`
	GivenName           *string  `json:"given_name"`
	FamilyName          *string  `json:"family_name"`
	MiddleName          *string  `json:"middle_name"`
	Nickname            *string  `json:"nickname"`
	PreferredUsername   *string  `json:"preferred_username"`
	Gender              *string  `json:"gender"`
	Birthdate           *string  `json:"birthdate"`
	PhoneNumber         *string  `json:"phone_number"`
	PhoneNumberVerified *bool    `json:"phone_number_verified"`
	Picture             *string  `json:"picture"`
	Roles               []string `json:"roles"`
	CreatedAt           *int64   `json:"created_at"`
	UpdatedAt           *int64   `json:"updated_at"`
}
