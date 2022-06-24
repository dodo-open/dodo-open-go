package client

import (
	"context"
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
