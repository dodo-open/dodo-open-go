package model

import (
	"errors"
)

// MemberElement 取成员列表 list element
type MemberElement struct {
	DodoId       string `json:"dodoId"`       // DoDo号
	NickName     string `json:"nickName"`     // 在群昵称
	AvatarUrl    string `json:"avatarUrl"`    // 头像
	JoinTime     string `json:"joinTime"`     // 加群时间
	Sex          int    `json:"sex"`          // 性别，-1：保密，0：女，1：男
	Level        int    `json:"level"`        // 等级
	IsBot        int    `json:"isBot"`        // 是否机器人，0：否，1：是
	OnlineDevice int    `json:"onlineDevice"` // 在线设备，0：无，1：PC在线，2：手机在线
	OnlineStatus int    `json:"onlineStatus"` // 在线状态，0：离线，1：在线，2：请勿打扰，3：离开
}

type (
	// GetMemberListReq 取成员列表 request
	GetMemberListReq struct {
		IslandId string `json:"islandId" binding:"required"` // 群号
		PageSize int    `json:"pageSize" binding:"required"` // 页大小，最大100
		MaxId    uint64 `json:"maxId" binding:"required"`    // 上一页最大ID值，为提升分页查询性能，需要传入上一页查询记录中的最大ID值，首页请传0
	}

	// GetMemberListRsp 取成员列表 response
	GetMemberListRsp struct {
		MaxId uint64           `json:"maxId"` // 最大 ID 值
		List  []*MemberElement `json:"list"`  // 列表
	}
)

func (p *GetMemberListReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		return errors.New("invalid parameter PageSize (0 < PageSize <= 100)")
	}
	return nil
}

type (
	// GetMemberInfoReq 取成员信息 request
	GetMemberInfoReq struct {
		IslandId string `json:"islandId" binding:"required"` // 群号
		DodoId   string `json:"dodoId" binding:"required"`   // DoDo号
	}

	// GetMemberInfoRsp 取成员信息 response
	GetMemberInfoRsp struct {
		MemberElement
		IslandId string `json:"islandId"` // 群号
	}
)

func (p *GetMemberInfoReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	return nil
}

// GetMemberRoleListReq 取成员身份组列表 request
type GetMemberRoleListReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
	DodoId   string `json:"dodoId" binding:"required"`   // DoDo号
}

func (p *GetMemberRoleListReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	return nil
}

// SetMemberNickReq 设置成员昵称 request
type SetMemberNickReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
	DodoId   string `json:"dodoId" binding:"required"`   // DoDo号
	NickName string `json:"nickName" binding:"required"` // 在群昵称
}

func (p *SetMemberNickReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	if p.NickName == "" {
		return errors.New("invalid parameter NickName (empty detected)")
	}
	return nil
}

// SetMemberSilenceReq 设置成员禁言 request
type SetMemberSilenceReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
	DodoId   string `json:"dodoId" binding:"required"`   // DoDo号
	Duration int64  `json:"duration" binding:"required"` // 禁言时长（单位：秒），最长 7 天
	Reason   string `json:"reason,omitempty"`            // 禁言原因
}

func (p *SetMemberSilenceReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	if p.Duration <= 0 || p.Duration > (7*24*60*60) {
		return errors.New("invalid parameter Duration (0 second < Duration <= 7 days)")
	}
	return nil
}

type (
	// GetMemberInviteInfoReq 取成员邀请信息 request
	GetMemberInviteInfoReq struct {
		IslandId string `json:"islandId" binding:"required"` // 群号
		DodoId   string `json:"dodoId" binding:"required"`   // DoDo号
	}

	// GetMemberInviteInfoRsp 取成员邀请信息 response
	GetMemberInviteInfoRsp struct {
		DodoId          string `json:"dodoId"`          // DoDo号
		NickName        string `json:"nickName"`        // 在群昵称
		InvitationCount int    `json:"invitationCount"` // 邀请人数
	}
)

func (p *GetMemberInviteInfoReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	return nil
}

type (
	// GetMemberUPowerchainInfoReq 取成员高能链数字藏品信息 request
	GetMemberUPowerchainInfoReq struct {
		IslandId string `json:"islandId" binding:"required"` // 群号
		DodoId   string `json:"dodoId" binding:"required"`   // DoDo号
		Issuer   string `json:"issuer" binding:"required"`   // 发行商
		Series   string `json:"series,omitempty"`            // 系列
	}

	// GetMemberUPowerchainInfoRsp 取成员高能链数字藏品信息 response
	GetMemberUPowerchainInfoRsp struct {
		DodoId           string `json:"dodoId"`           // DoDo号
		NickName         string `json:"nickName"`         // 在群昵称
		PersonalNickName string `json:"personalNickName"` // 个人昵称
		IsHaveIssuer     int    `json:"isHaveIssuer"`     // 是否拥有该发行方的NFT，0：否，1：是
		IsHaveSeries     int    `json:"isHaveSeries"`     // 是否拥有该系列的NFT，0：否，1：是
		NftCount         int    `json:"nftCount"`         // 拥有的该系列下NFT数量
	}
)

func (p *GetMemberUPowerchainInfoReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	if p.Issuer == "" {
		return errors.New("invalid parameter Issuer (empty detected)")
	}
	return nil
}
