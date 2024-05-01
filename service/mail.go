package service

import (
	"recruit/models"
)

func SendMail(req models.SendMailRequest) (err error) {
	// 获取学生邮箱
	/*ids := make([]*models.Student, len(req.Ids))
	for i, id := range req.Ids {
		ids[i].ID = id
	}
	mysql.GetStudentMail(ids)
	// 发送消息
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress("2860719956@qq.com", "三月软件官方"))
	m.SetHeader("To", mailTo)*/
	return err
}
