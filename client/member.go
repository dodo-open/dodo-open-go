package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/errs"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetMemberList 取成员列表
func (c *client) GetMemberList(ctx context.Context, req *model.GetMemberListReq) (*model.GetMemberListRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getMemberListUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetMemberListRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMemberInfo 取成员信息
func (c *client) GetMemberInfo(ctx context.Context, req *model.GetMemberInfoReq) (*model.GetMemberInfoRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getMemberInfoUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetMemberInfoRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMemberRoleList 取成员身份组列表
func (c *client) GetMemberRoleList(ctx context.Context, req *model.GetMemberRoleListReq) ([]*model.RoleElement, error) {
	list := make([]*model.RoleElement, 0)

	if err := req.ValidParams(); err != nil {
		return list, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getMemberRoleListUri))
	if err != nil {
		return list, err
	}

	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &list); err != nil {
		return list, err
	}
	return list, nil
}

// SetMemberNick 设置成员昵称
func (c *client) SetMemberNick(ctx context.Context, req *model.SetMemberNickReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setMemberNickUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// SetMemberSilence 设置成员禁言，即不能在频道发布内容
func (c *client) SetMemberSilence(ctx context.Context, req *model.SetMemberSilenceReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setMemberSilenceUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}
