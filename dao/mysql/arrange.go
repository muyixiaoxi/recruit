package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// GetArrangeDetail 获取面试组详细信息
func GetArrangeDetail(id uint) (arrange models.Arrange, err error) {
	err = DB.Where("id = ?", id).Preload("Students.TimeArrange").First(&arrange).Error
	return arrange, err
}

// DeleteArrangeStudents 删除学生
func DeleteArrangeStudents(tx *gorm.DB, ids []int) error {
	err := tx.Model(&models.StudentArrange{}).Delete("arrange_id in ?", ids).Error
	return err
}

// DeleteArranges 删除
func DeleteArranges(tx *gorm.DB, ids []int) error {
	err := tx.Model(&models.Arrange{}).Delete("id in ?", ids).Error
	return err
}

// GetArrangeMenus 获取安排菜单
func GetArrangeMenus(t int) (arranges []models.Arrange, err error) {
	if t == 1 {
		err = DB.Find(&arranges).Error
	} else if t == 2 {
		err = DB.Where("type = 'interview'").Find(&arranges).Error
	} else if t == 3 {
		err = DB.Where("type = 'visit'").Find(&arranges).Error
	}
	return arranges, err
}

// DeleteArrangeVisit 删除宣讲
func DeleteArrangeVisit(tx *gorm.DB, ids []int) (err error) {
	err = tx.Model(&models.Student{}).Joins("join arranges on arranges.id = student_arrange.arrange_id").Where("arranges.type = visit and  student_arrange.student_id in ?", ids).Error
	return err
}

// DeleteArrangeInterview 删除面试
func DeleteArrangeInterview(tx *gorm.DB, ids []int) (err error) {
	err = tx.Model(&models.Student{}).Joins("join arranges on arranges.id = student_arrange.arrange_id").Where("arranges.type = interview and  student_arrange.student_id in ?", ids).Error
	return err
}

// UpdateInterviewTimeIsNil 更新面试时间为空
func UpdateInterviewTimeIsNil(tx *gorm.DB, arrange *models.TimeArrange) (err error) {
	err = tx.Table("time_arranges").Where("student_id = ?", arrange.StudentID).Update("interview", nil).Error
	if err != nil {
		zap.L().Error("TX.Table(\"time_arranges\").Where(\"student_id = ?\", arrange.StudentID).Update(\"interview\", nil) failed", zap.Error(err))
	}
	return err
}

// UpdateVisitTimeIsNil 更新宣讲时间为空
func UpdateVisitTimeIsNil(tx *gorm.DB, arrange *models.TimeArrange) (err error) {
	err = tx.Table("time_arranges").Where("student_id = ?", arrange.StudentID).Update("visit", nil).Error
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
func AddStudentArrange(tx *gorm.DB, sA *models.StudentArrange) (err error) {
	err = tx.Create(sA).Error
	if err != nil {
		zap.L().Error("DB.Create(sA) failed", zap.Error(err))
	}
	return
}

// AddArrange 添加安排
func AddArrange(tx *gorm.DB, arrange *models.Arrange) (err error) {
	err = tx.Create(arrange).Error
	if err != nil {
		zap.L().Error("DB.Create(arrange) failed", zap.Error(err))
	}
	return
}

// UpdateArrangeTime 修改安排时间
func UpdateArrangeTime(tx *gorm.DB, ta *models.TimeArrange) error {
	res := tx.Where("student_id = ?", ta.StudentID).Updates(ta)
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
