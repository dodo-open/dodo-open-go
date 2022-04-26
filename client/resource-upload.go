package client

import (
	"bytes"
	"context"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// UploadImageByBytes 上传图片资源（字节方式）
// you can use os.ReadFile (or other ways you want) to get the bytes of image file
func (c *client) UploadImageByBytes(ctx context.Context, req *model.UploadImageByBytesReq) (*model.UploadImageRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).
		SetFileReader("file", req.Filename, bytes.NewBuffer(req.Bytes)).
		Post(c.getApi(uploadImageUri))
	if err != nil {
		return nil, err
	}

	result := &model.UploadImageRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// UploadImageByPath 上传图片资源（路径方式）
// we recommend you to use absolute path to upload image file
func (c *client) UploadImageByPath(ctx context.Context, req *model.UploadImageByPathReq) (*model.UploadImageRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).
		SetFile("file", req.Path).
		Post(c.getApi(uploadImageUri))
	if err != nil {
		return nil, err
	}

	result := &model.UploadImageRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
