package service

import (
	"recruit/dao/mysql"
	"recruit/dao/redis"
	"recruit/models"
	"recruit/pkg/jwt"
)

// Signup 注册
func Signup(param *models.ParamPhoneAndAuthCode) (n int) {
	// n = 0 注册成功
	// n = 1 服务器忙
	// n = 2 用户存在
	// n = 3 验证码错误
	// 判断验证码是否正确
	if ok := redis.GetAuthCode(param); !ok {
		return 3
	}
	// 判断用户是否存在
	user := mysql.GetUserByPhone(param.Phone)

	// 用户不存在
	if user.ID == 0 {
		// 注册
		if err := mysql.AddUser(param.Phone); err != nil {
			return 1
		}
	} else { // 用户存在
		return 2
	}

	// 创建
	return 0
}

// LoginByPassword 通过密码登录
func Login(user *models.User) (ok bool, err error) {
	// 密码登录
	if user.Password != "" {
		ok = mysql.LoginByPassword(user)
		if !ok {
			return
		}
		// 生成JWT
		token, err := jwt.GenToken(user.ID)
		if err != nil {
			return ok, err
		}
		user.Token = token
		return ok, err
	}
	// 验证码登录
	param := models.ParamPhoneAndAuthCode{
		Phone:    user.Phone,
		AuthCode: user.AuthCode,
	}
	if ok := redis.GetAuthCode(&param); ok {
		// 生成JWT
		token, err := jwt.GenToken(user.ID)
		if err != nil {
			return ok, err
		}
		temp := mysql.GetUserByPhone(user.Phone)
		user.ID = temp.ID
		user.Token = token
		return ok, err
	}
	return ok, err
}
