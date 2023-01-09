package websocket

import "github.com/dodo-open/dodo-open-go/tools"

// EventType event type code
type EventType string

const (
	PersonalMessageEvent         EventType = "1001" // 个人消息事件
	ChannelMessageEvent          EventType = "2001" // 频道消息事件
	MessageReactionEvent         EventType = "3001" // 消息反应事件
	MemberJoinEvent              EventType = "4001" // 成员加入事件
	MemberLeaveEvent             EventType = "4002" // 成员退出事件
	ChannelVoiceMemberJoinEvent  EventType = "5001" // 成员加入语音频道事件
	ChannelVoiceMemberLeaveEvent EventType = "5002" // 成员退出语音频道事件
	ChannelArticleEvent          EventType = "6001" // 帖子发布事件
	ChannelArticleCommentEvent   EventType = "6002" // 帖子评论回复事件
	GiftSendEvent                EventType = "7001" // 赠礼成功事件
)

// eventParserMap event parser map, for safety, do not modify this map
var eventParserMap = map[EventType]eventParser{
	PersonalMessageEvent:         personalMessageHandler,
	ChannelMessageEvent:          channelMessageHandler,
	MessageReactionEvent:         messageReactionHandler,
	MemberJoinEvent:              memberJoinHandler,
	MemberLeaveEvent:             memberLeaveHandler,
	ChannelVoiceMemberJoinEvent:  channelVoiceMemberJoinHandler,
	ChannelVoiceMemberLeaveEvent: channelVoiceMemberLeaveHandler,
	ChannelArticleEvent:          channelArticleHandler,
	ChannelArticleCommentEvent:   channelArticleCommentHandler,
	GiftSendEvent:                giftSendHandler,
}

// ParseDataAndHandle parse message data and handle it
func (c *client) ParseDataAndHandle(event *WSEventMessage) error {
	data := &EventData{}
	if err := tools.JSON.Unmarshal(event.Data, &data); err != nil {
		return err
	}
	// match event message parser by EventType
	if handle, ok := eventParserMap[data.EventType]; ok {
		return handle(c, event, event.Data)
	}
	// else treat the message as plain text message
	if DefaultHandlers.PlainTextHandler != nil {
		return DefaultHandlers.PlainTextHandler(event, event.Data)
	}
	return nil
}

// ParseData parse message data
func ParseData(message []byte, v interface{}) error {
	data := tools.JSON.Get(message, "eventBody")
	return tools.JSON.Unmarshal([]byte(data.ToString()), v)
}

// eventParser WebSocket message parser func
type eventParser func(c *client, event *WSEventMessage, message []byte) error

func personalMessageHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &PersonalMessageEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.PersonalMessage != nil {
		return c.conf.messageHandlers.PersonalMessage(event, data)
	}
	if DefaultHandlers.PersonalMessage != nil {
		return DefaultHandlers.PersonalMessage(event, data)
	}
	return nil
}

func channelMessageHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &ChannelMessageEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.ChannelMessage != nil {
		return c.conf.messageHandlers.ChannelMessage(event, data)
	}
	if DefaultHandlers.ChannelMessage != nil {
		return DefaultHandlers.ChannelMessage(event, data)
	}
	return nil
}

func messageReactionHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &MessageReactionEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.MessageReaction != nil {
		return c.conf.messageHandlers.MessageReaction(event, data)
	}
	if DefaultHandlers.MessageReaction != nil {
		return DefaultHandlers.MessageReaction(event, data)
	}
	return nil
}

func memberJoinHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &MemberJoinEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.MemberJoin != nil {
		return c.conf.messageHandlers.MemberJoin(event, data)
	}
	if DefaultHandlers.MemberJoin != nil {
		return DefaultHandlers.MemberJoin(event, data)
	}
	return nil
}

func memberLeaveHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &MemberLeaveEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.MemberLeave != nil {
		return c.conf.messageHandlers.MemberLeave(event, data)
	}
	if DefaultHandlers.MemberLeave != nil {
		return DefaultHandlers.MemberLeave(event, data)
	}
	return nil
}

func channelVoiceMemberJoinHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &ChannelVoiceMemberJoinEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.ChannelVoiceMemberJoin != nil {
		return c.conf.messageHandlers.ChannelVoiceMemberJoin(event, data)
	}
	if DefaultHandlers.ChannelVoiceMemberJoin != nil {
		return DefaultHandlers.ChannelVoiceMemberJoin(event, data)
	}
	return nil
}

func channelVoiceMemberLeaveHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &ChannelVoiceMemberLeaveEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.ChannelVoiceMemberLeave != nil {
		return c.conf.messageHandlers.ChannelVoiceMemberLeave(event, data)
	}
	if DefaultHandlers.ChannelVoiceMemberLeave != nil {
		return DefaultHandlers.ChannelVoiceMemberLeave(event, data)
	}
	return nil
}

func channelArticleHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &ChannelArticleEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.ChannelArticle != nil {
		return c.conf.messageHandlers.ChannelArticle(event, data)
	}
	if DefaultHandlers.ChannelArticle != nil {
		return DefaultHandlers.ChannelArticle(event, data)
	}
	return nil
}

func channelArticleCommentHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &ChannelArticleCommentEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.ChannelArticleComment != nil {
		return c.conf.messageHandlers.ChannelArticleComment(event, data)
	}
	if DefaultHandlers.ChannelArticleComment != nil {
		return DefaultHandlers.ChannelArticleComment(event, data)
	}
	return nil
}

func giftSendHandler(c *client, event *WSEventMessage, message []byte) error {
	data := &GiftSendEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if c.conf.messageHandlers.GiftSend != nil {
		return c.conf.messageHandlers.GiftSend(event, data)
	}
	if DefaultHandlers.GiftSend != nil {
		return DefaultHandlers.GiftSend(event, data)
	}
	return nil
}
