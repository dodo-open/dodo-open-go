package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetGiftAccount 获取群收入
func (c *client) GetGiftAccount(ctx context.Context, req *model.GetGiftAccountReq) (*model.GetGiftAccountRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getGiftAccountUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetGiftAccountRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGiftShareRatioInfo 获取成员分成管理
func (c *client) GetGiftShareRatioInfo(ctx context.Context, req *model.GetGiftShareRatioInfoReq) (*model.GetGiftShareRatioInfoRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getGiftShareRatioInfoUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetGiftShareRatioInfoRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGiftList 获取内容礼物列表
func (c *client) GetGiftList(ctx context.Context, req *model.GetGiftListReq) ([]*model.GetGiftListRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getGiftListUri))
	if err != nil {
		return nil, err
	}

	result := make([]*model.GetGiftListRsp, 0)
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGiftMemberList 获取内容礼物内成员列表
func (c *client) GetGiftMemberList(ctx context.Context, req *model.GetGiftMemberListReq) (*model.GetGiftMemberListRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getGiftMemberListUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetGiftMemberListRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetGiftGrossValueList 获取内容礼物总值列表
func (c *client) GetGiftGrossValueList(ctx context.Context, req *model.GetGiftGrossValueListReq) (*model.GetGiftGrossValueListRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getGiftGrossValueListUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetGiftGrossValueListRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
