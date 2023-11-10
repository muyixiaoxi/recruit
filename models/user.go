package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:longtext"`
	Mobile   string `json:"mobile" gorm:"type:longtext" binding:"required"` //手机号
	Password string `json:"password,omitempty" gorm:"type:longtext" binding:"required"`
	Token    string `json:"token" gorm:"-"`
	Role     int    `json:"role"gorm:"type:uint"`
}

// RecruitPermission 招新用户权限     表里无数据只能查看
type RecruitPermission struct {
	gorm.Model
	UserId     string `json:"user_id"`
	Permission uint   `json:"permission"`
}
