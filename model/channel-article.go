package model

import "errors"

// SetChannelArticleAddReq 发布帖子
// [https://open.imdodo.com/dev/api/channel-article.html#%E5%8F%91%E5%B8%83%E5%B8%96%E5%AD%90]
type SetChannelArticleAddReq struct {
	ChannelId string `json:"channelId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageUrl  string `json:"imageUrl"`
}

func (p *SetChannelArticleAddReq) ValidParams() error {
	if p.ChannelId == "" {
		return errors.New("invalid parameter channelId (empty detected)")
	}
	if p.Title == "" {
		return errors.New("invalid parameter title (empty detected)")
	}
	if p.Content == "" && p.ImageUrl == "" {
		return errors.New("invalid parameter content/imageUrl (empty detected)")
	}

	return nil
}

// SetChannelArticleAddRsp 发布帖子
// [https://open.imdodo.com/dev/api/channel-article.html#%E5%8F%91%E5%B8%83%E5%B8%96%E5%AD%90]
type SetChannelArticleAddRsp struct {
	ArticleId string `json:"articleId"`
}

// SetChannelArticleRemoveReq 删除帖子评论回复
// [https://open.imdodo.com/dev/api/channel-article.html#%E5%88%A0%E9%99%A4%E5%B8%96%E5%AD%90%E8%AF%84%E8%AE%BA%E5%9B%9E%E5%A4%8D]
type SetChannelArticleRemoveReq struct {
	ChannelId string `json:"channelId"`
	Type      int    `json:"type"`
	Id        string `json:"id"`
}

func (p *SetChannelArticleRemoveReq) ValidParams() error {
	if p.ChannelId == "" {
		return errors.New("invalid parameter channelId (empty detected)")
	}
	if p.Id == "" {
		return errors.New("invalid parameter id (empty detected)")
	}
	if p.Type != 1 && p.Type != 2 && p.Type != 3 {
		return errors.New("invalid parameter type (empty detected)")
	}
	return nil
}
