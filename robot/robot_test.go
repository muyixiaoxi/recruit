package robot

import (
	"testing"
)

func TestSendMsg(t *testing.T) {
	Init()
	// 获取所有频道
	guild, _ := GetGuilds()
	// 创建子频道
	//channel, _ := AddChannel(guild[0])
	channels, _ := GetChannelList(guild[0])
	// 发送消息

	CreateGuildAnnounces(guild[0], channels[0])
}

func TestSendMsg2(t *testing.T) {

}
