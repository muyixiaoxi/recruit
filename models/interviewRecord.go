package models

// InterviewRecord 面试记录
type InterviewRecord struct {
	ArrangeId uint   `json:"arrange_id" gorm:"primary_key;not null"`
	StudentID uint   `json:"student_id" gorm:"primary_key;not null;constraint:false"`
	ContentId uint   `json:"content_id" gorm:"not null"`
	Content   string `json:"content" gorm:"type:text;not null"`
	UserId    uint   `json:"-" gorm:"-"`
}
