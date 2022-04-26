package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetWebsocketConnection 获取 Websocket 连接
func (c *client) GetWebsocketConnection(ctx context.Context) (*model.GetWebsocketConnectionRsp, error) {
	resp, err := c.request(ctx).SetBody(&model.GetWebsocketConnectionReq{}).Post(c.getApi(getWebsocketConnectionUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetWebsocketConnectionRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
