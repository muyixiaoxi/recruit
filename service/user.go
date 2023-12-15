package service

import (
	"recruit/dao/mysql"
	"recruit/models"
)

const (
	Visited        = "已宣讲"
	Unvisited      = "未宣讲"
	Interviewed    = "已面试"
	UniInterviewed = "未面试"
	Enroll         = "录取"
)
const (
	ArrangeVisit     = "visit"
	ArrangeInterview = "interview"
)

// OnlineUser 在线用户
var OnlineUser = map[uint]chan models.InterviewRecord{}

// UpdateStudentsState 修改学生状态
func UpdateStudentsState(par *models.ParamStudents) (err error) {
	tx := mysql.DB.Begin()
	for _, v := range par.StudentsId {
		var student models.Student
		student.ID = v
		student.State = par.State
		err = mysql.UpdateStudentsState(tx, &student)
		if err != nil {
			tx.Rollback()
			break
		}
	}
	tx.Commit()
	return err
}

// GetDetailInfo 获取详细信息
func GetDetailInfo(student *models.Student) (err error) {
	return mysql.GetDetailInfo(student)
}

// ScreenStudents 筛选学生
func ScreenStudents(s *models.ParamScreen) ([]*models.Student, error) {
	/*筛选结果
	默认全部已报名1
	已宣讲2 已面试3
	待宣讲4 待面试5
	*/
	return mysql.ScreenStudents(s)
}

// UserLogin 用户登录
func UserLogin(user *models.User) error {
	return mysql.UserLogin(user)
}
