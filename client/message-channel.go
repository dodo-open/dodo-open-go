package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/errs"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// SendChannelMessage 发送消息
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

// EditChannelMessage 编辑消息
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

// WithdrawChannelMessage 撤回消息
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

// AddChannelMessageReaction 添加表情反应
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

// RemChannelMessageReaction 取消表情反应
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

// SetChannelMessageTop 置顶消息
func (c *client) SetChannelMessageTop(ctx context.Context, req *model.SetChannelMessageTopReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setChannelMessageTopUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// GetChannelMessageReactionList 获取消息反应列表
func (c *client) GetChannelMessageReactionList(ctx context.Context, req *model.GetChannelMessageReactionListReq) ([]*model.GetChannelMessageReactionListRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getChannelMessageReactionListUri))
	if err != nil {
		return nil, err
	}

	result := make([]*model.GetChannelMessageReactionListRsp, 0)
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetChannelMessageReactionMemberList 获取消息反应内成员列表
func (c *client) GetChannelMessageReactionMemberList(ctx context.Context, req *model.GetChannelMessageReactionMemberListReq) (*model.GetChannelMessageReactionMemberListRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getChannelMessageReactionMemberListUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetChannelMessageReactionMemberListRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
