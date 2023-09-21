package models

type ParamPhoneAndAuthCode struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	AuthCode string `json:"auth_code" form:"auth_code"`
}

type ParamLogin struct {
	Phone    string `form:"phone" binding:"required"`
	AuthCode string `form:"auth_code"`
	Password string `form:"password"`
}
