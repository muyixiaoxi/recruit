package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Id               string `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	StudentNumber    string `json:"student_number" gorm:"type:varchar(11)"`
	Name             string `json:"name" `
	Gender           int64  `json:"gender" "`
	Class            string `json:"class"`
	QQ               string `json:"QQ" gorm:"type:varchar(10)"`
	Phone            string `json:"phone" gorm:"type:varchar(11)"`
	Job              string `json:"job"`
	SelfIntroduction string `json:"selfIntroduction"`
}
