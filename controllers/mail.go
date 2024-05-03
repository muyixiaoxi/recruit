package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"recruit/models"
	"recruit/service"
)

func SendMail(c *gin.Context) {
	req := &models.SendMailRequest{}
	if err := c.ShouldBind(req); err != nil {
		zap.L().Error("SendMail c.ShouldBind(req) fail:", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	failIds, _ := service.SendMail(*req)

	ResponseSuccess(c, models.SendMailResponse{SendFailIds: failIds})
}
