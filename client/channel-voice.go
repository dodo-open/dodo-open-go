package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/errs"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetChannelVoiceMemberStatus 获取成员语音频道状态
func (c *client) GetChannelVoiceMemberStatus(ctx context.Context, req *model.GetChannelVoiceMemberStatusReq) (*model.GetChannelVoiceMemberStatusRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getChannelVoiceMemberStatusUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetChannelVoiceMemberStatusRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetChannelVoiceMemberMove 移动语音频道成员
func (c *client) SetChannelVoiceMemberMove(ctx context.Context, req *model.SetChannelVoiceMemberMoveReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setChannelVoiceMemberMoveUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// SetChannelVoiceMemberEdit 管理语音中的成员
func (c *client) SetChannelVoiceMemberEdit(ctx context.Context, req *model.SetChannelVoiceMemberEditReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setChannelVoiceMemberEditUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}
