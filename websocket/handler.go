package websocket

// MessageHandlers instance message handlers
type MessageHandlers struct {
	PersonalMessage PersonalMessageEventHandler
	ChannelMessage  ChannelMessageEventHandler
	MessageReaction MessageReactionEventHandler
	MemberJoin      MemberJoinEventHandler
	MemberLeave     MemberLeaveEventHandler

	PlainTextHandler PlainTextHandler
	ErrorHandler     ErrorHandler
}

// DefaultHandlers default handlers to manage all supported message
var DefaultHandlers struct {
	MessageHandlers
}

func fillHandler(handlers *MessageHandlers) *MessageHandlers {
	if handlers.PersonalMessage == nil {
		handlers.PersonalMessage = DefaultHandlers.PersonalMessage
	}
	if handlers.ChannelMessage == nil {
		handlers.ChannelMessage = DefaultHandlers.ChannelMessage
	}
	if handlers.MessageReaction == nil {
		handlers.MessageReaction = DefaultHandlers.MessageReaction
	}
	if handlers.MemberJoin == nil {
		handlers.MemberJoin = DefaultHandlers.MemberJoin
	}
	if handlers.MemberLeave == nil {
		handlers.MemberLeave = DefaultHandlers.MemberLeave
	}
	if handlers.PlainTextHandler == nil {
		handlers.PlainTextHandler = DefaultHandlers.PlainTextHandler
	}
	if handlers.ErrorHandler == nil {
		handlers.ErrorHandler = DefaultHandlers.ErrorHandler
	}
	return handlers
}

// PersonalMessageEventHandler 个人消息事件 handler
type PersonalMessageEventHandler func(event *WSEventMessage, data *PersonalMessageEventBody) error

// ChannelMessageEventHandler 频道消息事件 handler
type ChannelMessageEventHandler func(event *WSEventMessage, data *ChannelMessageEventBody) error

// MessageReactionEventHandler 消息反应事件 handler
type MessageReactionEventHandler func(event *WSEventMessage, data *MessageReactionEventBody) error

// MemberJoinEventHandler 成员加入事件 handler
type MemberJoinEventHandler func(event *WSEventMessage, data *MemberJoinEventBody) error

// MemberLeaveEventHandler 成员退出事件 handler
type MemberLeaveEventHandler func(event *WSEventMessage, data *MemberLeaveEventBody) error

// PlainTextHandler plain text message handler
type PlainTextHandler func(event *WSEventMessage, message []byte) error

// ErrorHandler error handler
type ErrorHandler func(err error)

// RegisterHandlers Register global level event message handlers
func RegisterHandlers(handlers ...interface{}) {
	for _, h := range handlers {
		switch handle := h.(type) {
		case PersonalMessageEventHandler:
			DefaultHandlers.PersonalMessage = handle
		case ChannelMessageEventHandler:
			DefaultHandlers.ChannelMessage = handle
		case MessageReactionEventHandler:
			DefaultHandlers.MessageReaction = handle
		case MemberJoinEventHandler:
			DefaultHandlers.MemberJoin = handle
		case MemberLeaveEventHandler:
			DefaultHandlers.MemberLeave = handle
		default:
			// other handlers will be registered in the following functions
			// non-business handler will be registered here
			registerOtherHandlers(handlers)
		}
	}
}

// registerOtherHandlers Register global level non-business handlers
func registerOtherHandlers(handlers ...interface{}) {
	for _, h := range handlers {
		switch handle := h.(type) {
		case PlainTextHandler:
			DefaultHandlers.PlainTextHandler = handle
		case ErrorHandler:
			DefaultHandlers.ErrorHandler = handle
		default:
		}
	}
}
