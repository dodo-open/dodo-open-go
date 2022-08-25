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

// ChannelElement 取频道列表 list element
type ChannelElement struct {
	ChannelId   string      `json:"channelId"`   // 频道号
	ChannelName string      `json:"channelName"` // 频道名称
	ChannelType ChannelType `json:"channelType"` // 频道类型，1：文字频道，2：语音频道，4：帖子频道，5：链接频道，6：资料频道
	DefaultFlag int         `json:"defaultFlag"` // 默认频道标识，0：否，1：是
	GroupId     string      `json:"groupId"`     // 分组ID
	GroupName   string      `json:"groupName"`   // 分组名称
}

// GetChannelListReq 取频道列表 request
type GetChannelListReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
}

func (p *GetChannelListReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	return nil
}

type (
	// GetChannelInfoReq 取频道信息 request
	GetChannelInfoReq struct {
		ChannelId string `json:"channelId" binding:"required"` // 频道号
	}

	// GetChannelInfoRsp 取频道信息 response
	GetChannelInfoRsp struct {
		ChannelElement

		IslandId string `json:"islandId"` // 群号
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
		IslandId    string      `json:"islandId" binding:"required"`    // 群号
		ChannelName string      `json:"channelName"`                    // 频道名称，非必传，不传时默认使用名称`新的频道`，不能大于32个字符或16个汉字
		ChannelType ChannelType `json:"channelType" binding:"required"` // 频道类型，1：文字频道，2：语音频道（默认自由模式），4：帖子频道（默认详细模式）
	}

	// CreateChannelRsp 创建频道 response
	CreateChannelRsp struct {
		ChannelId string `json:"channelId"` // 频道ID
	}
)

func (p *CreateChannelReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.ChannelType == 0 {
		return errors.New("invalid parameter ChannelType (zero detected)")
	}
	return nil
}

// EditChannelReq 编辑频道 request
type EditChannelReq struct {
	IslandId    string `json:"islandId" binding:"required"`  // 群号
	ChannelId   string `json:"channelId" binding:"required"` // 频道号
	ChannelName string `json:"channelName"`
}

func (p *EditChannelReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.ChannelId == "" {
		return errors.New("invalid parameter ChannelId (empty detected)")
	}
	return nil
}

// RemoveChannelReq 删除频道 request
type RemoveChannelReq struct {
	IslandId  string `json:"islandId" binding:"required"`  // 群号
	ChannelId string `json:"channelId" binding:"required"` // 频道号
}

func (p *RemoveChannelReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.ChannelId == "" {
		return errors.New("invalid parameter ChannelId (empty detected)")
	}
	return nil
}

type (
	// SendChannelMessageReq 发送频道消息 request
	SendChannelMessageReq struct {
		ChannelId           string       `json:"channelId" binding:"required"`   // 频道号
		MessageType         MessageType  `json:"messageType" binding:"required"` // 消息类型，该参数会在SDK中重新赋值，所以无需开发者主动设值
		MessageBody         IMessageBody `json:"messageBody" binding:"required"` // 消息内容
		ReferencedMessageId string       `json:"referencedMessageId,omitempty"`  // 回复消息ID
	}

	// SendChannelMessageRsp 发送频道消息 response
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
	// EditChannelMessageReq 编辑频道消息 request
	EditChannelMessageReq struct {
		MessageId   string       `json:"messageId" binding:"required"`   // 欲编辑的消息 ID
		MessageType MessageType  `json:"messageType" binding:"required"` // 消息类型，该参数会在SDK中重新赋值，所以无需开发者主动设值
		MessageBody IMessageBody `json:"messageBody" binding:"required"` // 消息内容
	}

	// EditChannelMessageRsp 发送频道消息 response
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

// WithdrawChannelMessageReq 撤回频道消息 request
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

// SetChannelMessageReactionReq 频道消息添加反应 request
type SetChannelMessageReactionReq struct {
	ReactionTarget *ReactionTarget `json:"reactionTarget" binding:"required"` // 反应对象
	ReactionEmoji  *ReactionEmoji  `json:"reactionEmoji" binding:"required"`  // 反应表情
	ReactionType   int             `json:"reactionType" binding:"required"`   // 反应类型，0：删除，1：新增
}

func (p *SetChannelMessageReactionReq) ValidParams() error {
	if p.ReactionTarget == nil {
		return errors.New("invalid parameter ReactionTarget (nil detected)")
	}
	if p.ReactionTarget.Id == "" {
		return errors.New("invalid parameter ReactionTarget.Id (empty detected)")
	}
	if p.ReactionEmoji == nil {
		return errors.New("invalid parameter ReactionEmoji (nil detected)")
	}
	if p.ReactionEmoji.Id == "" {
		return errors.New("invalid parameter ReactionEmoji.Id (empty detected)")
	}
	if p.ReactionType != 0 && p.ReactionType != 1 {
		return errors.New("invalid parameter ReactionType (should be 0 or 1)")
	}
	return nil
}

// AddChannelMessageReactionReq 新增文字频道消息反应 request
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

// RemChannelMessageReactionReq 移除文字频道消息反应 request
type RemChannelMessageReactionReq struct {
	MessageId string         `json:"messageId" binding:"required"` // 消息 ID
	Emoji     *ReactionEmoji `json:"emoji" binding:"required"`     // 反应表情
	DoDoId    string         `json:"dodoId,omitempty"`             // DoDo号，不传或传空时表示移除机器人自身的反应
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
