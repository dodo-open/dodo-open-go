package main

import (
	"context"
	"fmt"
	"github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 创建实例
	clientId := "在开放平台创建的 Bot 的 ClientID"
	token := "在开放平台创建的 Bot 的 Token"
	instance, err := client.New(clientId, token,
		// 请求超时是可以自定义的
		client.WithTimeout(time.Second*3),
		// 可以设置为 Debug 模式（未来 SDK 会针对 Debug 模式输出相应的日志，现在仅仅只是预留）
		client.WithDebugMode(false),
	)
	if err != nil {
		fmt.Printf("创建实例失败：%v\n", err)
		return
	}

	// ================================================================================

	// 举例：发送频道文字消息
	content := fmt.Sprintf("example: send-channel-message, time: %s", time.Now().Format("2006-01-02 15:04:05"))
	sendTextResp, err := instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
		ChannelId:   "171204",
		MessageBody: &model.TextMessage{Content: content},
	})
	if err != nil {
		fmt.Printf("发送消息失败：%v\n", err)
		return
	}
	fmt.Printf("回报消息 ID：%v\n", sendTextResp.MessageId)

	// ================================================================================

	// 举例：发送频道图片消息
	// Step 1. 读取文件
	abs, _ := filepath.Abs("./dodo.png")
	bytes, err := os.ReadFile(abs)
	if err != nil {
		fmt.Printf("读取文件失败：%v\n", err)
		return
	}
	// Step 2. 上传图片资源，获取 CDN 链接和图片宽高
	resourceResp, err := instance.UploadImageByBytes(context.Background(), &model.UploadImageByBytesReq{
		Filename: "dodo.png",
		Bytes:    bytes,
	})
	if err != nil {
		fmt.Printf("上传资源失败：%v\n", err)
		return
	}
	// Step 3. 发送消息
	sendImageResp, err := instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
		ChannelId: "171204",
		MessageBody: &model.ImageMessage{
			Url:        resourceResp.Url,
			Width:      resourceResp.Width,
			Height:     resourceResp.Height,
			IsOriginal: 0,
		},
	})
	if err != nil {
		fmt.Printf("发送消息失败：%v\n", err)
		return
	}
	fmt.Printf("回报消息 ID：%v\n", sendImageResp.MessageId)
}
