package client

import "fmt"

type uri string

const (
	getBotInfoUri         uri = "/api/v1/bot/info"             // 取机器人信息
	setBotIslandLeaveUri  uri = "/api/v1/bot/island/leave"     // 置机器人群退出
	getIslandListUri      uri = "/api/v1/island/list"          // 取群列表
	getIslandInfoUri      uri = "/api/v1/island/info"          // 取群信息
	getChannelListUri     uri = "/api/v1/channel/list"         // 取频道列表
	getChannelInfoUri     uri = "/api/v1/channel/info"         // 取频道信息
	sendChannelMessageUri uri = "/api/v1/channel/message/send" // 发送频道消息
)

// getApi build the full api url
func (c *client) getApi(u uri) string {
	return fmt.Sprintf("%s%s", c.conf.BaseApi, u)
}
