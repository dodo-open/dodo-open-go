package model

import "errors"

// SendDirectMessageReq 发送私聊消息 request
type SendDirectMessageReq struct {
	DodoId      string       `json:"dodoId" binding:"required"`      // DoDo号
	MessageType int          `json:"messageType" binding:"required"` // 消息类型，该参数会在SDK中重新赋值，所以无需开发者主动设值
	MessageBody IMessageBody `json:"messageBody" binding:"required"` // 消息内容
}

func (p *SendDirectMessageReq) ValidParams() error {
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	if p.MessageBody == nil {
		return errors.New("invalid parameter MessageBody (nil detected)")
	}
	return nil
}

// SendDirectMessageRsp 发送私聊消息 response
type SendDirectMessageRsp struct {
	MessageId string `json:"messageId"` // 消息 ID
}
