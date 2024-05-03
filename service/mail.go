package service

import (
	"gopkg.in/gomail.v2"
	"recruit/dao/mysql"
	"recruit/models"
	"recruit/settings"
	"strings"
)

func SendMail(req models.SendMailRequest) (failIds []uint, err error) {
	// 获取学生邮箱
	mailTo, err := mysql.GetStudentMail(req.Ids)
	if err != nil {
		return
	}
	// 发送消息
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress("2860719956@qq.com", "三月软件官方"))
	m.SetHeader("Subject", req.Subject)
	for _, to := range mailTo {
		m.SetHeader("To", to.Mail)
		m.SetBody("text/html", strings.ReplaceAll(req.Message, "{{name}}", to.Name))
		d := gomail.NewDialer(settings.Conf.Mail.Host, settings.Conf.Mail.Port, settings.Conf.Mail.User, settings.Conf.Mail.Pass)
		err = d.DialAndSend(m)
		if err != nil {
			failIds = append(failIds, to.ID)
		}
	}
	return
}
