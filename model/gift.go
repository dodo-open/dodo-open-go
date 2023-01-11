package model

import "errors"

// GetGiftAccountReq 获取群收入 请求
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E7%BE%A4%E6%94%B6%E5%85%A5]
type GetGiftAccountReq struct {
	IslandSourceId string `json:"islandSourceId"`
}

func (p *GetGiftAccountReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter islandSourceId (empty detected)")
	}
	return nil
}

// GetGiftAccountRsp 获取群收入 响应
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E7%BE%A4%E6%94%B6%E5%85%A5]
type GetGiftAccountRsp struct {
	TotalIncome        float64 `json:"totalIncome"`
	SettlableIncome    float64 `json:"settlableIncome"`
	TransferableIncome float64 `json:"transferableIncome"`
}

// GetGiftShareRatioInfoReq 获取成员分成管理 请求
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98%E5%88%86%E6%88%90%E7%AE%A1%E7%90%86]
type GetGiftShareRatioInfoReq struct {
	IslandSourceId string `json:"islandSourceId"`
}

func (p *GetGiftShareRatioInfoReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter islandSourceId (empty detected)")
	}
	return nil
}

// GetGiftShareRatioInfoRsp 获取成员分成管理 响应
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98%E5%88%86%E6%88%90%E7%AE%A1%E7%90%86]
type GetGiftShareRatioInfoRsp struct {
	DefaultRatio struct {
		IslandRatio   float64 `json:"islandRatio"`
		UserRatio     float64 `json:"userRatio"`
		PlatformRatio float64 `json:"platformRatio"`
	} `json:"defaultRatio"`
	RoleRatioList []struct {
		RoleId        string  `json:"roleId"`
		RoleName      string  `json:"roleName"`
		IslandRatio   float64 `json:"islandRatio"`
		UserRatio     float64 `json:"userRatio"`
		PlatformRatio float64 `json:"platformRatio"`
	} `json:"roleRatioList"`
}

// GetGiftListReq 获取内容礼物列表 请求
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E5%86%85%E5%AE%B9%E7%A4%BC%E7%89%A9%E5%88%97%E8%A1%A8]
type GetGiftListReq struct {
	TargetType int    `json:"targetType"`
	TargetId   string `json:"targetId"`
}

func (p *GetGiftListReq) ValidParams() error {
	if p.TargetType != 1 && p.TargetType != 2 {
		return errors.New("invalid parameter targetType (empty detected)")
	}
	if p.TargetId == "" {
		return errors.New("invalid parameter targetId (empty detected)")
	}
	return nil
}

// GetGiftListRsp 获取内容礼物列表 响应
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E5%86%85%E5%AE%B9%E7%A4%BC%E7%89%A9%E5%88%97%E8%A1%A8]
type GetGiftListRsp struct {
	GiftId          string  `json:"giftId"`
	GiftCount       int     `json:"giftCount"`
	GiftTotalAmount float64 `json:"giftTotalAmount"`
}

// GetGiftMemberListReq 获取内容礼物内成员列表 请求
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E5%86%85%E5%AE%B9%E7%A4%BC%E7%89%A9%E5%86%85%E6%88%90%E5%91%98%E5%88%97%E8%A1%A8]
type GetGiftMemberListReq struct {
	TargetType int    `json:"targetType"`
	TargetId   string `json:"targetId"`
	GiftId     string `json:"giftId"`
	PageSize   int    `json:"pageSize"`
	MaxId      int    `json:"maxId"`
}

func (p *GetGiftMemberListReq) ValidParams() error {
	if p.TargetType != 1 && p.TargetType != 2 {
		return errors.New("invalid parameter targetType (empty detected)")
	}
	if p.TargetId == "" {
		return errors.New("invalid parameter targetId (empty detected)")
	}
	if p.GiftId == "" {
		return errors.New("invalid parameter giftId (empty detected)")
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		return errors.New("invalid parameter pageSize ")
	}
	if p.MaxId < 0 {
		return errors.New("invalid parameter maxId ")
	}
	return nil
}

// GetGiftMemberListRsp 获取内容礼物内成员列表 响应
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E5%86%85%E5%AE%B9%E7%A4%BC%E7%89%A9%E5%86%85%E6%88%90%E5%91%98%E5%88%97%E8%A1%A8]
type GetGiftMemberListRsp struct {
	List []struct {
		DodoSourceId string `json:"dodoSourceId"`
		NickName     string `json:"nickName"`
		GiftCount    int    `json:"giftCount"`
	} `json:"list"`
	MaxId int `json:"maxId"`
}

// GetGiftGrossValueListReq 获取内容礼物总值列表 请求
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E5%86%85%E5%AE%B9%E7%A4%BC%E7%89%A9%E6%80%BB%E5%80%BC%E5%88%97%E8%A1%A8]
type GetGiftGrossValueListReq struct {
	TargetType int    `json:"targetType"`
	TargetId   string `json:"targetId"`
	PageSize   int    `json:"pageSize"`
	MaxId      int    `json:"maxId"`
}

func (p *GetGiftGrossValueListReq) ValidParams() error {
	if p.TargetType != 1 && p.TargetType != 2 {
		return errors.New("invalid parameter targetType (empty detected)")
	}
	if p.TargetId == "" {
		return errors.New("invalid parameter targetId (empty detected)")
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		return errors.New("invalid parameter pageSize ")
	}
	if p.MaxId < 0 {
		return errors.New("invalid parameter maxId ")
	}
	return nil
}

// GetGiftGrossValueListRsp 获取内容礼物总值列表 响应
// [https://open.imdodo.com/dev/api/gift.html#%E8%8E%B7%E5%8F%96%E5%86%85%E5%AE%B9%E7%A4%BC%E7%89%A9%E6%80%BB%E5%80%BC%E5%88%97%E8%A1%A8]
type GetGiftGrossValueListRsp struct {
	List []struct {
		DodoSourceId    string  `json:"dodoSourceId"`
		NickName        string  `json:"nickName"`
		GiftTotalAmount float64 `json:"giftTotalAmount"`
	} `json:"list"`
	MaxId int `json:"maxId"`
}
