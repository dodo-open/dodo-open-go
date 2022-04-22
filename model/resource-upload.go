package model

import "errors"

// UploadImageByBytesReq 上传图片资源（字节方式） request
type UploadImageByBytesReq struct {
	Filename string `binding:"required"` // 文件名
	Bytes    []byte `binding:"required"` // 文件字节数组
}

func (p *UploadImageByBytesReq) ValidParams() error {
	if p.Filename == "" {
		return errors.New("invalid parameter Filename (empty detected)")
	}
	if p.Bytes == nil || len(p.Bytes) == 0 {
		return errors.New("invalid parameter DataBytes (empty detected)")
	}
	return nil
}

// UploadImageByPathReq 上传图片资源（路径方式） request
type UploadImageByPathReq struct {
	Path string `binding:"required"` // 文件绝对路径
}

func (p *UploadImageByPathReq) ValidParams() error {
	if p.Path == "" {
		return errors.New("invalid parameter Path (empty detected)")
	}
	return nil
}

// UploadImageRsp 上传图片资源 response
type UploadImageRsp struct {
	Url    string `json:"url"`    // 图片链接
	Width  int    `json:"width"`  // 图片宽度
	Height int    `json:"height"` // 图片高度
}
