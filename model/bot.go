package model

import "errors"

// GetBotInfoRsp 获取机器人信息 response
type GetBotInfoRsp struct {
	ClientId  string `json:"clientId"`  // 机器人唯一标识
	DodoId    string `json:"dodoId"`    // 机器人DoDo号
	NickName  string `json:"nickName"`  // 机器人昵称
	AvatarUrl string `json:"avatarUrl"` // 机器人图标
}

// SetBotLeaveIslandReq 机器人退群 request
type SetBotLeaveIslandReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
}

func (p *SetBotLeaveIslandReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	return nil
}
