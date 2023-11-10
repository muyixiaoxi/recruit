package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"recruit/models"
	"recruit/service"
)

// AddReadMsg 添加已读消息
func AddReadMsg(c *gin.Context) {
	openid := c.Query("openid")
	var rm models.ReadMessage
	if err := c.ShouldBind(&rm); err != nil {
		zap.L().Error("err := c.ShouldBind(&rm) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	err := service.AddReadMsg(&rm, openid)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// GetMyAllMsg 获取我收到的全部消息
func GetMyAllMsg(c *gin.Context) {
	openid := c.Query("openid")
	msg, err := service.GetSignUpInfoByOpenid(openid)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}
