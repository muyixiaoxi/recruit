package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// DeleteContentId 删除指定列
func DeleteContentId(arrange, content int) error {
	return DB.Where("arrange_id = ? and content_id = ?", arrange, content).Delete(&models.InterviewRecord{}).Error
}

// GetInterviewRecordCol 获取列数
func GetInterviewRecordCol(arrangeId uint) (max int64, err error) {
	err = DB.Model(&models.InterviewRecord{}).Where("arrange_id = ? and student_id = 0", arrangeId).Select("MAX(content_id)").Scan(&max).Error
	return max, err
}

// AddInterviewRecord 增加面试记录
func AddInterviewRecord(tx *gorm.DB, records []models.InterviewRecord) error {
	return tx.Create(&records).Error
}

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
