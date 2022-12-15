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
