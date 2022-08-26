package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/errs"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetMemberList 获取成员列表
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

// GetMemberInfo 获取成员信息
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

// GetMemberRoleList 获取成员身份组列表
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

// GetMemberInviteInfo 获取成员邀请信息
func (c *client) GetMemberInviteInfo(ctx context.Context, req *model.GetMemberInviteInfoReq) (*model.GetMemberInviteInfoRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getMemberInviteInfo))
	if err != nil {
		return nil, err
	}

	result := &model.GetMemberInviteInfoRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetMemberNick 编辑成员群昵称
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

// MuteMember 禁言成员
func (c *client) MuteMember(ctx context.Context, req *model.MuteMemberReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(muteMemberUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// UnmuteMember 取消成员禁言
func (c *client) UnmuteMember(ctx context.Context, req *model.UnmuteMemberReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(unmuteMemberUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// BanMember 永久封禁成员
func (c *client) BanMember(ctx context.Context, req *model.BanMemberReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(banMemberUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}

// UnbanMember 取消成员永久封禁
func (c *client) UnbanMember(ctx context.Context, req *model.UnbanMemberReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(unbanMemberUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}
