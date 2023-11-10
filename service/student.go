package service

import (
	"recruit/dao/mysql"
	"recruit/models"
)

// StudentLogin 学生登录
func StudentLogin(openid string, avtar string) error {
	s := &models.Student{
		Openid: openid,
		Avatar: avtar,
	}
	// 判断用户是否注册
	sid := mysql.GetStudentIdByOpenid(openid)
	if sid != 0 { // 用户已注册,同步头像
		return mysql.UpdateStudent(s)
	}
	// 未注册，注册
	return mysql.Signup(s)
}

// SignUp 报名
func SignUp(s *models.Student) error {
	s.State = 1
	return mysql.SignUp(s)
}

// GetSignUpInfoByOpenid 获取报名信息
func GetSignUpInfoByOpenid(openid string) (*models.Student, error) {
	return mysql.GetSignUpInfoByOpenid(openid)
}
