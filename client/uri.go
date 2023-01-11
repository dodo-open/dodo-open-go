package client

import "fmt"

type uri string

const (
	// 机器人 API
	getBotInfoUri         uri = "/api/v2/bot/info"          // 获取机器人信息
	setBotIslandLeaveUri  uri = "/api/v2/bot/island/leave"  // 机器人退群
	getBotInviteListUri   uri = "/api/v2/bot/invite/list"   // 获取机器人邀请列表
	setBotInviteAddUri    uri = "/api/v2/bot/invite/add"    // 添加成员到机器人邀请列表
	setBotInviteRemoveUri uri = "/api/v2/bot/invite/remove" // 移除成员出机器人邀请列表

	// 群 API
	getIslandListUri          uri = "/api/v2/island/list"            // 获取群列表
	getIslandInfoUri          uri = "/api/v2/island/info"            // 获取群信息
	getIslandLevelRankListUri uri = "/api/v2/island/level/rank/list" // 获取群等级排行榜
	getIslandMuteListUri      uri = "/api/v2/island/mute/list"       // 获取群禁言名单
	getIslandBanListUri       uri = "/api/v2/island/ban/list"        // 获取群封禁名单

	// 频道 API
	getChannelListUri uri = "/api/v2/channel/list"   // 获取频道列表
	getChannelInfoUri uri = "/api/v2/channel/info"   // 获取频道信息
	createChannelUri  uri = "/api/v2/channel/add"    // 创建频道
	editChannelUri    uri = "/api/v2/channel/edit"   // 编辑频道
	removeChannelUri  uri = "/api/v2/channel/remove" // 删除频道

	// 文字频道 API
	sendChannelMessageUri                  uri = "/api/v2/channel/message/send"                 // 发送消息
	editChannelMessageUri                  uri = "/api/v2/channel/message/edit"                 // 编辑消息
	withdrawChannelMessageUri              uri = "/api/v2/channel/message/withdraw"             // 撤回消息
	addChannelMessageReaction              uri = "/api/v2/channel/message/reaction/add"         // 添加表情反应
	remChannelMessageReaction              uri = "/api/v2/channel/message/reaction/remove"      // 取消表情反应
	setChannelMessageTopUri                uri = "/api/v2/channel/message/top"                  // 置顶消息
	getChannelMessageReactionListUri       uri = "/api/v2/channel/message/reaction/list"        // 获取消息反应列表
	getChannelMessageReactionMemberListUri uri = "/api/v2/channel/message/reaction/member/list" // 获取消息反应内成员列表

	// 语音频道 API
	getChannelVoiceMemberStatusUri uri = "/api/v2/channel/voice/member/status" // 获取成员语音频道状态
	setChannelVoiceMemberMoveUri   uri = "/api/v2/channel/voice/member/move"   // 移动语音频道成员
	setChannelVoiceMemberEditUri   uri = "/api/v2/channel/voice/member/edit"   // 管理语音中的成员

	// 帖子频道
	setChannelArticleAddUri    uri = "/api/v2/channel/article/add"    // 发布帖子
	setChannelArticleRemoveUri uri = "/api/v2/channel/article/remove" // 删除帖子评论回复

	// 身份组 API
	getRoleListUri       uri = "/api/v2/role/list"          // 获取身份组列表
	createRoleUri        uri = "/api/v2/role/add"           // 创建身份组
	editRoleUri          uri = "/api/v2/role/edit"          // 编辑身份组
	removeRoleUri        uri = "/api/v2/role/remove"        // 删除身份组
	addRoleMemberUri     uri = "/api/v2/role/member/add"    // 赋予成员身份组
	removeRoleMemberUri  uri = "/api/v2/role/member/remove" // 取消成员身份组
	getRoleMemberListUri uri = "/api/v2/role/member/list"   // 获取身份组成员列表

	// 成员 API
	getMemberListUri     uri = "/api/v2/member/list"            // 获取成员列表
	getMemberInfoUri     uri = "/api/v2/member/info"            // 获取成员信息
	getMemberRoleListUri uri = "/api/v2/member/role/list"       // 获取成员身份组列表
	getMemberInviteInfo  uri = "/api/v2/member/invitation/info" // 获取成员邀请信息
	setMemberNickUri     uri = "/api/v2/member/nickname/edit"   // 编辑成员群昵称
	muteMemberUri        uri = "/api/v2/member/mute/add"        // 禁言成员
	unmuteMemberUri      uri = "/api/v2/member/mute/remove"     // 取消成员禁言
	banMemberUri         uri = "/api/v2/member/ban/set"         // 永久封禁成员
	unbanMemberUri       uri = "/api/v2/member/ban/remove"      // 取消成员永久封禁

	// 赠礼系统 API
	getGiftAccountUri        uri = "/api/v2/gift/account/info"     // 获取群收入
	getGiftShareRatioInfoUri uri = "/api/v2/gift/share/ratio/info" // 获取成员分成管理
	getGiftListUri           uri = "/api/v2/gift/list"             // 获取内容礼物列表
	getGiftMemberListUri     uri = "/api/v2/gift/member/list"      // 获取内容礼物内成员列表
	getGiftGrossValueListUri uri = "/api/v2/gift/gross/value/list" // 获取内容礼物总值列表

	// 数字藏品 API
	getMemberNFTStatusUri    uri = "/api/v2/member/nft/status"       // 获取成员数字藏品判断
	getMemberUPowerchainInfo uri = "/api/v2/member/upowerchain/info" // 取成员高能链数字藏品信息

	// 私信 API
	sendDirectMessageUri uri = "/api/v2/personal/message/send" // 发送私信

	// 资源 API
	uploadImageUri uri = "/api/v2/resource/picture/upload" // 上传图片资源

	// 事件 API
	getWebsocketConnectionUri uri = "/api/v2/websocket/connection" // 获取 Websocket 连接
)

// getApi build the full api url
func (c *client) getApi(u uri) string {
	return fmt.Sprintf("%s%s", c.conf.BaseApi, u)
}
