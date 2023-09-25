package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"recruit/models"
	"recruit/service"
)

// GetUnreadMsg 获取未读消息
func GetUnreadMsg(c *gin.Context) {
	openid := c.Query("openid")
	msg, err := service.GetUnreadMsg(openid)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// SignUp 报名
func SignUp(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBind(&student); err != nil {
		zap.L().Error("c.ShouldBind(&student) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := service.SignUp(&student); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// Login 登录
func Login(c *gin.Context) {
	code := c.Query("code")
	// 获取openid
	response, err := http.Get(fmt.Sprintf("https://api.q.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, secret, code))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to QQ API"})
		return
	}
	defer response.Body.Close()
	var loginResponse models.LoginResponse
	if err := json.NewDecoder(response.Body).Decode(&loginResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse QQ API response"})
		return
	}

	// 注册
	service.Signup(loginResponse.Openid)

	c.JSON(http.StatusOK, loginResponse)
}

// GetMySignUp 报名进程
func GetMySignUp(c *gin.Context) {
	//登录，返回个人进程
	openid := c.Query("openid")
	date, err := service.GetMySignUp(openid)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, date)
}
