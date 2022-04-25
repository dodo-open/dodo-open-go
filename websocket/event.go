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
func ParseDataAndHandle(event *WSEventMessage) error {
	data := &EventData{}
	if err := tools.JSON.Unmarshal(event.Data, &data); err != nil {
		return err
	}
	// match event message parser by EventType
	if handle, ok := eventParserMap[data.EventType]; ok {
		return handle(event, event.RawData)
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
type eventParser func(event *WSEventMessage, message []byte) error

func personalMessageHandler(event *WSEventMessage, message []byte) error {
	data := &PersonalMessageEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.PersonalMessage != nil {
		return DefaultHandlers.PersonalMessage(event, data)
	}
	return nil
}

func channelMessageHandler(event *WSEventMessage, message []byte) error {
	data := &ChannelMessageEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.ChannelMessage != nil {
		return DefaultHandlers.ChannelMessage(event, data)
	}
	return nil
}

func messageReactionHandler(event *WSEventMessage, message []byte) error {
	data := &MessageReactionEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageReaction != nil {
		return DefaultHandlers.MessageReaction(event, data)
	}
	return nil
}

func memberJoinHandler(event *WSEventMessage, message []byte) error {
	data := &MemberJoinEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MemberJoin != nil {
		return DefaultHandlers.MemberJoin(event, data)
	}
	return nil
}

func memberLeaveHandler(event *WSEventMessage, message []byte) error {
	data := &MemberLeaveEventBody{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MemberLeave != nil {
		return DefaultHandlers.MemberLeave(event, data)
	}
	return nil
}
