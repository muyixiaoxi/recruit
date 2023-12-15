package mysql

import "gorm.io/gorm"

// DeleteArrangeStudents 删除学生
func DeleteArrangeStudents(tx *gorm.DB, ids []int) error {
	err := tx.Table("student_arrange").Where("arrange_id in ?", ids).Delete(nil).Error
	return err
}

// DeleteStudentArrangesByTx 删除面试安排表记录
func DeleteStudentArrangesByTx(tx *gorm.DB, arrangeId uint, ids []int) error {
	err := tx.Table("student_arrange").Where("arrange_id= ? and student_id in ?", arrangeId, ids).Delete(nil).Error
	return err
}
