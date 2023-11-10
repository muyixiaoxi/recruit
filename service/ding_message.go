package service

import (
	"gorm.io/gorm"
	"recruit/dao/mysql"
	"recruit/models"
	"strings"
)

// GetDetailMsg 获取详细信息
func GetDetailMsg(message *models.Message) (err error) {
	if err = mysql.GetDetailMsg(message); err != nil { // 获取详细信息
		return
	}
	if len(message.StudentsID) == 0 { // 公告
		// 查询已读
		read, err := mysql.GetReadStudentByMessageId(message)
		if err != nil {
			return err
		}
		for i, _ := range read {
			read[i].IsRead = true
		}
		// 查询未读
		unread, err := mysql.GetUnreadStudentByMessageId(message)
		if err != nil {
			return err
		}
		message.Students = append(read, unread...)

	} else { // 给部分学生发送消息
		sids := strings.Split(message.StudentsID, ",")
		
		// 获取已读
		read, err := mysql.GetReadStudentByMessageId(message)
		if err != nil {
			return err
		}
		for i, _ := range read {
			read[i].IsRead = true
		}

		// 获取未读
		unread, err := mysql.GetUnreadStudentByMessageIdAndStudentIds(message, sids)
		message.Students = append(read, unread...)
		/*var students []*models.Student
		ss := strings.Split(message.StudentsID, ",")

		for _, v := range ss {
			id, err := strconv.Atoi(v)
			if err != nil {
				zap.L().Error("strconv.Atoi(v) failed", zap.Error(err))
				return err
			}
			// 获取学生信息
			var s models.Student
			s.ID = uint(id)
			mysql.GetStudent(&s)
			students = append(students, &s)
			// 判断是否读取消息
		}
		for _, student := range students {
			student.IsRead, err = mysql.GetStudentIsReadMessage(student, message.ID)
			if err != nil {
				return
			}
		}
		message.Students = students*/
	}
	return
}

// DeleteMsg 删除草稿箱
func DeleteMsg(par *models.ParamMsg) (err error) {
	// 遍历删除
	for _, id := range par.Id {
		draft := models.Message{Model: gorm.Model{ID: uint(id)}, State: 0}
		err := mysql.DeleteDraft(&draft)
		if err != nil {
			break
		}
	}
	return err
}

// UpdateDraft 修改草稿箱
func UpdateDraft(draft *models.Message) error {
	return mysql.UpdateDraft(draft)
}

// GetAllDraft 获取所有草稿
func GetAllDraft() (data []*models.Message, err error) {
	return mysql.GetAllDraft()
}

// AddDraft 添加草稿箱
func AddDraft(message *models.Message) error {
	return mysql.AddDraft(message)
}

// UserGetAllMsg 获取所有消息
func UserGetAllMsg() (data []*models.Message, err error) {
	return mysql.UserGetAllMsg()
}

// SendMsg 发送通知
func SendMsg(message *models.Message) error {
	return mysql.AddMsg(message)
}

// GetAllSignUp 获取所有报名人员信息
func GetAllSignUp() ([]*models.Student, error) {
	return mysql.GetAllSignUp()
}
