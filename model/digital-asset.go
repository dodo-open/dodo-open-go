package model

import "errors"

type (
	// GetMemberNFTStatusReq 获取成员数字藏品判断 request
	GetMemberNFTStatusReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
		DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
		Platform       string `json:"platform" binding:"required"`       // 数藏平台，upower：高能链，ubanquan：优版权，metamask：Opensea
		Issuer         string `json:"issuer,omitempty"`                  // 发行方，若填写了系列，则发行方必填
		Series         string `json:"series,omitempty"`                  // 系列
	}

	// GetMemberNFTStatusRsp 获取成员数字藏品判断 response
	GetMemberNFTStatusRsp struct {
		IsBandPlatform int `json:"isBandPlatform"` // 是否已绑定该数藏平台，0：否，1：是
		IsHaveIssuer   int `json:"isHaveIssuer"`   // 是否拥有该发行方的NFT，0：否，1：是
		IsHaveSeries   int `json:"isHaveSeries"`   // 是否拥有该系列的NFT，0：否，1：是
	}
)

func (p *GetMemberNFTStatusReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	if p.Platform == "" {
		return errors.New("invalid parameter Platform (empty detected)")
	}
	if p.Series != "" && p.Issuer == "" {
		return errors.New("invalid parameter Issuer (empty detected while presents Series)")
	}
	return nil
}

type (
	// GetMemberUPowerchainInfoReq 取成员高能链数字藏品信息 request
	GetMemberUPowerchainInfoReq struct {
		IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
		DodoSourceId   string `json:"dodoSourceId" binding:"required"`   // DoDoID
		Issuer         string `json:"issuer" binding:"required"`         // 发行商
		Series         string `json:"series,omitempty"`                  // 系列
	}

	// GetMemberUPowerchainInfoRsp 取成员高能链数字藏品信息 response
	GetMemberUPowerchainInfoRsp struct {
		DodoSourceId     string `json:"dodoSourceId"`     // DoDoID
		NickName         string `json:"nickName"`         // 在群昵称
		PersonalNickName string `json:"personalNickName"` // 个人昵称
		IsHaveIssuer     int    `json:"isHaveIssuer"`     // 是否拥有该发行方的NFT，0：否，1：是
		IsHaveSeries     int    `json:"isHaveSeries"`     // 是否拥有该系列的NFT，0：否，1：是
		NftCount         int    `json:"nftCount"`         // 拥有的该系列下NFT数量
	}
)

func (p *GetMemberUPowerchainInfoReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter DodoSourceId (empty detected)")
	}
	if p.Issuer == "" {
		return errors.New("invalid parameter Issuer (empty detected)")
	}
	return nil
}
