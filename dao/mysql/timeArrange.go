package mysql

import (
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
