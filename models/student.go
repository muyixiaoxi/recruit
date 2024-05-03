package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Openid           string             `json:"openid,omitempty" form:"openid" gorm:"unique;not null;comment:'openId'"`
	StudentNumber    string             `json:"student_number,omitempty" form:"student_number" gorm:"type:varchar(11);" binding:"required"`
	Name             string             `json:"name,omitempty" form:"name" binding:"required"`
	Gender           string             `json:"gender,omitempty" form:"gender" gorm:"type:varchar(1)" binding:"required"`
	Class            string             `json:"class,omitempty" form:"class" binding:"required"`
	Mail             string             `json:"mail,omitempty" gorm:"type:varchar(50);comment:'邮箱'"`
	QQ               string             `json:"QQ,omitempty" form:"QQ" gorm:"type:varchar(10)" binding:"required"`
	Phone            string             `json:"phone,omitempty" form:"phone" gorm:"type:varchar(11)" binding:"required"`
	Job              string             `json:"job,omitempty" form:"job" `
	SelfIntroduction string             `json:"selfIntroduction,omitempty" form:"selfIntroduction" `
	State            uint64             `json:"state,omitempty" gorm:"default:0"` // 0 未报名 1 已报名 2 已安排参观 3 已安排面试 4 已出结果
	InformState      uint64             `json:"inform_state,omitempty" gorm:"default:0"`
	TimeArrange      *TimeArrange       `gorm:"foreignKey,omitempty:StudentID"`
	Flag             bool               `json:"flag" gorm:"-"` // 前端单选框
	Avatar           string             `json:"avatar,omitempty"`
	IsRead           bool               `json:"is_read" gorm:"-"`
	Arranges         []*Arrange         `json:"arranges,omitempty" gorm:"many2many:student_arrange"`
	InterviewRecord  []*InterviewRecord `json:"interview_record,omitempty:StudentID" gorm:"constraint:false"`
}

// TimeArrange 时间安排
type TimeArrange struct {
	gorm.Model
	Visit     time.Time `json:"visit" `    // 参观
	Interview time.Time `json:"interview"` // 面试
	Enroll    string    `json:"enroll"`    // 录取结果
	StudentID uint      `json:"student_id"`
}
