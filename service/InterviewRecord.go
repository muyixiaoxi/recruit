package service

import (
	"recruit/dao/mysql"
	"recruit/models"
)

// SendServer 向服务端发送消息
func SendServer(record *models.InterviewRecord) error {
	err := mysql.InterviewRecord(record)
	if err != nil {
		return err
	}

	for id, ch := range OnlineUser {
		if id != record.UserId {
			ch <- *record
		}
	}
	return nil
}

// GetInterviewRecord 获取面试记录
func GetInterviewRecord(id uint) (*models.Arrange, error) {
	// 获取面试组
	arrange := &models.Arrange{
		ID: id,
	}
	return mysql.GetInterviewRecord(arrange)
}
