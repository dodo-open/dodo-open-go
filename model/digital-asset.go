package model

import "errors"

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
