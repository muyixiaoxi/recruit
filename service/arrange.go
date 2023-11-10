package service

import (
	"recruit/dao/mysql"
	"recruit/models"
	"time"
)

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
	mysql.TX = mysql.DB.Begin()

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

	// 安排宣讲
	if par.Type == "visit" {
		// 获取人员列表
		students, err := mysql.GetVisitStudents()
		if len(students) == 0 {
			return nil, mysql.TX.Rollback().Error
		}
		if err != nil {
			mysql.TX.Rollback()
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
			mysql.UpdateArrangeTime(ta)
			timeArranges = append(timeArranges, ta)
			n++
		}
	} else {
		// 获取人员列表
		// 如果需要不需要宣讲
		students, err := mysql.GetInterviewStudents(par.NeedVisit)
		if len(students) == 0 {
			return nil, mysql.TX.Rollback().Error
		}
		if err != nil {
			mysql.TX.Rollback()
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
			mysql.UpdateArrangeTime(ta)
			timeArranges = append(timeArranges, ta)
			n++
		}
	}
	// 插入
	for _, arrange := range arranges {
		if err := mysql.AddArrange(arrange); err != nil {
			mysql.TX.Rollback()
			return nil, err
		}
	}
	// 将timeArranges 与 student关联
	for _, arrange := range arranges {
		for i, s := range arrange.Students {
			s.TimeArrange = timeArranges[i]
		}
	}
	return arranges, mysql.TX.Commit().Error
}

// UpdateInterviewTime 修改面试时间
func UpdateInterviewTime(par *models.ParamArrange) (err error) {
	var ta = &models.TimeArrange{}
	// 开启事务
	mysql.TX = mysql.DB.Begin()
	// 宣讲
	for _, v := range par.StudentsId {
		ta.StudentID = v
		ta.Interview = par.Time
		if par.Time.IsZero() {
			err = mysql.UpdateInterviewTimeIsNil(ta)
			if err != nil {
				mysql.TX.Rollback()
				return err
			}
		}
		err = mysql.UpdateArrangeTime(ta)
		if err != nil {
			// 回滚事务
			mysql.TX.Rollback()
			break
		}
	}
	// 提交事务
	mysql.TX.Commit()
	return err
}

// UpdateVisitTime 修改宣讲时间
func UpdateVisitTime(par *models.ParamArrange) (err error) {
	var ta = &models.TimeArrange{}
	// 开启事务
	mysql.TX = mysql.DB.Begin()
	// 宣讲
	for _, v := range par.StudentsId {
		ta.StudentID = v
		ta.Visit = par.Time
		// 修改时间为默认值
		if par.Time.IsZero() {
			if err = mysql.UpdateVisitTimeIsNil(ta); err != nil {
				mysql.TX.Rollback()
				return err
			}
			continue
		}
		err = mysql.UpdateArrangeTime(ta)
		if err != nil {
			// 回滚事务
			mysql.TX.Rollback()
			return err
		}
	}
	// 提交事务
	mysql.TX.Commit()
	return err
}
