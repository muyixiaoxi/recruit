package models

import "gorm.io/gorm"

type interviewer struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(5)"` // 姓名
	Role int    `json:"role" gorm:"type:int"`        // 主、副n
}
