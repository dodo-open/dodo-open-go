package model

import (
	"errors"
)

// MemberElement 获取成员列表 list element
type MemberElement struct {
	DodoSourceId     string `json:"dodoSourceId"`     // DoDoID
	NickName         string `json:"nickName"`         // 在群昵称
	PersonalNickName string `json:"personalNickName"` // DoDo昵称
	AvatarUrl        string `json:"avatarUrl"`        // 头像
	JoinTime         string `json:"joinTime"`         // 加群时间
	Sex              int    `json:"sex"`              // 性别，-1：保密，0：女，1：男
	Level            int    `json:"level"`            // 等级
	IsBot            int    `json:"isBot"`            // 是否机器人，0：否，1：是
	OnlineDevice     int    `json:"onlineDevice"`     // 在线设备，0：无，1：PC在线，2：手机在线
	OnlineStatus     int    `json:"onlineStatus"`     // 在线状态，0：离线，1：在线，2：请勿打扰，3：离开
}

type (
	// GetMemberListReq 获取成员列表 request
	GetMemberListReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
		PageSize       int    `json:"pageSize" binding:"required"`       // 页大小，最大100
		MaxId          uint64 `json:"maxId" binding:"required"`          // 上一页最大ID值，为提升分页查询性能，需要传入上一页查询记录中的最大ID值，首页请传0
	}

	// GetMemberListRsp 获取成员列表 response
	GetMemberListRsp struct {
		MaxId uint64           `json:"maxId"` // 最大 ID 值
		List  []*MemberElement `json:"list"`  // 列表
	}
)

func (p *GetMemberListReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		return errors.New("invalid parameter PageSize (0 < PageSize <= 100)")
	}
	return nil
}

type (
	// GetMemberInfoReq 获取成员信息 request
	GetMemberInfoReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
		DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
	}

	// GetMemberInfoRsp 获取成员信息 response
	GetMemberInfoRsp struct {
		MemberElement
		IslandSourceId string `json:"islandSourceId"` // 群ID
	}
)

func (p *GetMemberInfoReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	return nil
}

// GetMemberRoleListReq 获取成员身份组列表 request
type GetMemberRoleListReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
}

func (p *GetMemberRoleListReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	return nil
}

// SetMemberNickReq 编辑成员群昵称 request
type SetMemberNickReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
	NickName       string `json:"nickName" binding:"required"`       // 在群昵称
}

func (p *SetMemberNickReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	if p.NickName == "" {
		return errors.New("invalid parameter NickName (empty detected)")
	}
	return nil
}

// MuteMemberReq 禁言成员 request
type MuteMemberReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
	Duration       int64  `json:"duration" binding:"required"`       // 禁言时长（单位：秒），最长 7 天
	Reason         string `json:"reason,omitempty"`                  // 禁言原因
}

func (p *MuteMemberReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	if p.Duration <= 0 || p.Duration > (7*24*60*60) {
		return errors.New("invalid parameter Duration (0 second < Duration <= 7 days)")
	}
	return nil
}

// UnmuteMemberReq 取消禁言成员 request
type UnmuteMemberReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
}

func (p *UnmuteMemberReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	return nil
}

// BanMemberReq 永久封禁成员 request
type BanMemberReq struct {
	IslandSourceId  string `json:"islandSourceId" binding:"required"` // 群ID
	DodoSourceId    string `json:"dodoSourceId" binding:"required"`   // DoDoID
	NoticeChannelId string `json:"noticeChannelId,omitempty"`         // 通知频道ID
	Reason          string `json:"reason,omitempty"`                  // 封禁理由，理由不能大于64个字符或32个汉字
}

func (p *BanMemberReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	return nil
}

// UnbanMemberReq 取消成员永久封禁 request
type UnbanMemberReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
}

func (p *UnbanMemberReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	return nil
}

type (
	// GetMemberInviteInfoReq 获取成员邀请信息 request
	GetMemberInviteInfoReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
		DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
	}

	// GetMemberInviteInfoRsp 获取成员邀请信息 response
	GetMemberInviteInfoRsp struct {
		DodoSourceId    string `json:"dodoSourceId"`    // DoDoID
		NickName        string `json:"nickName"`        // 在群昵称
		InvitationCount int    `json:"invitationCount"` // 邀请人数
	}
)

func (p *GetMemberInviteInfoReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	return nil
}
