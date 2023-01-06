package client

import (
	"context"
	"github.com/dodo-open/dodo-open-go/errs"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// SetChannelArticleAdd 发布帖子
func (c *client) SetChannelArticleAdd(ctx context.Context, req *model.SetChannelArticleAddReq) (*model.SetChannelArticleAddRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setChannelArticleAddUri))
	if err != nil {
		return nil, err
	}

	result := &model.SetChannelArticleAddRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// SetChannelArticleRemove 删除帖子评论回复
func (c *client) SetChannelArticleRemove(ctx context.Context, req *model.SetChannelArticleRemoveReq) (bool, error) {
	if err := req.ValidParams(); err != nil {
		return false, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setChannelArticleRemoveUri))
	if err != nil {
		return false, err
	}

	result := c.unmarshalResult(resp)
	if result.Status != 0 {
		return false, errs.New(result.Status, result.Message)
	}
	return true, nil
}
