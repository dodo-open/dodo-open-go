package model

// GetIslandListReq 取群列表 request
type GetIslandListReq struct {
}

// IslandElement 取群列表 list element
type IslandElement struct {
	IslandId         string `json:"islandId"`         // 群号
	IslandName       string `json:"islandName"`       // 群名称
	CoverUrl         string `json:"coverUrl"`         // 群头像
	DefaultChannelId int    `json:"defaultChannelId"` // 默认进入频道
	SystemChannelId  int    `json:"systemChannelId"`  // 系统消息频道
}

// GetIslandInfoReq 取群信息 request
type GetIslandInfoReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
}

// GetIslandInfoRsp 取群信息 response
type GetIslandInfoRsp struct {
	IslandElement

	Description string `json:"description"` // 群描述
}
