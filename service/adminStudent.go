package service

import (
	"recruit/dao/mysql"
	"recruit/models"
)

// UpdateInterviewState 同步学生面试状况
func UpdateInterviewState() error {
	return mysql.UpdateInterviewState()
}

// UpdateVisitState 同步学生宣讲状况
func UpdateVisitState() error {
	return mysql.UpdateVisitState()
}

// 查询的专业
var specialtys = map[string]string{"计算机科学": "计%科", "数据科学与大数据技术": "数%据", "物联网工程": "物%联", "信息工程": "信%工", "通信工程": "通%信"}

//var specialtys = []string{"计%科", "数%据", "物%联", "信%工", "通%信"}

// GroupBySpecialty 根据专业分组
func GroupBySpecialty() (data []*models.ParamGroupBySpecialty, err error) {
	// 获取所有报名学生
	students, err := mysql.GetAllSignUpAndTimeArrange()
	if err != nil {
		return nil, err
	}
	// 统计已面试，已通过
	inter, err := mysql.GetInterviewedStudent()
	if err != nil {
		return nil, err
	}
	// 已录取
	enroll, err := mysql.GetEnrollSuccess()
	if err != nil {
		return nil, err
	}
	// 男生
	boys, err := mysql.GetSignUpByGender("男")
	if err != nil {
		return nil, err
	}
	// 头部消息
	head := &models.ParamGroupBySpecialty{
		Type:           "总",
		Sum:            len(students),
		InterviewedNum: len(inter),
		EnrollNum:      len(enroll),
		BoyNum:         len(boys),
		GirlNum:        len(students) - len(boys),
	}
	data = append(data, head)

	// 查取不同专业
	var perBoys, perGirls, perEnrolls, perSum int
	for i, specialty := range specialtys {
		// 总数
		sum, err := mysql.GetSignUpBySpecialty(specialty)

		if err != nil {
			return nil, err
		}
		// 录取人数
		enroll, err := mysql.GetEnrollSuccessBySpecialty(specialty)
		if err != nil {
			return nil, err
		}
		// 男生
		boy, err := mysql.GetSignUpBySpecialtyAndGender(specialty, "男")
		if err != nil {
			return nil, err
		}
		per := &models.ParamGroupBySpecialty{
			Type:      i,
			Sum:       len(sum),
			BoyNum:    len(boy),
			GirlNum:   len(sum) - len(boy),
			EnrollNum: len(enroll),
		}
		perBoys += per.BoyNum
		perGirls += per.GirlNum
		perEnrolls += per.EnrollNum
		perSum += per.Sum
		if per.Sum != 0 {
			data = append(data, per)
		}
	}

	// 其他
	other := &models.ParamGroupBySpecialty{
		Type:      "其他",
		Sum:       head.Sum - perSum,
		BoyNum:    head.BoyNum - perBoys,
		GirlNum:   head.GirlNum - perGirls,
		EnrollNum: head.EnrollNum - perEnrolls,
	}
	if other.Sum != 0 {
		data = append(data, other)
	}
	return
}
