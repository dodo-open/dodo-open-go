package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/errs"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetBotInfo 获取机器人信息
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

// SetBotIslandLeave 机器人退群
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

// GetBotInviteList 获取机器人邀请列表
func (c *client) GetBotInviteList(ctx context.Context, req *model.GetBotInviteListReq) (*model.GetBotInviteListRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getBotInviteListUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetBotInviteListRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetBotInviteAdd 添加成员到机器人邀请列表
func (c *client) SetBotInviteAdd(ctx context.Context, req *model.SetBotInviteAddReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setBotInviteAddUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// SetBotInviteRemove 移除成员出机器人邀请列表
func (c *client) SetBotInviteRemove(ctx context.Context, req *model.SetBotInviteRemoveReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setBotInviteRemoveUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}
