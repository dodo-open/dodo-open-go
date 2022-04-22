package client

import "fmt"

type uri string

const (
	getBotInfoUri        uri = "/api/v1/bot/info"         // 取机器人信息
	setBotIslandLeaveUri uri = "/api/v1/bot/island/leave" // 置机器人群退出

	getIslandListUri uri = "/api/v1/island/list" // 取群列表
	getIslandInfoUri uri = "/api/v1/island/info" // 取群信息

	getChannelListUri uri = "/api/v1/channel/list" // 取频道列表
	getChannelInfoUri uri = "/api/v1/channel/info" // 取频道信息

	sendChannelMessageUri     uri = "/api/v1/channel/message/send"     // 发送频道消息
	editChannelMessageUri     uri = "/api/v1/channel/message/edit"     // 发送频道消息
	withdrawChannelMessageUri uri = "/api/v1/channel/message/withdraw" // 撤回频道消息

	getRoleListUri      uri = "/api/v1/role/list"          // 取身份组列表
	addRoleMemberUri    uri = "/api/v1/role/member/add"    // 身份组新增成员
	removeRoleMemberUri uri = "/api/v1/role/member/remove" // 身份组移除成员

	getMemberListUri     uri = "/api/v1/member/list"      // 取成员列表
	getMemberInfoUri     uri = "/api/v1/member/info"      // 取成员信息
	getMemberRoleListUri uri = "/api/v1/member/role/list" // 取成员身份组列表
	setMemberNickUri     uri = "/api/v1/member/nick/set"  // 设置成员昵称
	setMemberSilenceUri  uri = "/api/v1/member/ban/set"   // 设置成员禁言

	sendDirectMessageUri uri = "/api/v1/personal/message/send" // 发送私聊消息

	uploadImageUri uri = "/api/v1/resource/picture/upload" // 上传图片资源

	getWebsocketConnectionUri uri = "/api/v1/websocket/connection" // 获取 Websocket 连接
)

// getApi build the full api url
func (c *client) getApi(u uri) string {
	return fmt.Sprintf("%s%s", c.conf.BaseApi, u)
}
