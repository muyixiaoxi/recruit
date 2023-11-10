package service

import (
	"recruit/dao/mysql"
	"recruit/models"
)

// AddReadMsg 添加消息为已读
func AddReadMsg(rm *models.ReadMessage, openid string) error {
	sid := mysql.GetStudentIdByOpenid(openid)
	rm.StudentID = sid
	return mysql.AddReadMsg(rm)
}

// GetAllMsgByStudentOpenid 获取学生收到的全部消息
func GetAllMsgByStudentOpenid(openid string) ([]*models.Message, error) {
	return mysql.GetAllMsg()
}
