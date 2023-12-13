package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"recruit/models"
	"recruit/service"
	"strconv"
)

// GetDetailMsg 获取详细通知
func GetDetailMsg(c *gin.Context) {
	var sId = c.Query("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		zap.L().Error("strconv.Atoi(sId) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	msg := models.Message{Model: gorm.Model{ID: uint(id)}}
	err = service.GetDetailMsg(&msg)
	if err != nil {
		zap.L().Error("service.GetDetailMsg(&msg) failed", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// DeleteMsg 删除通知
func DeleteMsg(c *gin.Context) {
	var draft models.ParamIds
	if err := c.ShouldBind(&draft); err != nil {
		zap.L().Error("c.ShouldBind(&draft) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := service.DeleteMsg(&draft); err != nil {
		zap.L().Error("service.UpdateMsg(&draft) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeInvalidPassword)
}

// UpdateDraft 修改草稿箱
func UpdateDraft(c *gin.Context) {
	var draft models.Message
	if err := c.ShouldBind(&draft); err != nil {
		zap.L().Error("c.ShouldBind(&draft) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := service.UpdateDraft(&draft); err != nil {
		zap.L().Error("service.UpdateDraft(&draft) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeInvalidPassword)
}

// GetAllDraft 获取所有草稿
func GetAllDraft(c *gin.Context) {
	data, err := service.GetAllDraft()
	if err != nil {
		zap.L().Error("service.GetAllDraft() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, data)
}

// AddDraft 保存到草稿箱
func AddDraft(c *gin.Context) {
	var draft models.Message
	if err := c.ShouldBind(&draft); err != nil {
		zap.L().Error("c.ShouldBind(&draft) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// state = 0 为草稿
	draft.State = 0
	err := service.AddDraft(&draft)
	if err != nil {
		zap.L().Error("service.AddDraft() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// UserGetAllMsg 获取所有消息
func UserGetAllMsg(c *gin.Context) {
	data, err := service.UserGetAllMsg()
	if err != nil {
		zap.L().Error("service.GetAllMsg failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// SendMsg 发送自定义通知
func SendMsg(c *gin.Context) {
	var psm = models.ParamSendMsg{}
	if err := c.ShouldBind(&psm); err != nil {
		zap.L().Error("err := c.ShouldBind(&msg) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 将数组处理成字符串
	ids := ""
	for i, v := range psm.StudentsId {
		if i != 0 {
			ids += "," + strconv.Itoa(v)
		} else {
			ids += strconv.Itoa(v)
		}
	}
	var msg = models.Message{
		Title:      psm.Title,
		Content:    psm.Content,
		StudentsID: ids,
		State:      1,
	}
	if err := service.SendMsg(&msg); err != nil {
		zap.L().Error("err := service.SendMsg(&msg) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
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
