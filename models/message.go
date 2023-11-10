package models

import "gorm.io/gorm"

// Message 自定义消息模板
type Message struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(30)"`
	Content     string `json:"content" gorm:"type:longtext"`
	StudentsID  string `json:"students_id"`
	State       int    `json:"state" gorm:"type:int"` //0 草稿 1消息
	IsRead      bool   `json:"is_read" gorm:"-"`
	Flag        bool   `json:"flag" gorm:"-"`
	ReadMessage []ReadMessage
	Students    []*Student `gorm:"-"`
}

// ReadMessage 已读消息表
type ReadMessage struct {
	gorm.Model
	MessageID uint `json:"message_id"`
	StudentID uint `json:"student_id"`
}
