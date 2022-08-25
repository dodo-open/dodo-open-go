package client

import (
	"context"
	"errors"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/go-resty/resty/v2"
	"time"
)

type (
	// Client DoDoBot API Interface
	Client interface {
		Base
		IslandAPI
		ChannelAPI
		MessageAPI
		RoleAPI
		MemberAPI
		DirectMessageAPI
		ResourceUploadAPI
		WebsocketAPI
	}

	// Base client basic interface
	Base interface {
		GetConfig() *Config
		GetBotInfo(ctx context.Context) (*model.GetBotInfoRsp, error)                         // GetBotInfo 取机器人信息
		SetBotIslandLeave(ctx context.Context, req *model.SetBotLeaveIslandReq) (bool, error) // SetBotIslandLeave 置机器人群退出
	}

	// IslandAPI island API interface
	IslandAPI interface {
		GetIslandList(ctx context.Context) ([]*model.IslandElement, error)                                                            // GetIslandList 取群列表
		GetIslandInfo(ctx context.Context, req *model.GetIslandInfoReq) (*model.GetIslandInfoRsp, error)                              // GetIslandInfo 取群信息
		GetIslandLevelRankList(ctx context.Context, req *model.GetIslandLevelRankListReq) ([]*model.GetIslandLevelRankElement, error) // GetIslandLevelRankList 获取群等级排行榜
		GetIslandMuteList(ctx context.Context, req *model.GetIslandMuteListReq) (*model.GetIslandMuteListRsp, error)                  // GetIslandMuteList 获取群禁言名单
		GetIslandBanList(ctx context.Context, req *model.GetIslandBanListReq) (*model.GetIslandBanListRsp, error)                     // GetIslandBanList 获取群封禁名单
	}

	// ChannelAPI channel basic API interface
	ChannelAPI interface {
		GetChannelList(ctx context.Context, req *model.GetChannelListReq) ([]*model.ChannelElement, error)  // GetChannelList 取频道列表
		GetChannelInfo(ctx context.Context, req *model.GetChannelInfoReq) (*model.GetChannelInfoRsp, error) // GetChannelInfo 取频道信息
		CreateChannel(ctx context.Context, req *model.CreateChannelReq) (*model.CreateChannelRsp, error)    // CreateChannel 创建频道
		EditChannel(ctx context.Context, req *model.EditChannelReq) (bool, error)                           // EditChannel 编辑频道
		RemoveChannel(ctx context.Context, req *model.RemoveChannelReq) (bool, error)                       // RemoveChannel 编辑频道
	}

	// MessageAPI message API interface
	MessageAPI interface {
		SendChannelMessage(ctx context.Context, req *model.SendChannelMessageReq) (*model.SendChannelMessageRsp, error) // SetChannelMessageSend 发送频道消息
		EditChannelMessage(ctx context.Context, req *model.EditChannelMessageReq) (*model.EditChannelMessageRsp, error) // SetChannelMessageEdit 编辑频道消息
		WithdrawChannelMessage(ctx context.Context, req *model.WithdrawChannelMessageReq) (bool, error)                 // SetChannelMessageWithdraw 撤回频道消息
		AddChannelMessageReaction(ctx context.Context, req *model.AddChannelMessageReactionReq) (bool, error)           // SetChannelMessageReactionAdd 添加频道消息反应
		RemChannelMessageReaction(ctx context.Context, req *model.RemChannelMessageReactionReq) (bool, error)           // SetChannelMessageReactionRemove 移除文字频道消息反应
	}

	// RoleAPI role API interface
	RoleAPI interface {
		GetRoleList(ctx context.Context, req *model.GetRoleListReq) ([]*model.RoleElement, error) // GetRoleList 取身份组列表
		CreateRole(ctx context.Context, req *model.CreateRoleReq) (*model.CreateRoleRsp, error)   // CreateRole 创建身份组
		EditRole(ctx context.Context, req *model.EditRoleReq) (bool, error)                       // EditRole 编辑身份组
		RemoveRole(ctx context.Context, req *model.RemoveRoleReq) (bool, error)                   // RemoveRole 删除身份组
		AddRoleMember(ctx context.Context, req *model.AddRoleMemberReq) (bool, error)             // SetRoleMemberAdd 身份组新增成员
		RemoveRoleMember(ctx context.Context, req *model.RemoveRoleMemberReq) (bool, error)       // SetRoleMemberRemove 身份组移除成员
	}

	// MemberAPI member API interface
	MemberAPI interface {
		GetMemberList(ctx context.Context, req *model.GetMemberListReq) (*model.GetMemberListRsp, error)                                  // GetMemberList 取成员列表
		GetMemberInfo(ctx context.Context, req *model.GetMemberInfoReq) (*model.GetMemberInfoRsp, error)                                  // GetMemberInfo 取成员信息
		GetMemberRoleList(ctx context.Context, req *model.GetMemberRoleListReq) ([]*model.RoleElement, error)                             // GetMemberRoleList 取成员身份组列表
		SetMemberNick(ctx context.Context, req *model.SetMemberNickReq) (bool, error)                                                     // SetMemberNick 设置成员昵称
		SetMemberSilence(ctx context.Context, req *model.SetMemberSilenceReq) (bool, error)                                               // SetMemberBan 设置成员禁言，即不能在频道发布内容
		GetMemberInviteInfo(ctx context.Context, req *model.GetMemberInviteInfoReq) (*model.GetMemberInviteInfoRsp, error)                // GetMemberInvitationInfo 取成员邀请信息
		GetMemberUPowerchainInfo(ctx context.Context, req *model.GetMemberUPowerchainInfoReq) (*model.GetMemberUPowerchainInfoRsp, error) // GetMemberUPowerchainInfo 取成员高能链数字藏品信息
	}

	// DirectMessageAPI direct message (a.k.a. DM) API interface
	DirectMessageAPI interface {
		SendDirectMessage(ctx context.Context, req *model.SendDirectMessageReq) (*model.SendDirectMessageRsp, error)
	}

	// ResourceUploadAPI resource upload API interface
	ResourceUploadAPI interface {
		UploadImageByBytes(ctx context.Context, req *model.UploadImageByBytesReq) (*model.UploadImageRsp, error)
		UploadImageByPath(ctx context.Context, req *model.UploadImageByPathReq) (*model.UploadImageRsp, error)
	}

	// WebsocketAPI websocket API interface
	WebsocketAPI interface {
		GetWebsocketConnection(ctx context.Context) (*model.GetWebsocketConnectionRsp, error)
	}
)

type (
	// client DoDoBot Instance
	client struct {
		conf *Config
		r    *resty.Client // resty client
	}

	// Config DoDoBot client configuration
	Config struct {
		BaseApi  string        // DoDo OpenAPI Server Domain
		ClientId string        // DoDoBot ClientID
		Token    string        // DoDoBot Bot token
		IsDebug  bool          // debug mode
		Timeout  time.Duration // resty client request timeout
	}
)

// New create a new DoDoBot instance
func New(clientId, token string, options ...OptionHandler) (Client, error) {
	config := getDefaultConfig()
	config.ClientId = clientId
	config.Token = token

	// handle custom options
	for _, optionHandler := range options {
		if optionHandler == nil {
			return nil, errors.New("invalid OptionHandler (nil detected)")
		}
		if err := optionHandler(config); err != nil {
			return nil, err
		}
	}

	instance := &client{conf: config}
	instance.setupResty()

	return instance, nil
}

// getDefaultConfig Get the default configuration
func getDefaultConfig() *Config {
	return &Config{
		BaseApi: "https://botopen.imdodo.com",
		IsDebug: false,
		Timeout: time.Second * 5,
	}
}

// GetConfig get instance configuration
func (c *client) GetConfig() *Config {
	return c.conf
}
