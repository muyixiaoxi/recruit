package robot

import (
	"context"
	"fmt"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"log"
	"time"
)

var (
	appid  = uint64(102069444)
	aToken = "pUUfGUZNM8Vs0cp9DQ5ojAFlIYnC8nQP"
)

var (
	api   openapi.OpenAPI
	ctx   context.Context
	robot *dto.User
)

func Init() (err error) {
	token := token.BotToken(appid, aToken)
	api = botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
	ctx = context.Background()

	robot, err = api.Me(ctx)
	if err != nil {
		log.Fatalln("调用 Me 接口失败, err = ", err)
	}
	fmt.Println("robotId:", robot.ID)
	return err
}

// GetGuilds 获取所有频道
func GetGuilds() (guilds []*dto.Guild, err error) {
	guilds, err = api.MeGuilds(ctx, &dto.GuildPager{})
	if err != nil {
		log.Fatalln("调用 MeGuild 接口失败, err = ", err)
	}
	fmt.Println(guilds)
	return guilds, err
}

// GetGuild 获取频道详细信息
func GetGuild(g *dto.Guild) (*dto.Guild, error) {
	guild, guildError := api.Guild(ctx, g.ID)
	if guildError != nil {
		log.Fatalln("调用 Guild 接口失败, err = ", guildError)
	}
	return guild, guildError
}

// GetGuildUserList 获取频道成员列表
func GetGuildUserList(g *dto.Guild) ([]*dto.Member, error) {
	members, err := api.GuildMembers(ctx, g.ID, &dto.GuildMembersPager{
		After: "0",
		Limit: "100",
	})
	if err != nil {
		log.Fatalln("调用 GuildMembers 接口失败, err = ", err)
	}
	return members, err
}

// GetChannelList 获取子频道列表
func GetChannelList(g *dto.Guild) ([]*dto.Channel, error) {
	channels, channelsError := api.Channels(ctx, g.ID)
	if channelsError != nil {
		log.Fatalln("调用 Channels, err = ", channelsError)
	}
	return channels, channelsError
}

// AddChannel 创建子频道
func AddChannel(g *dto.Guild) (*dto.Channel, error) {
	channel, err := api.PostChannel(ctx, g.ID, &dto.ChannelValueObject{
		Name:    "测试子频道名",
		Type:    dto.ChannelTypeText,
		SubType: dto.ChannelSubTypeChat,
	})
	if err != nil {
		log.Fatalln("调用 PostChannel 接口失败, err = ", err)
	}
	return channel, err
}

// SetChannelLimit 设置子频道发言权限
func SetChannelLimit(c *dto.Channel) {
	fmt.Println("robotId:", robot.ID)
	err := api.PutChannelPermissions(ctx, c.ID, robot.ID, &dto.UpdateChannelPermissions{
		Add: "4",
	})
	if err != nil {
		log.Fatalln("调用 PutChannelPermissions 接口失败, err = ", err)
	}

}

// SendMsg 发送消息
func SendMsg(c *dto.Channel) {

	_, err := api.PostMessage(ctx, c.ID, &dto.MessageToCreate{
		MsgID:   "", //如果未空则表示主动消息
		Content: "hello world",
	})
	if err != nil {
		log.Fatalln("调用 PostMessage 接口失败, err = ", err)
	}
}

func CreateGuildAnnounces(g *dto.Guild, channel *dto.Channel) {

	_, err := api.CreateGuildAnnounces(ctx, g.ID, &dto.GuildAnnouncesToCreate{
		MessageID: "测试消息id",
		ChannelID: channel.ID,
	})
	if err != nil {
		log.Fatalln("调用 CreateGuildAnnounces 接口失败, err = ", err)
	}

}
