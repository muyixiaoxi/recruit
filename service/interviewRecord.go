package service

import (
	"gorm.io/gorm"
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

// InterviewRecordInit 面试记录初始化
func InterviewRecordInit(tx *gorm.DB, arrange uint, student []uint) error {
	// 默认三列
	data := make([]models.InterviewRecord, 3*len(student)+3)
	for i := 0; i < len(data)/3; i++ {
		if i == 0 {
			data[i*3] = models.InterviewRecord{
				ArrangeId: arrange,
				StudentID: 0,
				ContentId: uint(0),
				Content:   "面试官一",
			}
			data[i*3+1] = models.InterviewRecord{
				ArrangeId: arrange,
				StudentID: 0,
				ContentId: uint(1),
				Content:   "面试官二",
			}
			data[i*3+2] = models.InterviewRecord{
				ArrangeId: arrange,
				StudentID: 0,
				ContentId: uint(2),
				Content:   "面试官三",
			}
		} else {
			data[i*3] = models.InterviewRecord{
				ArrangeId: arrange,
				StudentID: student[i-1],
				ContentId: uint(0),
			}
			data[i*3+1] = models.InterviewRecord{
				ArrangeId: arrange,
				StudentID: student[i-1],
				ContentId: uint(1),
			}
			data[i*3+2] = models.InterviewRecord{
				ArrangeId: arrange,
				StudentID: student[i-1],
				ContentId: uint(2),
			}
		}
	}

	data = append(data)
	return mysql.AddInterviewRecord(tx, data)
}

// AddContentId 添加内容id
func AddContentId(id uint) error {
	arrange, err := mysql.GetArrangeDetail(id)
	if err != nil {
		return err
	}
	n, err := mysql.GetInterviewRecordCol(id)
	cId := uint(n) + 1
	if err != nil {
		return err
	}
	inter := make([]models.InterviewRecord, len(arrange.Students)+1)
	inter[len(arrange.Students)] = models.InterviewRecord{
		ArrangeId: id,
		StudentID: 0,
		ContentId: cId,
	}
	for i, student := range arrange.Students {
		inter[i] = models.InterviewRecord{
			ArrangeId: id,
			StudentID: student.ID,
			ContentId: cId,
		}
	}
	err = mysql.AddInterviewRecord(mysql.DB, inter)
	return err
}

// DeleteContentId 删除指定列
func DeleteContentId(arrange, content int) error {
	return mysql.DeleteContentId(arrange, content)
}
