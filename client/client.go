package client

import (
	"context"
	"dodo-open-go/model"
	"errors"
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
		GetBotInfo(ctx context.Context) (*model.GetBotInfoRsp, error)
		SetBotIslandLeave(ctx context.Context, req *model.SetBotLeaveIslandReq) (bool, error)
	}

	// IslandAPI island API interface
	IslandAPI interface {
		GetIslandList(ctx context.Context) ([]*model.IslandElement, error)
		GetIslandInfo(ctx context.Context, req *model.GetIslandInfoReq) (*model.GetIslandInfoRsp, error)
	}

	// ChannelAPI channel basic API interface
	ChannelAPI interface {
		GetChannelList(ctx context.Context, req *model.GetChannelListReq) ([]*model.ChannelElement, error)
		GetChannelInfo(ctx context.Context, req *model.GetChannelInfoReq) (*model.GetChannelInfoRsp, error)
	}

	// MessageAPI message API interface
	MessageAPI interface {
		SendChannelMessage(ctx context.Context, req *model.SendChannelMessageReq) (*model.SendChannelMessageRsp, error)
		EditChannelMessage(ctx context.Context, req *model.EditChannelMessageReq) (*model.EditChannelMessageRsp, error)
		WithdrawChannelMessage(ctx context.Context, req *model.WithdrawChannelMessageReq) (bool, error)
	}

	// RoleAPI role API interface
	RoleAPI interface {
		GetRoleList(ctx context.Context, req *model.GetRoleListReq) ([]*model.RoleElement, error)
		AddRoleMember(ctx context.Context, req *model.AddRoleMemberReq) (bool, error)
		RemoveRoleMember(ctx context.Context, req *model.AddRoleMemberReq) (bool, error)
	}

	// MemberAPI member API interface
	MemberAPI interface {
		GetMemberList(ctx context.Context, req *model.GetMemberListReq) (*model.GetMemberListRsp, error)
		GetMemberInfo(ctx context.Context, req *model.GetMemberInfoReq) (*model.GetMemberInfoRsp, error)
		GetMemberRoleList(ctx context.Context, req *model.GetMemberRoleListReq) ([]*model.RoleElement, error)
		SetMemberNick(ctx context.Context, req *model.SetMemberNickReq) (bool, error)
		SetMemberSilence(ctx context.Context, req *model.SetMemberSilenceReq) (bool, error)
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
