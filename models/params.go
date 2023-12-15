package models

import (
	"time"
)

type LoginResponse struct {
	Openid      string `json:"openid"`
	SessionKey  string `json:"session_key"`
	AccessToken string `json:"access_token"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

// TemplateMessage 模板消息
type TemplateMessage struct {
	AccessToken string `json:"access_token"`
	Appid       string `json:"appid"`
	Touser      string `json:"touser"`
	TemplateId  string `json:"templateId"`
	FormId      string `json:"formId"`
}

// ParamScreen 筛选参数
type ParamScreen struct {
	Gender string `json:"gender" form:"gender"`
	State  uint64 `json:"state" form:"state"`
	Name   string `json:"name" form:"name"`
}

// ParamArrange 安排参数
type ParamArrange struct {
	StudentsId []uint    `json:"students_id"`
	Type       string    `json:"type"`
	Time       time.Time `json:"time"`
}

// ParamSendMsg 发送消息参数
type ParamSendMsg struct {
	Id         uint   `json:"id"` //草稿箱发送消息时传id，新建消息不传id
	Title      string `json:"title"`
	Content    string `json:"content"`
	StudentsId []int  `json:"students_id"`
}

// ParamArrangeGroup 面试组参数
type ParamArrangeGroup struct {
	Name      string    `json:"name"`       // 名称
	Type      string    `json:"type"`       // 面试/宣讲
	StartTime time.Time `json:"start_time"` // 开始时间
	EndTime   time.Time `json:"end_time"`   // 结束时间
	OnceTime  int       `json:"once_time"`  // 单场时间
	Place     []string  `json:"place"`      // 地点 可以同时有多场
	NeedVisit bool      `json:"need_visit"` // 安排面试时使用，是否需要宣讲
}

// ParamStudent 学生参数
type ParamStudent struct {
	Id uint `json:"id" form:"id"`
}

type ParamStudents struct {
	StudentsId []uint `json:"students_id"`
	State      uint64 `json:"state"`
}

type ParamLogin struct {
	Code   string `json:"code,"`
	Avatar string `json:"avatar"` // 头像
}

type ParamIds struct {
	Id []int `json:"id"`
}

type ParamGroupBySpecialty struct {
	Type           string `json:"type"`
	Sum            int    `json:"sum"`
	BoyNum         int    `json:"boy_num"`
	GirlNum        int    `json:"girl_num"`
	InterviewedNum int    `json:"interviewed_num"`
	EnrollNum      int    `json:"enroll_num"`
}

type ParamCancelArrangeTime struct {
	ArrangeID uint  `json:"arrange_id"`
	Ids       []int `json:"ids"`
	Type      int   `json:"type" binding:"required"` // 1宣讲，2面试
}

// ParamArrangeMenus 安排菜单
type ParamArrangeMenus struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Type   string `json:"type"`
}

type RequestRecord struct {
	ID       uint            `json:"id"`
	Type     string          `json:"type"`
	Place    string          `json:"place"`
	Name     string          `json:"name"`
	Status   int             `json:"status"`
	Students []RecordStudent `json:"Students"`
}

type RecordStudent struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Record []*InterviewRecord
}
