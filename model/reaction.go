package model

// ReactionTarget 反映对象
type ReactionTarget struct {
	Type int    `json:"type"` // 对象类型，0：消息
	Id   string `json:"id"`   // 对象ID，若对象类型为0，则代表消息ID
}

// ReactionEmoji 表情内容
// 表情 ID 参考文档：https://open.imdodo.com/api/message/emoji.html#%E8%A1%A8%E6%83%85%E5%86%85%E5%AE%B9
type ReactionEmoji struct {
	Type int    `json:"type"` // 表情类型，1：Emoji
	Id   string `json:"id"`   // 表情 ID
}
