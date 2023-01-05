package model

import "errors"

// GetChannelVoiceMemberStatusReq 获取成员语音频道状态
// [https://open.imdodo.com/dev/api/channel-voice.html#%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98%E8%AF%AD%E9%9F%B3%E9%A2%91%E9%81%93%E7%8A%B6%E6%80%81]
type GetChannelVoiceMemberStatusReq struct {
	IslandSourceId string `json:"islandSourceId"`
	DodoSourceId   string `json:"dodoSourceId"`
}

func (p *GetChannelVoiceMemberStatusReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter islandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter dodoSourceId (empty detected)")
	}
	return nil
}

// GetChannelVoiceMemberStatusRsp 获取成员语音频道状态
// [https://open.imdodo.com/dev/api/channel-voice.html#%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98%E8%AF%AD%E9%9F%B3%E9%A2%91%E9%81%93%E7%8A%B6%E6%80%81]
type GetChannelVoiceMemberStatusRsp struct {
	ChannelId     string `json:"channelId"`
	MicStatus     int    `json:"micStatus"`
	SpkStatus     int    `json:"spkStatus"`
	MicSortStatus int    `json:"micSortStatus"`
}

// SetChannelVoiceMemberMoveReq 移动语音频道成员
// [https://open.imdodo.com/dev/api/channel-voice.html#%E7%A7%BB%E5%8A%A8%E8%AF%AD%E9%9F%B3%E9%A2%91%E9%81%93%E6%88%90%E5%91%98]
type SetChannelVoiceMemberMoveReq struct {
	IslandSourceId string `json:"islandSourceId"`
	DodoSourceId   string `json:"dodoSourceId"`
	ChannelId      string `json:"channelId"`
}

func (p *SetChannelVoiceMemberMoveReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter islandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter dodoSourceId (empty detected)")
	}
	if p.ChannelId == "" {
		return errors.New("invalid parameter channelId (empty detected)")
	}
	return nil
}

// SetChannelVoiceMemberEditReq 管理语音中的成员
// [https://open.imdodo.com/dev/api/channel-voice.html#%E7%AE%A1%E7%90%86%E8%AF%AD%E9%9F%B3%E4%B8%AD%E7%9A%84%E6%88%90%E5%91%98]
type SetChannelVoiceMemberEditReq struct {
	ChannelId    string `json:"channelId"`
	DodoSourceId string `json:"dodoSourceId"`
	OperateType  int    `json:"operateType"`
}

func (p *SetChannelVoiceMemberEditReq) ValidParams() error {
	if p.OperateType != 0 && p.OperateType != 1 && p.OperateType != 2 && p.OperateType != 3 {
		return errors.New("invalid parameter operateType (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter dodoSourceId (empty detected)")
	}
	if p.ChannelId == "" {
		return errors.New("invalid parameter channelId (empty detected)")
	}
	return nil
}
