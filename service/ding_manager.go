package service

import (
	"recruit/dao/mysql"
	"recruit/models"
)

// SendMsg 发送通知
func SendMsg(message *models.Message) error {
	return mysql.AddMsg(message)
}

// GetAllSignUp 获取所有报名人员信息
func GetAllSignUp() ([]*models.Student, error) {
	return mysql.GetAllSignUp()
}
