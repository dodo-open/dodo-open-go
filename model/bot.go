package model

// GetBotInfoRsp 取机器人信息 response
type GetBotInfoRsp struct {
	ClientId  string `json:"clientId"`  // 机器人唯一标识
	DodoId    string `json:"dodoId"`    // 机器人DoDo号
	NickName  string `json:"nickName"`  // 机器人昵称
	AvatarUrl string `json:"avatarUrl"` // 机器人图标
}

// SetBotLeaveIslandReq 置机器人群退出 request
type SetBotLeaveIslandReq struct {
	IslandId string `json:"islandId"` // 群号
}
