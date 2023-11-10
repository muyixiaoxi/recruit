package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// UpdateInterviewState 同步学生面试状况
func UpdateInterviewState() error {
	res := DB.Exec("UPDATE `students` JOIN time_arranges ON students.id = time_arranges.student_id SET `state`=3 WHERE time_arranges.interview < NOW()")
	if res.Error != nil {
		zap.L().Error("DB.Exec(\"UPDATE `students` JOIN time_arranges ON students.id = time_arranges.student_id SET `state`=3 WHERE time_arranges.interview < NOW()\") failed", zap.Error(res.Error))
	}
	return res.Error
}

// UpdateVisitState 同步学生宣讲状况
func UpdateVisitState() error {
	res := DB.Exec("UPDATE `students` JOIN time_arranges ON students.id = time_arranges.student_id SET `state`=2 WHERE time_arranges.visit < NOW() AND students.id = 1;")
	if res.Error != nil {
		zap.L().Error("DB.Exec(\"UPDATE `students` JOIN time_arranges ON students.id = time_arranges.student_id SET `state`=2 WHERE time_arranges.visit < NOW() AND students.id = 1;\") failed", zap.Error(res.Error))
	}
	return res.Error
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
