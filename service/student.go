package service

import (
	"gorm.io/gorm"
	"recruit/dao/mysql"
	"recruit/models"
)

// GetUnreadMsg 获取未读消息
func GetUnreadMsg(openid string) ([]*models.Message, error) {
	// 1、获取所有发送的消息
	// 2、获取未读

	return mysql.GetUnreadMsg(openid)
}

// Signup 注册
func Signup(openid string) error {
	// 判断用户是否注册
	_, err := mysql.GetMySignUp(openid)
	// 未注册过，注册
	if err == gorm.ErrRecordNotFound {
		var student = models.Student{Openid: openid}
		return mysql.Signup(&student)
	}
	return err
}

// SignUp 新生报名
func SignUp(s *models.Student) error {
	return mysql.SignUp(s)
}

// GetMySignUp 获取我的报名信息
func GetMySignUp(openId string) (*models.Student, error) {
	return mysql.GetMySignUp(openId)
}
