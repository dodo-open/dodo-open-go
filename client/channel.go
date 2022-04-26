package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/errs"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetChannelList 取频道列表
func (c *client) GetChannelList(ctx context.Context, req *model.GetChannelListReq) ([]*model.ChannelElement, error) {
	list := make([]*model.ChannelElement, 0)

	if err := req.ValidParams(); err != nil {
		return list, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getChannelListUri))
	if err != nil {
		return list, err
	}

	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &list); err != nil {
		return list, err
	}
	return list, nil
}

// GetChannelInfo 取频道信息
func (c *client) GetChannelInfo(ctx context.Context, req *model.GetChannelInfoReq) (*model.GetChannelInfoRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getChannelInfoUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetChannelInfoRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

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
