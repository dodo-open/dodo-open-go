package model

type MessageType int

const (
	TextMsg  MessageType = 1 // 频道文本消息
	ImageMsg MessageType = 2 // 频道图片消息
	VideoMsg MessageType = 3 // 频道视频消息
	FileMsg  MessageType = 5 // 频道文件消息
	CardMsg  MessageType = 6 // 频道卡片消息
	Unknown  MessageType = 0 // 未知格式消息
)

type (
	// IMessageBody 消息内容
	IMessageBody interface {
		MessageType() MessageType // 获取消息类型
	}

	// TextMessage 频道文本消息内容
	TextMessage struct {
		Content string `json:"content" binding:"required"` // 文本内容
	}

	// ImageMessage 频道图片消息内容
	ImageMessage struct {
		Url        string `json:"url" binding:"required"`    // 图片链接，必须是官方的
		Width      int    `json:"width" binding:"required"`  // 图片宽度
		Height     int    `json:"height" binding:"required"` // 图片高度
		IsOriginal int    `json:"isOriginal,omitempty"`      // 是否原图，0：压缩图，1：原图
	}

	// VideoMessage 频道视频消息内容
	VideoMessage struct {
		Url      string `json:"url" binding:"required"` // 视频链接，必须是官方的（后期会提供视频上传接口，前期可通过DoDo群内上传资源，事件内接收消息，从而获取到官方资源链接）
		CoverUrl string `json:"coverUrl,omitempty"`     // 封面链接，必须是官方的
		Duration uint64 `json:"duration,omitempty"`     // 视频时长
		Size     uint64 `json:"size,omitempty"`         // 视频大小
	}

	// FileMessage 频道文件消息内容
	FileMessage struct {
		Url  string `json:"url" binding:"required"`  // 文件链接
		Name string `json:"name" binding:"required"` // 文件名称
		Size uint64 `json:"size" binding:"required"` // 文件大小
	}

	// CardMessage 频道卡片消息内容
	CardMessage struct {
		Content string           `json:"content,omitempty"`       // 附加文本，支持Markdown语法、菱形语法
		Card    *CardBodyElement `json:"card" binding:"required"` // 卡片，限制 10000 个字符，支持 Markdown 语法，不支持菱形语法
	}

	// UnknownMessage 未知格式消息内容
	UnknownMessage struct {
	}
)

func (m *TextMessage) MessageType() MessageType {
	return TextMsg
}

func (m *ImageMessage) MessageType() MessageType {
	return ImageMsg
}

func (m *VideoMessage) MessageType() MessageType {
	return VideoMsg
}

func (m *FileMessage) MessageType() MessageType {
	return FileMsg
}

func (m *CardMessage) MessageType() MessageType {
	return CardMsg
}

func (m *UnknownMessage) MessageType() MessageType {
	return Unknown
}

// CardBodyElement 卡片消息结构体
type CardBodyElement struct {
	Type       string        `json:"type" binding:"required"`       // 类型，固定填写 card
	Components []interface{} `json:"components" binding:"required"` // 内容组件
	Theme      string        `json:"theme" binding:"required"`      // 卡片风格，grey，red，orange，yellow ，green，indigo，blue，purple，black，default
	Title      string        `json:"title,omitempty"`               // 卡片标题，只支持普通文本，可以为空字符串
}
