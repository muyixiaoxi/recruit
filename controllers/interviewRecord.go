package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"recruit/models"
	"recruit/service"
	"strconv"
	"sync"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 互斥锁
var wsMutex sync.Mutex

// InterviewRecord 面试记录
func InterviewRecord(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Error("upgrader.Upgrade(c.Writer,c.Request,nil) failed", zap.Error(err))
		return
	}
	uid, _ := getCurrentUser(c)
	userChan := make(chan models.InterviewRecord, 100)
	service.OnlineUser[uid] = userChan

	defer func() {
		// 关闭WebSocket连接
		ws.Close()
		// 删除用户通道
		close(userChan)
		delete(service.OnlineUser, uid)
	}()

	// 处理WebSocket消息
	go SendServer(ws)
	go ReceiveServer(ws, userChan)
	select {} // 阻塞函数
}

// SendServer 向服务端发送
func SendServer(ws *websocket.Conn) {
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			zap.L().Error("ws.ReadMessage() failed", zap.Error(err))
			break
		}

		// 解析JSON
		var per models.InterviewRecord
		if err := json.Unmarshal(p, &per); err != nil {
			ws.WriteJSON(gin.H{
				"msg": "json解析错误",
			})
			continue
		}
		service.SendServer(&per)
		err = ws.WriteJSON(per)
		if err != nil {
			zap.L().Error("ws.WriteMessage(messageType,p) failed", zap.Error(err))
			break
		}
	}
}

// ReceiveServer 接收服务端消息
func ReceiveServer(ws *websocket.Conn, userChan chan models.InterviewRecord) {
	for {
		wsMutex.Lock()
		err := ws.WriteJSON(<-userChan)
		wsMutex.Unlock()
		fmt.Println("ReceiveServer接收")
		if err != nil {
			zap.L().Error("ws.WriteMessage(messageType,p) failed", zap.Error(err))
			break
		}
	}
}

// GetInterviewRecord 获取面试记录
func GetInterviewRecord(c *gin.Context) {
	par := c.Query("arrange_id")
	id, _ := strconv.Atoi(par)
	data, err := service.GetInterviewRecord(uint(id))
	if err != nil {
		zap.L().Error("service.GetAllInterviewRecord() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// AddContentId 添加列
func AddContentId(c *gin.Context) {
	par := c.Query("id")
	id, _ := strconv.Atoi(par)
	if err := service.AddContentId(uint(id)); err != nil {
		zap.L().Error("service.AddContentId(uint(id)) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// DeleteContentId 删除列
func DeleteContentId(c *gin.Context) {
	para := c.Query("arrange_id")
	parc := c.Query("content_id")
	arrange, _ := strconv.Atoi(para)
	content, _ := strconv.Atoi(parc)
	err := service.DeleteContentId(arrange, content)
	if err != nil {
		zap.L().Error("service.DeleteContentId(arrange,content) failed", zap.Error(err))
		ResponseError(c, CodeSuccess)
	}
	ResponseSuccess(c, nil)
}
