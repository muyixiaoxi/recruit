package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// GetUnreadStudentByMessageIdAndStudentIds 获取未读学生
func GetUnreadStudentByMessageIdAndStudentIds(message *models.Message, ids []string) (s []*models.Student, err error) {
	res := DB.Table("students").Joins("LEFT JOIN read_messages ON students.id = read_messages.student_id AND read_messages.message_id = ?", message.ID).Where("read_messages.student_id IS NULL AND students.id IN (?)", ids).Find(&s)
	if res.Error != nil {
		zap.L().Error("DB.Table(\"students\").Joins(\"LEFT JOIN read_messages ON students.id = read_messages.student_id\").Where(\"read_messages.student_id IS NULL\").Find(&s) failed", zap.Error(res.Error))
	}
	return s, res.Error
}

// GetReadStudentByMessageId 获取已读消息的学生
func GetReadStudentByMessageId(message *models.Message) (s []*models.Student, err error) {
	res := DB.Model(&models.Student{}).Joins("JOIN read_messages ON students.id = read_messages.student_id").Where("read_messages.message_id = ?", message.ID).Find(&s)
	if res.Error != nil {
		zap.L().Error("DB.Model(&models.Student{}).Joins(\"JOIN read_messages ON students.id = read_messages.student_id\").Where(\"read_messages.message = ?\",message.ID).Find(&s) failed", zap.Error(res.Error))
	}
	return s, res.Error
}

// GetUnreadStudentByMessageId 获取未读消息的学生
func GetUnreadStudentByMessageId(message *models.Message) (s []*models.Student, err error) {
	res := DB.Table("students").Joins("LEFT JOIN read_messages ON students.id = read_messages.student_id").Where("read_messages.student_id IS NULL AND read_messages.message_id = ?", message.ID).Find(&s)
	if res.Error != nil {
		zap.L().Error("DB.Table(\"students\").Joins(\"LEFT JOIN read_messages ON students.id = read_messages.student_id\").Where(\"read_messages.student_id IS NULL\").Find(&s) failed", zap.Error(res.Error))
	}
	return s, res.Error
}

// GetStudentIsReadMessage 获取消息是否已读
func GetStudentIsReadMessage(student *models.Student, mId uint) (bool, error) {
	res := DB.Where("message_id = ? AND student_id = ?", mId, student.ID).First(&models.ReadMessage{})
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		zap.L().Error("DB.Where(\"message_id = ? AND student_id = ?\", mId, student.ID).Find(&models.ReadMessage{}) failed", zap.Error(res.Error))
		return false, res.Error
	}
	return true, nil
}

// GetDetailMsg 获取详细信息
func GetDetailMsg(message *models.Message) (err error) {
	err = DB.First(message).Error
	if err != nil {
		zap.L().Error("DB.First(message) failed", zap.Error(err))
	}
	return
}

// DeleteDraft 删除草稿箱
func DeleteDraft(message *models.Message) error {
	err := DB.Delete(message).Error
	return err
}

// UpdateDraft 修改草稿箱
func UpdateDraft(message *models.Message) error {
	err := DB.Updates(message).Error
	return err
}

// GetAllDraft 获取所有草稿
func GetAllDraft() (data []*models.Message, err error) {
	err = DB.Where("state = 0").Find(&data).Error
	if err != nil || err != gorm.ErrRecordNotFound {
		return
	}
	return data, nil
}

// AddDraft 添加草稿箱
func AddDraft(message *models.Message) error {
	err := DB.Create(message).Error
	return err
}

// UserGetAllMsg 用户获取所有信息
func UserGetAllMsg() (data []*models.Message, err error) {
	err = DB.Where("state = 1").Find(&data).Error
	if err != gorm.ErrRecordNotFound {
		return data, err
	}
	return data, nil
}
