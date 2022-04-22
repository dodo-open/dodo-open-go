package client

import (
	"context"
	"dodo-open-go/model"
	"dodo-open-go/tools"
	"errors"
)

// GetChannelList 取频道列表
func (c *client) GetChannelList(ctx context.Context, req *model.GetChannelListReq) ([]*model.ChannelElement, error) {
	list := make([]*model.ChannelElement, 0)

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getChannelListUri))
	if err != nil {
		return list, err
	}

	if err = tools.JSON.Unmarshal(resp.Result().(*model.OpenAPIRsp).Data, &list); err != nil {
		return list, err
	}
	return list, nil
}

// GetChannelInfo 取频道信息
func (c *client) GetChannelInfo(ctx context.Context, req *model.GetChannelInfoReq) (*model.GetChannelInfoRsp, error) {
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getChannelInfoUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetChannelInfoRsp{}
	if err = tools.JSON.Unmarshal(resp.Result().(*model.OpenAPIRsp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SendChannelMessage 发送频道消息
// model.SendChannelMessageReq 对象中的 MessageType 参数会在SDK中重新赋值，所以无需开发者关注
func (c *client) SendChannelMessage(ctx context.Context, req *model.SendChannelMessageReq) (*model.SendChannelMessageRsp, error) {
	if req.MessageBody == nil {
		return nil, errors.New("invalid SendChannelMessageReq object (MessageBody nil detected)")
	}
	req.MessageType = req.MessageBody.MessageType()
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(sendChannelMessageUri))
	if err != nil {
		return nil, err
	}

	result := &model.SendChannelMessageRsp{}
	if err = tools.JSON.Unmarshal(resp.Result().(*model.OpenAPIRsp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
