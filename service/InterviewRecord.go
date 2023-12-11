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
func GetInterviewRecord(id uint) (*models.RequestRecord, error) {
	// 获取面试组
	arrange := &models.Arrange{
		ID: id,
	}

	data, err := mysql.GetInterviewRecord(arrange)
	if err != nil {
		return nil, err
	}
	return tranRequestRecord(*data), nil
}

func tranRequestRecord(record models.Arrange) *models.RequestRecord {
	res := models.RequestRecord{
		ID:       record.ID,
		Type:     record.Type,
		Place:    record.Place,
		Name:     record.Name,
		Status:   record.Status,
		Students: make([]models.RecordStudent, len(record.Students)),
	}
	for i, student := range record.Students {
		child := models.RecordStudent{
			ID:     student.ID,
			Name:   student.Name,
			Record: student.InterviewRecord,
		}
		res.Students[i] = child
	}
	return &res
}
