package websocket

import "dodo-open-go/model"

// WSEventMessage 事件主体内容
type WSEventMessage struct {
	Type    TypeCode    `json:"type"` // 数据类型
	Data    WSEventData `json:"data"` // 数据内容
	RawData []byte      `json:"-"`    // 原始数据
}

// WSEventData 事件业务数据
type WSEventData struct {
	EventId   string    `json:"eventId"`   // 事件ID
	EventType EventType `json:"eventType"` // 事件类型
	EventBody EventBody `json:"eventBody"` // 事件内容
	Timestamp uint64    `json:"timestamp"` // 发送时间戳
}

// WSBeatData 心跳类型数据
type WSBeatData struct {
	Type TypeCode `json:"type"`
}

// EventBody 事件内容
type EventBody interface {
	EventType() EventType
}

type (
	// PersonalModel 个人信息
	PersonalModel struct {
		NickName  string `json:"nickName"`  // 个人昵称
		AvatarUrl string `json:"avatarUrl"` // 个人头像
		Sex       int    `json:"sex"`       // 个人性别，-1：保密，0：女，1：男
	}

	// MemberModel 成员信息
	MemberModel struct {
		NickName string `json:"nickName"` // 在群昵称
		JoinTime string `json:"joinTime"` // 加群时间
	}

	// ReferenceModel 回复信息
	ReferenceModel struct {
		MessageId string `json:"messageId"` // 被回复消息ID
		DodoId    string `json:"dodoId"`    // 被回复者DoDo号
		NickName  string `json:"nickName"`  // 被回复者在群昵称
	}
)

type (
	// ReactionTargetModel 反映对象
	ReactionTargetModel struct {
		Type int    `json:"type"` // 对象类型，0：消息
		Id   string `json:"id"`   // 对象ID，若对象类型为0，则代表消息ID
	}

	// EmojiModel 表情内容
	// 表情 ID 参考文档：https://open.imdodo.com/api/message/emoji.html#%E8%A1%A8%E6%83%85%E5%86%85%E5%AE%B9
	EmojiModel struct {
		Type int    `json:"type"` // 表情类型，1：Emoji
		Id   string `json:"id"`   // 表情 ID
	}
)

type (
	// PersonalMessageEventBody 个人消息事件
	PersonalMessageEventBody struct {
		DodoId      string             `json:"dodoId"`      // 来源DoDo号
		Personal    *PersonalModel     `json:"personal"`    // 个人信息
		MessageId   string             `json:"messageId"`   // 消息ID
		MessageType int                `json:"messageType"` // 消息类型，1：文本消息，2：图片消息，3：视频消息
		MessageBody model.IMessageBody `json:"messageBody"` // 消息内容
	}

	// ChannelMessageEventBody 频道消息事件
	ChannelMessageEventBody struct {
		IslandId    string             `json:"islandId"`    // 来源群号
		ChannelId   string             `json:"channelId"`   // 来源频道号
		DodoId      string             `json:"dodoId"`      // 来源DoDo号
		Personal    *PersonalModel     `json:"personal"`    // 个人信息
		Member      *MemberModel       `json:"member"`      // 成员信息
		Reference   *ReferenceModel    `json:"reference"`   // 回复信息
		MessageId   string             `json:"messageId"`   // 消息ID
		MessageType int                `json:"messageType"` // 消息类型，1：文本消息，2：图片消息，3：视频消息，5：文件消息
		MessageBody model.IMessageBody `json:"messageBody"` // 消息内容
	}

	// MessageReactionEventBody 消息反应事件
	MessageReactionEventBody struct {
		IslandId       string               `json:"islandId"`       // 来源群号
		ChannelId      string               `json:"channelId"`      // 来源频道号
		DodoId         string               `json:"dodoId"`         // 来源DoDo号
		ReactionTarget *ReactionTargetModel `json:"reactionTarget"` // 反应对象
		ReactionEmoji  *EmojiModel          `json:"reactionEmoji"`  // 反应表情
		ReactionType   int                  `json:"reactionType"`   // 反应类型，0：删除，1：新增
	}

	// MemberJoinEventBody 成员加入事件
	MemberJoinEventBody struct {
		IslandId   string `json:"islandId"`   // 来源群号
		DodoId     string `json:"dodoId"`     // 来源DoDo号
		ModifyTime string `json:"modifyTime"` // 变动时间
	}

	// MemberLeaveEventBody 成员退出事件
	MemberLeaveEventBody struct {
		IslandId      string `json:"islandId"`      // 来源群号
		DodoId        string `json:"dodoId"`        // 来源DoDo号
		LeaveType     int    `json:"leaveType"`     // 退出类型，1：主动，2：被踢
		OperateDoDoId string `json:"operateDoDoId"` // 操作者DoDo号（执行踢出操作的人）
		ModifyTime    string `json:"modifyTime"`    // 变动时间
	}
)

func (e *PersonalMessageEventBody) EventType() EventType {
	return PersonalMessageEvent
}

func (e *ChannelMessageEventBody) EventType() EventType {
	return ChannelMessageEvent
}

func (e *MessageReactionEventBody) EventType() EventType {
	return MessageReactionEvent
}

func (e *MemberJoinEventBody) EventType() EventType {
	return MemberJoinEvent
}

func (e *MemberLeaveEventBody) EventType() EventType {
	return MemberLeaveEvent
}
