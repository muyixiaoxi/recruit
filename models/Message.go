package models

import "gorm.io/gorm"

// Message 自定义消息模板
type Message struct {
	gorm.Model
	Content             string                `json:"content" gorm:"type:longtext"`
	NotificationMessage []NotificationMessage `gorm:"foreignkey:MsgId"`
}

// NotificationMessage 自定义通知消息
type NotificationMessage struct {
	gorm.Model
	MsgId  uint64 `json:"msg_id"`
	Openid string `json:"openid" `
	IsRead uint64 `json:"is_read" gorm:"default:0"` //0 未读 1 已读
}
