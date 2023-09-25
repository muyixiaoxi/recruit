package mysql

import "recruit/models"

// SignUp 报名  Signup 注册

// GetUnreadMsg 获取未读消息
func GetUnreadMsg(openid string) ([]*models.Message, error) {
	msg := make([]*models.Message, 0)
	DB.Find(&msg).Joins("JOIN notification_messages nmsg ON message.id =nmsg.msg_id").Where("nmsg.openid = ? AND is_read = 0", openid)
	return msg, DB.Error
}

// Signup 注册
func Signup(s *models.Student) error {
	res := DB.Create(s)
	return res.Error
}

// SignUp 报名
func SignUp(s *models.Student) error {
	res := DB.Updates(s)
	return res.Error
}

// GetMySignUp 获取我的报名信息
func GetMySignUp(openid string) (*models.Student, error) {
	var data = &models.Student{
		Openid: openid,
	}

	res := DB.First(data)
	return data, res.Error
}
