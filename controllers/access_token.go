package controllers

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"recruit/models"
	"time"
)

var (
	appId      = "1112250789"
	secret     = "e8DGwrNUDIxID6TP"
	tickerTime = 110 * time.Minute
)

var accessToken models.AccessToken

func TickerSetAccessToken() {
	ticker := time.NewTicker(tickerTime)
	for range ticker.C {
		SetAccessToken()
	}
}

// SetAccessToken 设置AccessToken
func SetAccessToken() {
	url := fmt.Sprintf("https://api.q.qq.com/api/getToken?grant_type=client_credential&appid=%s&secret=%s", appId, secret)
	resp, err := http.Get(url)
	if err != nil {
		zap.L().Error("resp, err := http.Get(url) failed", zap.Error(err))
		return
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&accessToken); err != nil {
		zap.L().Error("err := json.NewDecoder(resp.Body).Decode(&accessToken) failed", zap.Error(err))
		return
	}
}
