package model

import "errors"

// GetRoleListReq 取身份组列表 request
type GetRoleListReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
}

func (p *GetRoleListReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	return nil
}

// RoleElement 取身份组列表 list element
type RoleElement struct {
	RoleId   string `json:"roleId"`   // 身份组ID
	RoleName string `json:"roleName"` // 身份组名称
}

// AddRoleMemberReq 身份组新增成员 request
type AddRoleMemberReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
	DodoId   string `json:"dodoId" binding:"required"`   // DoDo号
	RoleId   string `json:"roleId" binding:"required"`   // 身份组ID
}

func (p *AddRoleMemberReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	if p.RoleId == "" {
		return errors.New("invalid parameter RoleId (empty detected)")
	}
	return nil
}

// RemoveRoleMemberReq 身份组移除成员 request
type RemoveRoleMemberReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
	DodoId   string `json:"dodoId" binding:"required"`   // DoDo号
	RoleId   string `json:"roleId" binding:"required"`   // 身份组ID
}

func (p *RemoveRoleMemberReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	if p.DodoId == "" {
		return errors.New("invalid parameter DodoId (empty detected)")
	}
	if p.RoleId == "" {
		return errors.New("invalid parameter RoleId (empty detected)")
	}
	return nil
}
