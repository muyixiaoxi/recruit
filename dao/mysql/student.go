package mysql

import (
	"go.uber.org/zap"
	"recruit/models"
)

// SignUp 报名  Signup 注册

// GetStudentIdByOpenid 获取学生id
func GetStudentIdByOpenid(openid string) uint {
	s := models.Student{
		Openid: openid,
	}
	DB.Where("openid = ?", openid).First(&s)
	return s.ID
}

// GetAllMsg 获取全部消息
func GetAllMsg() ([]*models.Message, error) {
	msg := make([]*models.Message, 0)
	res := DB.Where("state = 1").Find(&msg)
	if res.Error != nil {
		zap.L().Error(" DB.Find(&msg) failed", zap.Error(res.Error))
	}

	return msg, DB.Error
}

// Signup 注册
func Signup(s *models.Student) error {
	res := DB.Create(s)
	if res.Error != nil {
		zap.L().Error("DB.Create(s) failed", zap.Error(res.Error))
	}
	return res.Error
}

// SignUp 报名
func SignUp(s *models.Student) error {
	res := DB.Where("openid = ?", s.Openid).Updates(s)
	if res.Error != nil {
		zap.L().Error("DB.Updates(s) failed", zap.Error(res.Error))
	}
	return res.Error
}

// GetSignUpInfoByOpenid 获取报名信息
func GetSignUpInfoByOpenid(openid string) (*models.Student, error) {
	var data = &models.Student{
		Openid: openid,
	}

	res := DB.First(data)
	return data, res.Error
}

// UpdateStudent 修改学生信息
func UpdateStudent(student *models.Student) error {
	res := DB.Updates(student)
	return res.Error
}
