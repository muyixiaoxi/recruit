package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// UpdateInterviewTimeIsNil 更新面试时间为空
func UpdateInterviewTimeIsNil(arrange *models.TimeArrange) (err error) {
	err = TX.Table("time_arranges").Where("student_id = ?", arrange.StudentID).Update("interview", nil).Error
	if err != nil {
		zap.L().Error("TX.Table(\"time_arranges\").Where(\"student_id = ?\", arrange.StudentID).Update(\"interview\", nil) failed", zap.Error(err))
	}
	return err
}

// UpdateVisitTimeIsNil 更新宣讲时间为空
func UpdateVisitTimeIsNil(arrange *models.TimeArrange) (err error) {
	err = TX.Table("time_arranges").Where("student_id = ?", arrange.StudentID).Update("visit", nil).Error
	if err != nil {
		zap.L().Error("TX.Table(\"time_arranges\").Where(\"student_id = ?\", arrange.StudentID).Update(\"visit\", nil) failed", zap.Error(err))
	}
	return err
}

// GetAllArrangeGroup 获取所有面试组
func GetAllArrangeGroup() (arr []*models.Arrange, err error) {
	err = DB.Preload("Students").Preload("Students.TimeArrange").Find(&arr).Error
	if err != nil {
		zap.L().Error("DB.Preload(\"Students\").Find(arr) failed", zap.Error(err))
	}
	return
}

// AddStudentArrange 添加学生面试关系表
func AddStudentArrange(sA *models.StudentArrange) (err error) {
	err = TX.Create(sA).Error
	if err != nil {
		zap.L().Error("DB.Create(sA) failed", zap.Error(err))
	}
	return
}

// AddArrange 添加安排
func AddArrange(arrange *models.Arrange) (err error) {
	err = TX.Create(arrange).Error
	if err != nil {
		zap.L().Error("DB.Create(arrange) failed", zap.Error(err))
	}
	return
}

// UpdateArrangeTime 修改安排时间
func UpdateArrangeTime(ta *models.TimeArrange) error {
	res := TX.Where("student_id = ?", ta.StudentID).Updates(ta)
	if res.Error != nil {
		zap.L().Error("DB.Where(\"student_id = ?\", ta.StudentID).Updates(ta) failed", zap.Error(res.Error))
	}
	return res.Error
}

// GetVisitStudents 获取待宣讲学生人员列表
func GetVisitStudents() (s []*models.Student, err error) {
	err = DB.Raw("SELECT s.* FROM students s,time_arranges t WHERE s.state = 1 and s.id = t.student_id and t.visit IS NULL").Scan(&s).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if err != nil {
		zap.L().Error("DB.Preload(\"TimeArrange\", \"visit = DEFAULT\").Where(\"state = 1\").Find(&s) failed", zap.Error(err))
	}
	return
}

// GetInterviewStudents 获取待面试学生人员列表
func GetInterviewStudents(needVisit bool) (s []*models.Student, err error) {
	if needVisit {
		err = DB.Raw("SELECT s.* FROM students s,time_arranges t WHERE s.state = 2 and s.id = t.student_id and t.interview IS NULL").Scan(&s).Error
	} else {
		err = DB.Raw("SELECT s.* FROM students s,time_arranges t WHERE s.state IN(1,2) and s.id = t.student_id and t.interview IS NULL").Scan(&s).Error
	}
	if err != nil {
		zap.L().Error("DB.Preload(\"TimeArrange\", \"visit = DEFAULT\").Where(\"state = 1\").Find(&s) failed", zap.Error(err))
	}
	return
}
