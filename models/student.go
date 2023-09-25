package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Openid           string `json:"openid" form:"openid" gorm:"primaryKey"`
	StudentNumber    string `json:"student_number" form:"student_number" gorm:"type:varchar(11)" binding:"required"`
	Name             string `json:"name" form:"name" binding:"required"`
	Gender           string `json:"gender" form:"gender" gorm:"type:varchar(1)" binding:"required"`
	Class            string `json:"class" form:"class" binding:"required"`
	QQ               string `json:"QQ" form:"QQ" gorm:"type:varchar(10)" binding:"required"`
	Phone            string `json:"phone" form:"phone" gorm:"type:varchar(11)" binding:"required"`
	Job              string `json:"job" form:"job" `
	SelfIntroduction string `json:"selfIntroduction" form:"selfIntroduction" `
	State            uint64 `json:"state" gorm:"default:0"` // 0 未报名 1 已报名 2 参观 3 面试 4 结果
}
