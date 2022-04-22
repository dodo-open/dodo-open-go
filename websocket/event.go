package websocket

// EventType event type code
type EventType int

const (
	PersonalMessageEvent EventType = 1001 // 个人消息事件
	ChannelMessageEvent  EventType = 2001 // 频道消息事件
	MessageReactionEvent EventType = 3001 // 消息反应事件
	MemberJoinEvent      EventType = 4001 // 成员加入事件
	MemberLeaveEvent     EventType = 4002 // 成员退出事件
)
