package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// CancelTimeVisit 取消参观
func CancelTimeVisit(tx *gorm.DB, ids []int) error {
	err := tx.Model(&models.TimeArrange{}).Where("id in ?", ids).Update("visit", nil).Error
	return err
}

// CancelTimeInterview 取消面试
func CancelTimeInterview(tx *gorm.DB, ids []int) error {
	err := tx.Model(&models.TimeArrange{}).Where("id in ?", ids).Update("interview", nil).Error
	return err
}

// UpdateInterviewTimeIsNil 更新面试时间为空
func UpdateInterviewTimeIsNil(tx *gorm.DB, arrange *models.TimeArrange) (err error) {
	err = tx.Table("time_arranges").Where("student_id = ?", arrange.StudentID).Update("interview", nil).Error
	if err != nil {
		zap.L().Error("TX.Table(\"time_arranges\").Where(\"student_id = ?\", arrange.StudentID).Update(\"interview\", nil) failed", zap.Error(err))
	}
	return err
}

// UpdateVisitTimeIsNil 更新宣讲时间为空
func UpdateVisitTimeIsNil(tx *gorm.DB, arrange *models.TimeArrange) (err error) {
	err = tx.Table("time_arranges").Where("student_id = ?", arrange.StudentID).Update("visit", nil).Error
	if err != nil {
		zap.L().Error("TX.Table(\"time_arranges\").Where(\"student_id = ?\", arrange.StudentID).Update(\"visit\", nil) failed", zap.Error(err))
	}
	return err
}
