package model

import "errors"

type ChannelType int

const (
	TextChannel    ChannelType = 1
	VoiceChannel   ChannelType = 2
	ArticleChannel ChannelType = 4
	LinkChannel    ChannelType = 5
	ResChannel     ChannelType = 6
)

// ChannelElement 获取频道列表 list element
type ChannelElement struct {
	ChannelId   string      `json:"channelId"`   // 频道号
	ChannelName string      `json:"channelName"` // 频道名称
	ChannelType ChannelType `json:"channelType"` // 频道类型，1：文字频道，2：语音频道，4：帖子频道，5：链接频道，6：资料频道
	DefaultFlag int         `json:"defaultFlag"` // 默认频道标识，0：否，1：是
	GroupId     string      `json:"groupId"`     // 分组ID
	GroupName   string      `json:"groupName"`   // 分组名称
}

// GetChannelListReq 获取频道列表 request
type GetChannelListReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
}

func (p *GetChannelListReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	return nil
}

type (
	// GetChannelInfoReq 获取频道信息 request
	GetChannelInfoReq struct {
		ChannelId string `json:"channelId" binding:"required"` // 频道号
	}

	// GetChannelInfoRsp 获取频道信息 response
	GetChannelInfoRsp struct {
		ChannelElement

		IslandSourceId string `json:"islandSourceId"` // 群ID
	}
)

func (p *GetChannelInfoReq) ValidParams() error {
	if p.ChannelId == "" {
		return errors.New("invalid parameter ChannelId (empty detected)")
	}
	return nil
}

type (
	// CreateChannelReq 创建频道 request
	CreateChannelReq struct {
		IslandSourceId string      `json:"islandSourceId" binding:"required"` // 群ID
		ChannelName    string      `json:"channelName"`                       // 频道名称，非必传，不传时默认使用名称`新的频道`，不能大于32个字符或16个汉字
		ChannelType    ChannelType `json:"channelType" binding:"required"`    // 频道类型，1：文字频道，2：语音频道（默认自由模式），4：帖子频道（默认详细模式）
	}

	// CreateChannelRsp 创建频道 response
	CreateChannelRsp struct {
		ChannelId string `json:"channelId"` // 频道ID
	}
)

func (p *CreateChannelReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.ChannelType == 0 {
		return errors.New("invalid parameter ChannelType (zero detected)")
	}
	return nil
}

// EditChannelReq 编辑频道 request
type EditChannelReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	ChannelId      string `json:"channelId" binding:"required"`      // 频道号
	ChannelName    string `json:"channelName,omitempty"`
}

func (p *EditChannelReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.ChannelId == "" {
		return errors.New("invalid parameter ChannelId (empty detected)")
	}
	return nil
}

// RemoveChannelReq 删除频道 request
type RemoveChannelReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	ChannelId      string `json:"channelId" binding:"required"`      // 频道号
}

func (p *RemoveChannelReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.ChannelId == "" {
		return errors.New("invalid parameter ChannelId (empty detected)")
	}
	return nil
}

type (
	// SendChannelMessageReq 发送消息 request
	SendChannelMessageReq struct {
		ChannelId           string       `json:"channelId" binding:"required"`   // 频道号
		MessageType         MessageType  `json:"messageType" binding:"required"` // 消息类型，该参数会在SDK中重新赋值，所以无需开发者主动设值
		MessageBody         IMessageBody `json:"messageBody" binding:"required"` // 消息内容
		ReferencedMessageId string       `json:"referencedMessageId,omitempty"`  // 回复消息ID
		DodoSourceId        string       `json:"dodoSourceId,omitempty"`         // DoDoID，非必传，如果传了，则给该成员发送频道私信
	}

	// SendChannelMessageRsp 发送消息 response
	SendChannelMessageRsp struct {
		MessageId string `json:"messageId"` // 消息 ID
	}
)

func (p *SendChannelMessageReq) ValidParams() error {
	if p.ChannelId == "" {
		return errors.New("invalid parameter ChannelId (empty detected)")
	}
	if p.MessageBody == nil {
		return errors.New("invalid parameter MessageBody (nil detected)")
	}
	return nil
}

type (
	// EditChannelMessageReq 编辑消息 request
	EditChannelMessageReq struct {
		MessageId   string       `json:"messageId" binding:"required"`   // 欲编辑的消息 ID
		MessageType MessageType  `json:"messageType" binding:"required"` // 消息类型，该参数会在SDK中重新赋值，所以无需开发者主动设值
		MessageBody IMessageBody `json:"messageBody" binding:"required"` // 消息内容
	}

	// EditChannelMessageRsp 编辑消息 response
	EditChannelMessageRsp struct {
		MessageId string `json:"messageId"` // 消息 ID
	}
)

func (p *EditChannelMessageReq) ValidParams() error {
	if p.MessageId == "" {
		return errors.New("invalid parameter MessageId (empty detected)")
	}
	if p.MessageBody == nil {
		return errors.New("invalid parameter MessageBody (nil detected)")
	}
	return nil
}

// WithdrawChannelMessageReq 撤回消息 request
type WithdrawChannelMessageReq struct {
	MessageId string `json:"messageId" binding:"required"` // 消息ID
	Reason    string `json:"reason,omitempty"`             // 撤回原因
}

func (p *WithdrawChannelMessageReq) ValidParams() error {
	if p.MessageId == "" {
		return errors.New("invalid parameter MessageId (empty detected)")
	}
	return nil
}

// AddChannelMessageReactionReq 添加表情反应 request
type AddChannelMessageReactionReq struct {
	MessageId string         `json:"messageId" binding:"required"` // 消息 ID
	Emoji     *ReactionEmoji `json:"emoji" binding:"required"`     // 反应表情
}

func (p *AddChannelMessageReactionReq) ValidParams() error {
	if p.MessageId == "" {
		return errors.New("invalid parameter MessageId (empty detected)")
	}
	if p.Emoji == nil {
		return errors.New("invalid parameter Emoji (nil detected)")
	}
	if p.Emoji.Id == "" {
		return errors.New("invalid parameter Emoji.Id (empty detected)")
	}
	return nil
}

// RemChannelMessageReactionReq 取消表情反应 request
type RemChannelMessageReactionReq struct {
	MessageId    string         `json:"messageId" binding:"required"` // 消息 ID
	Emoji        *ReactionEmoji `json:"emoji" binding:"required"`     // 反应表情
	DodoSourceId string         `json:"dodoSourceId,omitempty"`       // DoDoID，不传或传空时表示移除机器人自身的反应
}

func (p *RemChannelMessageReactionReq) ValidParams() error {
	if p.MessageId == "" {
		return errors.New("invalid parameter MessageId (empty detected)")
	}
	if p.Emoji == nil {
		return errors.New("invalid parameter Emoji (nil detected)")
	}
	if p.Emoji.Id == "" {
		return errors.New("invalid parameter Emoji.Id (empty detected)")
	}
	return nil
}
