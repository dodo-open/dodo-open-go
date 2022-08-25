package client

import "fmt"

type uri string

const (
	// 机器人 API
	getBotInfoUri        uri = "/api/v1/bot/info"         // 取机器人信息
	setBotIslandLeaveUri uri = "/api/v1/bot/island/leave" // 置机器人群退出

	// 群 API
	getIslandListUri          uri = "/api/v1/island/list"            // 取群列表
	getIslandInfoUri          uri = "/api/v1/island/info"            // 取群信息
	getIslandLevelRankListUri uri = "/api/v1/island/level/rank/list" // 获取群等级排行榜
	getIslandMuteListUri      uri = "/api/v1/island/mute/list"       // 获取群禁言名单
	getIslandBanListUri       uri = "/api/v1/island/ban/list"        // 获取群封禁名单

	// 频道 API
	getChannelListUri uri = "/api/v1/channel/list"   // 取频道列表
	getChannelInfoUri uri = "/api/v1/channel/info"   // 取频道信息
	createChannelUri  uri = "/api/v1/channel/add"    // 创建频道
	editChannelUri    uri = "/api/v1/channel/edit"   // 编辑频道
	removeChannelUri  uri = "/api/v1/channel/remove" // 删除频道

	// 文字频道 API
	sendChannelMessageUri     uri = "/api/v1/channel/message/send"            // 发送频道消息
	editChannelMessageUri     uri = "/api/v1/channel/message/edit"            // 发送频道消息
	withdrawChannelMessageUri uri = "/api/v1/channel/message/withdraw"        // 撤回频道消息
	addChannelMessageReaction uri = "/api/v1/channel/message/reaction/add"    // 新增文字频道消息反应
	remChannelMessageReaction uri = "/api/v1/channel/message/reaction/remove" // 移除文字频道消息反应

	// 身份组 API
	getRoleListUri      uri = "/api/v1/role/list"          // 取身份组列表
	addRoleMemberUri    uri = "/api/v1/role/member/add"    // 身份组新增成员
	removeRoleMemberUri uri = "/api/v1/role/member/remove" // 身份组移除成员

	// 成员 API
	getMemberListUri         uri = "/api/v1/member/list"             // 取成员列表
	getMemberInfoUri         uri = "/api/v1/member/info"             // 取成员信息
	getMemberRoleListUri     uri = "/api/v1/member/role/list"        // 取成员身份组列表
	setMemberNickUri         uri = "/api/v1/member/nick/set"         // 设置成员昵称
	setMemberSilenceUri      uri = "/api/v1/member/ban/set"          // 设置成员禁言
	getMemberInviteInfo      uri = "/api/v1/member/invitation/info"  // 取成员邀请信息
	getMemberUPowerchainInfo uri = "/api/v1/member/upowerchain/info" // 取成员高能链数字藏品信息

	// 私信 API
	sendDirectMessageUri uri = "/api/v1/personal/message/send" // 发送私聊消息

	// 资源 API
	uploadImageUri uri = "/api/v1/resource/picture/upload" // 上传图片资源

	// 事件 API
	getWebsocketConnectionUri uri = "/api/v1/websocket/connection" // 获取 Websocket 连接
)

// getApi build the full api url
func (c *client) getApi(u uri) string {
	return fmt.Sprintf("%s%s", c.conf.BaseApi, u)
}
