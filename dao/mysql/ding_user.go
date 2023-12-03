package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruit/models"
)

// UpdateStudentsState 更改学生状态
func UpdateStudentsState(tx *gorm.DB, student *models.Student) (err error) {
	err = tx.Updates(student).Error
	if err != nil {
		zap.L().Error("TX.Updates(student) failed", zap.Error(err))
		return err
	}
	return
}

// GetDetailInfo 获取详细信息
func GetDetailInfo(s *models.Student) (err error) {
	err = DB.Preload("TimeArrange").Find(s).Error
	return err
}

// GetStudent 初始化学生信息
func GetStudent(s *models.Student) {
	DB.Preload("TimeArrange").First(&s)
}

// ScreenStudents 筛选学生
func ScreenStudents(per *models.ParamScreen) ([]*models.Student, error) {
	/*筛选结果
	默认全部已报名1
	已宣讲2 已面试3 已通知4
	待宣讲5(正处于1)    待面试6（正处于2） 待通知7（正处于3）
	*/
	var data []*models.Student
	var err error
	if per.State <= 4 {
		err = DB.Where("IF (? = '',true,gender = ?) AND IF (? = '',true,state >= ?) AND IF (? = '',true,name like ?)", per.Gender, per.Gender, per.State, per.State, per.Name, "%"+per.Name+"%").Find(&data).Error
	} else {
		err = DB.Where("IF (? = '',true,gender = ?) AND IF (? = '',true,state = ?) AND IF (? = '',true,name like ?)", per.Gender, per.Gender, per.State, per.State-4, per.Name, "%"+per.Name+"%").Find(&data).Error
	}
	return data, err
}

// UserLogin 钉钉登录
func UserLogin(user *models.User) error {
	res := DB.Where("mobile =? and password = ?", user.Mobile, user.Password).First(user)
	return res.Error
}

// AddMsg 添加通知
func AddMsg(message *models.Message) error {
	res := DB.Create(message)
	return res.Error
}

// GetAllSignUp 获取所有报名人员信息
func GetAllSignUp() ([]*models.Student, error) {
	var data []*models.Student
	// state = 1 已报名
	res := DB.Where("state >= 1").Find(&data)
	if res.Error != nil {
		zap.L().Error("DB.First(&list) failed", zap.Error(res.Error))
	}
	return data, res.Error
}
