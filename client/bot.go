package client

import (
	"context"
	"dodo-open-go/errs"
	"dodo-open-go/model"
	"dodo-open-go/tools"
)

// GetBotInfo 取机器人信息
func (c *client) GetBotInfo(ctx context.Context) (*model.GetBotInfoRsp, error) {
	resp, err := c.request(ctx).Post(c.getApi(getBotInfoUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetBotInfoRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetBotIslandLeave 置机器人群退出
func (c *client) SetBotIslandLeave(ctx context.Context, req *model.SetBotLeaveIslandReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setBotIslandLeaveUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}
