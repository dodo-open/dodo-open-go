package model

import "errors"

// GetIslandListReq 获取群列表 request
type GetIslandListReq struct {
}

// IslandElement 获取群列表 list element
type IslandElement struct {
	IslandSourceId    string `json:"islandSourceId"`    // 群ID
	IslandName        string `json:"islandName"`        // 群名称
	CoverUrl          string `json:"coverUrl"`          // 群头像
	MemberCount       int    `json:"memberCount"`       // 成员数
	OnlineMemberCount int    `json:"onlineMemberCount"` // 在线成员数
	DefaultChannelId  string `json:"defaultChannelId"`  // 默认进入频道
	SystemChannelId   string `json:"systemChannelId"`   // 系统消息频道
}

// GetIslandInfoReq 获取群信息 request
type GetIslandInfoReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
}

func (p *GetIslandInfoReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	return nil
}

// GetIslandInfoRsp 获取群信息 response
type GetIslandInfoRsp struct {
	IslandElement

	Description string `json:"description"` // 群描述
}

type (
	// GetIslandLevelRankListReq 获取群等级排行榜 request
	GetIslandLevelRankListReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	}

	// GetIslandLevelRankElement 获取群等级排行榜 list element
	GetIslandLevelRankElement struct {
		DodoSourceId string `json:"dodoSourceId"` // DoDoID
		NickName     string `json:"nickName"`     // 群昵称
		Level        int    `json:"level"`        // 等级
		Rank         int    `json:"rank"`         // 排名，返回前100名
	}
)

func (p *GetIslandLevelRankListReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	return nil
}

type (
	// GetIslandMuteListReq 获取群禁言名单 request
	GetIslandMuteListReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
		PageSize       int    `json:"pageSize" binding:"required"`       // 页大小，最大100
		MaxId          uint64 `json:"maxId" binding:"required"`          // 上一页最大 ID 值，为提升分页查询性能，需要传入上一页查询记录中的最大 ID 值，首页请传 0
	}

	// GetIslandMuteListRsp 获取群禁言名单 response
	GetIslandMuteListRsp struct {
		MaxId uint64                  `json:"maxId"` // 最大 ID 值
		List  []*GetIslandMuteElement `json:"list"`  // 数据列表
	}

	// GetIslandMuteElement 获取群禁言名单 list element
	GetIslandMuteElement struct {
		DodoSourceId string `json:"dodoSourceId"` // DoDoID
	}
)

func (p *GetIslandMuteListReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.PageSize <= 0 {
		return errors.New("invalid parameter PageSize (PageSize must not less than 0)")
	}
	return nil
}

type (
	// GetIslandBanListReq 获取群封禁名单 request
	GetIslandBanListReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
		PageSize       int    `json:"pageSize" binding:"required"`       // 页大小，最大100
		MaxId          uint64 `json:"maxId" binding:"required"`          // 上一页最大 ID 值，为提升分页查询性能，需要传入上一页查询记录中的最大 ID 值，首页请传 0
	}

	// GetIslandBanListRsp 获取群封禁名单 response
	GetIslandBanListRsp struct {
		MaxId uint64                 `json:"maxId"` // 最大 ID 值
		List  []*GetIslandBanElement `json:"list"`  // 数据列表
	}

	// GetIslandBanElement 获取群封禁名单 list element
	GetIslandBanElement struct {
		DodoSourceId string `json:"dodoSourceId"` // DoDoID
	}
)

func (p *GetIslandBanListReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.PageSize <= 0 {
		return errors.New("invalid parameter PageSize (PageSize must not less than 0)")
	}
	return nil
}
