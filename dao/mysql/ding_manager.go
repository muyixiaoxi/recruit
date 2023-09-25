package mysql

import (
	"go.uber.org/zap"
	"recruit/models"
)

// AddMsg 添加信息
func AddMsg(message *models.Message) error {
	DB.Create(message)
	return DB.Error
}

// GetAllSignUp 获取所有报名人员信息
func GetAllSignUp() ([]*models.Student, error) {
	var data []*models.Student
	res := DB.First(&data)
	if res.Error != nil {
		zap.L().Error("DB.First(&list) failed", zap.Error(res.Error))
	}
	return data, res.Error
}
