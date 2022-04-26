package main

import (
	"context"
	"fmt"
	"github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
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
		fmt.Printf("创建实例失败：%v", err)
		return
	}

	// 举例：发送频道文字消息
	content := fmt.Sprintf("example: send-channel-message, time: %s", time.Now().Format("2006-01-02 15:04:05"))
	resp, err := instance.SendChannelMessage(context.Background(), &model.SendChannelMessageReq{
		ChannelId:   "171204",
		MessageBody: &model.ChannelTextMessage{Content: content},
	})
	if err != nil {
		fmt.Printf("发送消息失败：%v", err)
		return
	}
	fmt.Printf("回报消息 ID：%v", resp.MessageId)
}
