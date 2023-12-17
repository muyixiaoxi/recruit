package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// GetStudentInSeven 获取七天内报道的学生
func GetStudentInSeven() ([]models.ReportData, error) {
	var data []models.ReportData

	// 使用 RAW SQL 查询
	err := DB.Raw("SELECT DATE(created_at) as date, COUNT(*) as count FROM students WHERE created_at >= CURDATE() - INTERVAL 6 DAY GROUP BY DATE(created_at) ORDER BY DATE(created_at)").Scan(&data).Error

	return data, err
}

// StudentStatusAdd 学生状态+1
func StudentStatusAdd(tx *gorm.DB, ids []uint) error {
	return tx.Model(&models.Student{}).Where("id in ?", ids).Update("state", gorm.Expr("state+1")).Error
}

// StudentStatusSub 学生状态-1
func StudentStatusSub(tx *gorm.DB, ids []int) error {
	return tx.Model(&models.Student{}).Where("id in ?", ids).Update("state", gorm.Expr("state-1")).Error
}

// GetAllSignUpAndTimeArrange 获取所有已报名学生及安排表
func GetAllSignUpAndTimeArrange() (students []*models.Student, err error) {
	res := DB.Preload("TimeArrange").Where("state > 0").Find(&students)
	if res.Error != nil {
		zap.L().Error("DB.Preload(\"TimeArrange\").Where(\"state > 0\").Find(students) failed", zap.Error(res.Error))
	}
	return students, res.Error
}

// GetSignUpBySpecialty 通过专业获取报名学生
func GetSignUpBySpecialty(specialty string) (students []*models.Student, err error) {
	res := DB.Preload("TimeArrange").Where("state > 0 AND class like ?", "%"+specialty+"%").Find(&students)
	if res.Error != nil {
		zap.L().Error("DB.Preload(\"TimeArrange\").Where(\"state > 0 AND class like ?\", \"%\"+specialty+\"%\").Find(students) failed", zap.Error(res.Error))
	}
	return students, res.Error
}

// GetInterviewedStudent 获取已面试学生
func GetInterviewedStudent() (students []*models.Student, err error) {
	res := DB.Preload("TimeArrange", func(db *gorm.DB) *gorm.DB {
		return db.Where("interview IS NULl OR interview < CURDATE()")
	}).Where("state > 0").Find(&students)
	if res.Error != nil {
		zap.L().Error("DB.Preload(\"TimeArrange\", func(db *gorm.DB) *gorm.DB {\n\t\treturn db.Where(\"interview IS NULl OR interview < CURDATE()\")\n\t})"+
			".Where(\"state > 0\").Find(&students) failed", zap.Error(res.Error))
	}
	return students, res.Error
}

// GetEnrollSuccess 获取成功录取
func GetEnrollSuccess() (students []*models.Student, err error) {
	res := DB.Table("students").Joins("join time_arranges on students.id = time_arranges.student_id").Where("time_arranges.enroll='录取'").Scan(&students)
	return students, res.Error
}

// GetEnrollSuccessBySpecialty 通过专业获取成功录取
func GetEnrollSuccessBySpecialty(specialty string) (students []*models.Student, err error) {
	res := DB.Table("students").Joins("join time_arranges on students.id = time_arranges.student_id").Where("time_arranges.enroll='录取' AND students.class like ?", specialty).Scan(&students)
	return students, res.Error
}

// GetSignUpBySpecialtyAndGender 通过专业和性别获取学生
func GetSignUpBySpecialtyAndGender(specialty string, gender string) (students []*models.Student, err error) {
	res := DB.Where("state > 0 AND gender = ? AND class like ? ", gender, "%"+specialty+"%").Find(&students)
	if res.Error != nil {
		zap.L().Error("DB.Where(\"state > 0 AND gender = ? AND class like ? \", gender, \"%\"+specialty+\"%\").Find(&students) failed", zap.Error(res.Error))
	}
	return students, res.Error
}

// GetSignUpByGender 获取报名学生
func GetSignUpByGender(gender string) (students []*models.Student, err error) {
	res := DB.Where("state > 0 AND gender = ?", gender).Find(&students)
	if res.Error != nil {
		zap.L().Error("DB.Where(\"state > 0 AND gender = ?\",gender).Find(&students) failed", zap.Error(res.Error))
	}
	return students, res.Error
}
