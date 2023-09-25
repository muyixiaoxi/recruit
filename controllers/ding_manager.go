package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"recruit/models"
	"recruit/service"
)

// SendMsg 发送自定义消息
func SendMsg(c *gin.Context) {
	var msg = models.Message{}
	if err := c.ShouldBind(&msg); err != nil {
		zap.L().Error("err := c.ShouldBind(&msg) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := service.SendMsg(&msg); err != nil {
		zap.L().Error("err := service.SendMsg(&msg) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// GetAllSignUp 获取所有报名信息
func GetAllSignUp(c *gin.Context) {
	data, err := service.GetAllSignUp()
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

// SendTemplateMessage 通知模板消息
func SendTemplateMessage(c *gin.Context) {
	token := c.Query("access_token")
	url := fmt.Sprintf("https://api.q.qq.com/api/json/template/send?access_token=%s", token)
	var tm = models.TemplateMessage{}
	// 将结构体转换为JSON字符串
	json, err := json.Marshal(tm)
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	defer resp.Body.Close()
}
