package model

type (
	// IMessageBody 消息内容
	IMessageBody interface {
		MessageType() int // 获取消息类型
	}

	// ChannelTextMessage 频道文本消息内容
	ChannelTextMessage struct {
		Content string `json:"content" binding:"required"` // 文本内容
	}

	// ChannelImageMessage 频道图片消息内容
	ChannelImageMessage struct {
		Url        string `json:"url" binding:"required"`    // 图片链接，必须是官方的
		Width      int    `json:"width" binding:"required"`  // 图片宽度
		Height     int    `json:"height" binding:"required"` // 图片高度
		IsOriginal int    `json:"isOriginal,omitempty"`      // 是否原图，0：压缩图，1：原图
	}

	// ChannelVideoMessage 频道视频消息内容
	ChannelVideoMessage struct {
		Url      string `json:"url" binding:"required"` // 视频链接，必须是官方的（后期会提供视频上传接口，前期可通过DoDo群内上传资源，事件内接收消息，从而获取到官方资源链接）
		CoverUrl string `json:"coverUrl,omitempty"`     // 封面链接，必须是官方的
		Duration uint64 `json:"duration,omitempty"`     // 视频时长
		Size     uint64 `json:"size,omitempty"`         // 视频大小
	}

	// ChannelFileMessage 频道文件消息内容
	ChannelFileMessage struct {
		Url  string `json:"url" binding:"required"`  // 文件链接
		Name string `json:"name" binding:"required"` // 文件名称
		Size uint64 `json:"size" binding:"required"` // 文件大小
	}
)

func (m *ChannelTextMessage) MessageType() int {
	return 1
}

func (m *ChannelImageMessage) MessageType() int {
	return 2
}

func (m *ChannelVideoMessage) MessageType() int {
	return 3
}

func (m *ChannelFileMessage) MessageType() int {
	return 5
}