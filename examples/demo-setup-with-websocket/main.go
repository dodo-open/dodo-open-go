package main

import (
	"fmt"
	"github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/tools"
	"github.com/dodo-open/dodo-open-go/websocket"
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

	ws, err := websocket.New(instance,
		// 允许自定义消息 channel 的队列长度
		websocket.WithMessageQueueSize(1024),
		// 设置自定义的消息处理器中间件
		websocket.WithMessageHandlers(&websocket.MessageHandlers{
			// 私聊事件处理器
			PersonalMessage: func(event *websocket.WSEventMessage, data *websocket.PersonalMessageEventBody) error {
				j, _ := tools.JSON.Marshal(data)
				fmt.Printf("事件：%s", string(j))
				return nil
			},
			// 频道消息事件处理器
			ChannelMessage: func(event *websocket.WSEventMessage, data *websocket.ChannelMessageEventBody) error {
				j, _ := tools.JSON.Marshal(data)
				fmt.Printf("事件：%s", string(j))
				return nil
			},
			// 消息反应事件处理器
			MessageReaction: func(event *websocket.WSEventMessage, data *websocket.MessageReactionEventBody) error {
				j, _ := tools.JSON.Marshal(data)
				fmt.Printf("事件：%s", string(j))
				return nil
			},
			// 用户入群事件处理器
			MemberJoin: func(event *websocket.WSEventMessage, data *websocket.MemberJoinEventBody) error {
				j, _ := tools.JSON.Marshal(data)
				fmt.Printf("事件：%s", string(j))
				return nil
			},
			// 用户离开群事件处理器
			MemberLeave: func(event *websocket.WSEventMessage, data *websocket.MemberLeaveEventBody) error {
				j, _ := tools.JSON.Marshal(data)
				fmt.Printf("事件：%s", string(j))
				return nil
			},
			// 未分类的纯文本事件处理器
			PlainTextHandler: func(event *websocket.WSEventMessage, message []byte) error {
				fmt.Printf("文本数据：%s\n", string(message))
				return nil
			},
			// 异常处理器
			ErrorHandler: func(err error) {
				fmt.Printf("(error) | %v", err)
			},
		}),
	)
	if err != nil {
		fmt.Printf("创建 WebSocket 实例失败：%v", err)
		return
	}

	if err := ws.Connect(); err != nil {
		fmt.Printf("连接 WebSocket 远程服务器失败：%v", err)
		return
	}

	// 调试阶段可以使用下面的代码
	// 注意，一旦启动监听，会将当前进程阻塞
	if err := ws.Listen(); err != nil {
		fmt.Printf("监听异常：%v", err)
		return
	}

	// 所以生产环境建议使用协程运行监听
	go func() {
		if err := ws.Listen(); err != nil {
			fmt.Printf("监听异常：%v", err)
			return
		}
	}()
}
