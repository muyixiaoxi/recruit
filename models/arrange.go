package models

type Arrange struct {
	ID       uint       `json:"id"`
	Type     string     `json:"type"`                                                // 安排类型
	Place    string     `json:"place"`                                               // 地点
	Name     string     `json:"Name"`                                                // 本场名字
	Status   int        `json:"status" gorm:"default:0;not null;comment:0:未开始 1：结束"` // 0 未开始 1 已结束
	Students []*Student `gorm:"many2many:student_arrange;constraint:false"`
}

type StudentArrange struct {
	StudentID uint `json:"student_id"`
	ArrangeID uint `json:"arrange_id"`
}
