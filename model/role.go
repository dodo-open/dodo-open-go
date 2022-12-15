package model

import "errors"

// GetRoleListReq 获取身份组列表 request
type GetRoleListReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
}

func (p *GetRoleListReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	return nil
}

// RoleElement 身份组数据对象 list element
type RoleElement struct {
	RoleId     string `json:"roleId"`     // 身份组ID
	RoleName   string `json:"roleName"`   // 身份组名称
	RoleColor  string `json:"roleColor"`  // 身份组颜色，例：#ffffff
	Position   int    `json:"position"`   // 身份组排序位置
	Permission string `json:"permission"` // 身份组权限值，16进制
}

type (
	// CreateRoleReq 创建身份组 request
	CreateRoleReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
		RoleName       string `json:"roleName,omitempty"`                // 身份组名称，非必传，不传时默认使用`新的身份组`，不能大于32个字符或16个汉字
		RoleColor      string `json:"roleColor,omitempty"`               // 身份组颜色，非必传，不传时默认使用`#333333`，16进制HEX格式颜色码
		Position       int    `json:"position,omitempty"`                // 身份组排序位置，非必传，不传时默认为1，不可传比机器人身份组大的排序值
		Permission     string `json:"permission,omitempty"`              // 身份组权限值（16进制），非必传，不传时默认为0
	}

	// CreateRoleRsp 创建身份组 response
	CreateRoleRsp struct {
		RoleId string `json:"roleId"`
	}
)

func (p *CreateRoleReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	return nil
}

// EditRoleReq 编辑身份组 request
type EditRoleReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	RoleId         string `json:"roleId" binding:"required"`         // 身份组ID
	RoleName       string `json:"roleName,omitempty"`                // 身份组名称，非必传，不传时默认不改动，不能大于32个字符或16个汉字
	RoleColor      string `json:"roleColor,omitempty"`               // 身份组颜色，非必传，不传时默认不改动，16进制HEX格式颜色码
	Position       int    `json:"position,omitempty"`                // 身份组排序位置，非必传，不传时默认不改动，不可传比机器人身份组大的排序值
	Permission     string `json:"permission,omitempty"`              // 身份组权限值（16进制），非必传，不传时默认不改动
}

func (p *EditRoleReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.RoleId == "" {
		return errors.New("invalid parameter RoleId (empty detected)")
	}
	return nil
}

// RemoveRoleReq 删除身份组 request
type RemoveRoleReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	RoleId         string `json:"roleId" binding:"required"`         // 身份组ID
}

func (p *RemoveRoleReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.RoleId == "" {
		return errors.New("invalid parameter RoleId (empty detected)")
	}
	return nil
}

// AddRoleMemberReq 赋予成员身份组 request
type AddRoleMemberReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
	RoleId         string `json:"roleId" binding:"required"`         // 身份组ID
}

func (p *AddRoleMemberReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	if p.RoleId == "" {
		return errors.New("invalid parameter RoleId (empty detected)")
	}
	return nil
}

// RemoveRoleMemberReq 取消成员身份组 request
type RemoveRoleMemberReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
	DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
	RoleId         string `json:"roleId" binding:"required"`         // 身份组ID
}

func (p *RemoveRoleMemberReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	if p.RoleId == "" {
		return errors.New("invalid parameter RoleId (empty detected)")
	}
	return nil
}
