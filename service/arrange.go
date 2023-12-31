package service

import (
	"recruit/dao/mysql"
	"recruit/models"
	"time"
)

// GetArrangeDetail 获取面试组详细信息
func GetArrangeDetail(id uint) (models.Arrange, error) {
	return mysql.GetArrangeDetail(id)
}

// DeleteArrange 删除安排组
func DeleteArrange(ids []int) error {
	tx := mysql.DB.Begin()
	// 删除安排组
	if err := mysql.DeleteArranges(tx, ids); err != nil {
		tx.Rollback()
		return err
	}
	// 删除 安排组-学生
	err := mysql.DeleteArrangeStudents(tx, ids)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 学生状态-1
	err = mysql.StudentStatusSub(tx, ids)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// GetArrangeMenus 获取安排菜单
func GetArrangeMenus(t int) ([]models.ParamArrangeMenus, error) {
	arranges, err := mysql.GetArrangeMenus(t)
	if err != nil {
		return nil, err
	}
	data := make([]models.ParamArrangeMenus, len(arranges))
	for i, arrange := range arranges {
		tmp := models.ParamArrangeMenus{
			Id:     arrange.ID,
			Name:   arrange.Name,
			Status: arrange.Status,
			Type:   arrange.Type,
		}
		data[i] = tmp
	}

	return data, err
}

// CancelTime 取消时间
func CancelTime(par models.ParamCancelArrangeTime) (err error) {
	// 删除面试安排记录表记录
	// 修改学生面试时间表
	// 修改学生状态
	tx := mysql.DB.Begin()
	err = mysql.DeleteStudentArrangesByTx(tx, par.ArrangeID, par.Ids)
	if err != nil {
		tx.Rollback()
		return err
	}
	if par.Type == 1 {
		err = mysql.DeleteArrangeVisit(tx, par.Ids)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = mysql.CancelTimeVisit(tx, par.Ids)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else if par.Type == 2 {
		err = mysql.DeleteArrangeInterview(tx, par.Ids)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = mysql.CancelTimeInterview(tx, par.Ids)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = mysql.StudentStatusSub(tx, par.Ids)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

// GetAllArrangeGroup 获取安排组
func GetAllArrangeGroup() (arranges []*models.Arrange, err error) {
	return mysql.GetAllArrangeGroup()
}

// AddArrangeGroup 添加安排组
func AddArrangeGroup(par *models.ParamArrangeGroup) ([]*models.Arrange, error) {
	// 安排顺序：按照报名顺序依次插入到不同的地点
	// !!!!!!!!!!!!!!!!!!!这里每年招新流程不一样，让我很难搞!!!!!!!!!!!!!!!!!!!!!

	// 1、首先计算总共安排人数 = （总时长/单场间隔 (向下取整)）* 场次
	num := int(par.EndTime.Sub(par.StartTime)) / (par.OnceTime * 60000000000) * len(par.Place)
	// 学生面试时间
	t := par.StartTime

	// 2、按照报名顺序依次插入到不同的地点
	arranges := make([]*models.Arrange, len(par.Place))

	// 创建面试安排 	展示：	 面试安排{
	//								student{
	//										TimeArrange
	//										}
	//								}
	// 开启事务
	tx := mysql.DB.Begin()

	for i := 0; i < len(arranges); i++ {
		arranges[i] = &models.Arrange{
			Type:  par.Type,
			Place: par.Place[i],
			Name:  par.Name,
		}
	}

	// 安排宣讲
	n := 0

	var timeArranges []*models.TimeArrange

	ids := []uint{}

	// 安排宣讲
	if par.Type == "visit" {
		// 获取人员列表
		students, err := mysql.GetVisitStudents()
		if len(students) == 0 {
			return nil, tx.Rollback().Error
		}
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		for i, s := range students {
			if i >= num {
				break
			}
			if n >= len(par.Place) {
				n %= len(par.Place)
				t = t.Add(time.Duration(par.OnceTime) * time.Minute)
			}
			arranges[n].Students = append(arranges[n].Students, s)
			ta := &models.TimeArrange{
				StudentID: s.ID,
				Visit:     t,
			}
			mysql.UpdateArrangeTime(tx, ta)
			timeArranges = append(timeArranges, ta)
			ids = append(ids, s.ID)
			n++
		}
	} else {
		// 获取人员列表
		students, err := mysql.GetInterviewStudents(par.NeedVisit)
		if len(students) == 0 {
			return nil, tx.Rollback().Error
		}
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		for i, s := range students {
			if i >= num {
				break
			}
			if n >= len(par.Place) {
				n %= len(par.Place)
				t = t.Add(time.Duration(par.OnceTime) * time.Minute)
			}
			arranges[n].Students = append(arranges[n].Students, s)
			ta := &models.TimeArrange{
				StudentID: s.ID,
				Interview: t,
			}
			mysql.UpdateArrangeTime(tx, ta)
			timeArranges = append(timeArranges, ta)
			ids = append(ids, s.ID)
			n++
		}
	}
	// 插入
	for _, arrange := range arranges {
		if err := mysql.AddArrange(tx, arrange); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// 将timeArranges 与 student关联
	for _, arrange := range arranges {
		for i, s := range arrange.Students {
			s.TimeArrange = timeArranges[i]
		}
	}
	// 学生状态+1
	err := mysql.StudentStatusAdd(tx, ids)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 如果是面试初始化面试表
	if par.Type == "interview" {
		for _, arrange := range arranges {
			ids := make([]uint, len(arrange.Students))
			for i, student := range arrange.Students {
				ids[i] = student.ID
			}
			err := InterviewRecordInit(tx, arrange.ID, ids)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	return arranges, tx.Commit().Error
}

// UpdateInterviewTime 修改面试时间
func UpdateInterviewTime(par *models.ParamArrange) (err error) {
	var ta = &models.TimeArrange{}
	// 开启事务
	tx := mysql.DB.Begin()
	// 宣讲
	for _, v := range par.StudentsId {
		ta.StudentID = v
		ta.Interview = par.Time
		if par.Time.IsZero() {
			err = mysql.UpdateInterviewTimeIsNil(tx, ta)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		err = mysql.UpdateArrangeTime(tx, ta)
		if err != nil {
			// 回滚事务
			tx.Rollback()
			break
		}
	}
	// 提交事务
	tx.Commit()
	return err
}

// UpdateVisitTime 修改宣讲时间
func UpdateVisitTime(par *models.ParamArrange) (err error) {
	var ta = &models.TimeArrange{}
	// 开启事务
	tx := mysql.DB.Begin()
	// 宣讲
	for _, v := range par.StudentsId {
		ta.StudentID = v
		ta.Visit = par.Time
		// 修改时间为默认值
		if par.Time.IsZero() {
			if err = mysql.UpdateVisitTimeIsNil(tx, ta); err != nil {
				tx.Rollback()
				return err
			}
			continue
		}
		err = mysql.UpdateArrangeTime(tx, ta)
		if err != nil {
			// 回滚事务
			tx.Rollback()
			return err
		}
	}
	// 提交事务
	tx.Commit()
	return err
}
