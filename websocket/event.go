package websocket

import (
	"dodo-open-go/tools"
)

// EventType event type code
type EventType string

const (
	PersonalMessageEvent EventType = "1001" // 个人消息事件
	ChannelMessageEvent  EventType = "2001" // 频道消息事件
	MessageReactionEvent EventType = "3001" // 消息反应事件
	MemberJoinEvent      EventType = "4001" // 成员加入事件
	MemberLeaveEvent     EventType = "4002" // 成员退出事件
)

// eventParserMap event parser map, for safety, do not modify this map
var eventParserMap = map[EventType]eventParser{
	PersonalMessageEvent: personalMessageHandler,
	ChannelMessageEvent:  channelMessageHandler,
	MessageReactionEvent: messageReactionHandler,
	MemberJoinEvent:      memberJoinHandler,
	MemberLeaveEvent:     memberLeaveHandler,
}

// ParseDataAndHandle parse message data and handle it
func (c *client) ParseDataAndHandle(event *WSEventMessage) error {
	data := &EventData{}
	if err := tools.JSON.Unmarshal(event.Data, &data); err != nil {
		return err
	}
	// match event message parser by EventType
	if handle, ok := eventParserMap[data.EventType]; ok {
		return handle(c, event, event.RawData)
	}
	// else treat the message as plain text message
	if DefaultHandlers.PlainTextHandler != nil {
		return DefaultHandlers.PlainTextHandler(event, event.RawData)
	}
	return nil
}

// ParseData parse message data
func ParseData(message []byte, v interface{}) error {
	data := tools.JSON.Get(message, "data", "eventBody")
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
