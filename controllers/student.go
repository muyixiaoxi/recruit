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
	var par models.ParamLogin
	if err := c.ShouldBind(&par); err != nil {
		zap.L().Error("c.ShouldBind(&par) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 获取openid
	response, err := http.Get(fmt.Sprintf("https://api.q.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, secret, par.Code))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to QQ API"})
		return
	}
	defer response.Body.Close()
	var loginResponse models.LoginResponse
	fmt.Println(response.Body)
	if err := json.NewDecoder(response.Body).Decode(&loginResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse QQ API response"})
		return
	}

	// 登录
	service.StudentLogin(loginResponse.Openid, par.Avatar)

	c.JSON(http.StatusOK, loginResponse)
}

// GetMySignUp 报名进程
func GetMySignUp(c *gin.Context) {
	//登录，返回个人进程
	openid := c.Query("openid")
	date, err := service.GetSignUpInfoByOpenid(openid)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, date)
}
