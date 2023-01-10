package model

import "errors"

// SendDirectMessageReq 发送私信 request
type SendDirectMessageReq struct {
	IslandSourceId string       `json:"islandSourceId" binding:"required"` // 来源群ID，可自行指定，亦可从私信事件中获取
	DodoSourceId   string       `json:"dodoSourceId" binding:"required"`   // DoDoID
	MessageType    MessageType  `json:"messageType" binding:"required"`    // 消息类型，该参数会在SDK中重新赋值，所以无需开发者主动设值
	MessageBody    IMessageBody `json:"messageBody" binding:"required"`    // 消息内容
}

func (p *SendDirectMessageReq) ValidParams() error {
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	if p.MessageBody == nil {
		return errors.New("invalid parameter MessageBody (nil detected)")
	}
	return nil
}

// SendDirectMessageRsp 发送私信 response
type SendDirectMessageRsp struct {
	MessageId string `json:"messageId"` // 消息 ID
}
