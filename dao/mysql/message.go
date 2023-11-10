package mysql

import (
	"go.uber.org/zap"
	"recruit/models"
)

// AddReadMsg 添加已读消息
func AddReadMsg(rm *models.ReadMessage) error {
	res := DB.Create(rm)
	if res.Error != nil {
		zap.L().Error("DB.Create(rm) failed", zap.Error(res.Error))
	}
	return res.Error
}
