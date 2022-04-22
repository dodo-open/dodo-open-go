package client

import (
	"context"
	"dodo-open-go/model"
	"dodo-open-go/tools"
	"fmt"
)

// GetIslandList 取群列表
func (c *client) GetIslandList(ctx context.Context) ([]*model.IslandElement, error) {
	list := make([]*model.IslandElement, 0)

	resp, err := c.request(ctx).SetBody(&model.GetIslandListReq{}).Post(c.getApi(getIslandList))
	if err != nil {
		return list, err
	}

	data := resp.Result().(*model.OpenApiRpcRsp)
	fmt.Printf("%#v\n", data)

	if err = tools.JSON.Unmarshal(resp.Result().(*model.OpenApiRpcRsp).Data, &list); err != nil {
		return list, err
	}
	return list, nil
}

// GetIslandInfo 取群信息
func (c *client) GetIslandInfo(ctx context.Context, req *model.GetIslandInfoReq) (*model.GetIslandInfoRsp, error) {
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getIslandInfo))
	if err != nil {
		return nil, err
	}

	result := &model.GetIslandInfoRsp{}
	if err = tools.JSON.Unmarshal(resp.Result().(*model.OpenApiRpcRsp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
