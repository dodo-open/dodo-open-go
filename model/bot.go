package model

import "errors"

// GetBotInfoRsp 获取机器人信息 response
type GetBotInfoRsp struct {
	ClientId     string `json:"clientId"`     // 机器人唯一标识
	DodoSourceId string `json:"dodoSourceId"` // 机器人DoDoID
	NickName     string `json:"nickName"`     // 机器人昵称
	AvatarUrl    string `json:"avatarUrl"`    // 机器人图标
}

// SetBotLeaveIslandReq 机器人退群 request
type SetBotLeaveIslandReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
}

func (p *SetBotLeaveIslandReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	return nil
}

// GetBotInviteListReq 获取机器人邀请列表
// [https://open.imdodo.com/dev/api/bot.html#%E8%8E%B7%E5%8F%96%E6%9C%BA%E5%99%A8%E4%BA%BA%E9%82%80%E8%AF%B7%E5%88%97%E8%A1%A8]
type GetBotInviteListReq struct {
	PageSize int   `json:"pageSize"`
	MaxId    int64 `json:"maxId"`
}

func (p *GetBotInviteListReq) ValidParams() error {
	if p.PageSize <= 0 || p.PageSize > 100 {
		return errors.New("invalid parameter pageSize")
	}
	if p.MaxId < 0 {
		return errors.New("invalid parameter maxId")
	}
	return nil
}

// GetBotInviteListRsp 获取机器人邀请列表
// [https://open.imdodo.com/dev/api/bot.html#%E8%8E%B7%E5%8F%96%E6%9C%BA%E5%99%A8%E4%BA%BA%E9%82%80%E8%AF%B7%E5%88%97%E8%A1%A8]
type GetBotInviteListRsp struct {
	MaxId int64 `json:"maxId"`
	List  []struct {
		DodoSourceId string `json:"dodoSourceId"`
		NickName     string `json:"nickName"`
		AvatarUrl    string `json:"avatarUrl"`
	} `json:"list"`
}

// SetBotInviteAddReq  添加成员到机器人邀请列表 请求
// [https://open.imdodo.com/dev/api/bot.html#%E6%B7%BB%E5%8A%A0%E6%88%90%E5%91%98%E5%88%B0%E6%9C%BA%E5%99%A8%E4%BA%BA%E9%82%80%E8%AF%B7%E5%88%97%E8%A1%A8]
type SetBotInviteAddReq struct {
	DodoSourceId string `json:"dodoSourceId"`
}

func (p *SetBotInviteAddReq) ValidParams() error {
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter dodoSourceId")
	}
	return nil
}

// SetBotInviteRemoveReq  移除成员出机器人邀请列表 请求
// [https://open.imdodo.com/dev/api/bot.html#%E7%A7%BB%E9%99%A4%E6%88%90%E5%91%98%E5%87%BA%E6%9C%BA%E5%99%A8%E4%BA%BA%E9%82%80%E8%AF%B7%E5%88%97%E8%A1%A8]
type SetBotInviteRemoveReq struct {
	DodoSourceId string `json:"dodoSourceId"`
}

func (p *SetBotInviteRemoveReq) ValidParams() error {
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter dodoSourceId")
	}
	return nil
}
