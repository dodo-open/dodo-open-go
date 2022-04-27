package main

import (
	"context"
	"fmt"
	"github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/tools"
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

	// 举例：请求 Bot 的基础信息
	if info, err := instance.GetBotInfo(context.Background()); err != nil {
		fmt.Printf("获取 Bot 信息失败：%v\n", err)
		return
	} else {
		j, _ := tools.JSON.MarshalIndent(info, "", "    ")
		fmt.Printf("Bot 信息：%s\n", string(j))
	}

	// 举例：请求 Bot 加入过的群的列表
	if list, err := instance.GetIslandList(context.Background()); err != nil {
		fmt.Printf("获取群列表失败：%v\n", err)
		return
	} else {
		j, _ := tools.JSON.MarshalIndent(list, "", "    ")
		fmt.Printf("群列表：%s\n", string(j))
	}
}
