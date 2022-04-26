package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetIslandList 取群列表
func (c *client) GetIslandList(ctx context.Context) ([]*model.IslandElement, error) {
	list := make([]*model.IslandElement, 0)

	resp, err := c.request(ctx).SetBody(&model.GetIslandListReq{}).Post(c.getApi(getIslandListUri))
	if err != nil {
		return list, err
	}

	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &list); err != nil {
		return list, err
	}
	return list, nil
}

// GetIslandInfo 取群信息
func (c *client) GetIslandInfo(ctx context.Context, req *model.GetIslandInfoReq) (*model.GetIslandInfoRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getIslandInfoUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetIslandInfoRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
