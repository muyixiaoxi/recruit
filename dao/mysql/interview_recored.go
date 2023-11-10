package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// InterviewRecord 面试记录
func InterviewRecord(record *models.InterviewRecord) error {
	res := DB.Where("arrange_id =? and student_id =? and content_id = ?", record.ArrangeId, record.StudentID, record.ContentId).Save(record)
	if res.Error != nil {
		zap.L().Error("DB.Save(record) failed", zap.Error(res.Error))
	}
	return res.Error
}

// GetInterviewRecord 获取面试记录
func GetInterviewRecord(arrange *models.Arrange) (*models.Arrange, error) {
	res := DB.Preload("Students.InterviewRecord", func(db *gorm.DB) *gorm.DB {
		return db.Where("arrange_id = ?", arrange.ID)
	}).First(arrange)
	if res.Error != nil {
		zap.L().Error("DB.Preload(\"Students\").First(arrange) failed", zap.Error(res.Error))
	}
	return arrange, res.Error
}
