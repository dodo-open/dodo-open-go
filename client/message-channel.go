package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/errs"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// SendChannelMessage 发送频道消息
// model.SendChannelMessageReq 对象中的 MessageType 参数会在SDK中重新赋值，所以无需开发者关注
func (c *client) SendChannelMessage(ctx context.Context, req *model.SendChannelMessageReq) (*model.SendChannelMessageRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	req.MessageType = req.MessageBody.MessageType()
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(sendChannelMessageUri))
	if err != nil {
		return nil, err
	}

	result := &model.SendChannelMessageRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// EditChannelMessage 编辑频道消息
// model.EditChannelMessageReq 对象中的 MessageType 参数会在SDK中重新赋值，所以无需开发者关注
func (c *client) EditChannelMessage(ctx context.Context, req *model.EditChannelMessageReq) (*model.EditChannelMessageRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	req.MessageType = req.MessageBody.MessageType()
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(editChannelMessageUri))
	if err != nil {
		return nil, err
	}

	result := &model.EditChannelMessageRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// WithdrawChannelMessage 撤回频道消息
func (c *client) WithdrawChannelMessage(ctx context.Context, req *model.WithdrawChannelMessageReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(withdrawChannelMessageUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// SetChannelMessageReaction 频道消息添加反应
// Deprecated: 这个方法会在 ver.0.0.8 版本往后的某次平台更新后废弃，请尽快切换至新的方法
// Deprecated: 新的方法请参考 AddChannelMessageReaction 和 RemChannelMessageReaction
func (c *client) SetChannelMessageReaction(ctx context.Context, req *model.SetChannelMessageReactionReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setChannelMessageReaction))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// AddChannelMessageReaction 添加频道消息反应
func (c *client) AddChannelMessageReaction(ctx context.Context, req *model.AddChannelMessageReactionReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(addChannelMessageReaction))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// RemChannelMessageReaction 移除文字频道消息反应
func (c *client) RemChannelMessageReaction(ctx context.Context, req *model.RemChannelMessageReactionReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(remChannelMessageReaction))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}
