package receive_message

import (
	"github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
	"github.com/dodo-open/dodo-open-go/websocket"
	"testing"
	"time"
)

// setup instance before running test
func setup(t *testing.T) client.Client {
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
		t.Fatalf("创建实例失败：%v", err)
		return nil
	}
	return instance
}

// after setup websocket client
func after(t *testing.T, ws websocket.Client) {
	if err := ws.Connect(); err != nil {
		t.Fatal(err)
	}
	if err := ws.Listen(); err != nil {
		t.Fatal(err)
	}
}

// Test_ReceiveChannelMessage 接收到频道消息
func Test_ReceiveChannelMessage(t *testing.T) {
	instance := setup(t)

	// 定义消息处理器
	handlers := &websocket.MessageHandlers{
		// data 是已经解析好的消息事件数据对象，但其中的 MessageBody 是 jsoniter.RawMessage
		ChannelMessage: func(event *websocket.WSEventMessage, data *websocket.ChannelMessageEventBody) error {
			// 可以简单的通过 MessageType，根据不同的消息类型，解析到不同的结构
			switch data.MessageType {
			case model.TextMsg:
				// 文本消息
				messageBody := &model.TextMessage{}
				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
					return err
				}
				t.Logf("文本消息：%s", messageBody.Content)
			case model.ImageMsg:
				// 图片消息
				messageBody := &model.ImageMessage{}
				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
					return err
				}
				t.Logf("图片消息：%+v", messageBody)
			case model.VideoMsg:
				// 视频消息
				messageBody := &model.VideoMessage{}
				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
					return err
				}
				t.Logf("视频消息：%+v", messageBody)
			case model.FileMsg:
				// 文件消息
				messageBody := &model.FileMessage{}
				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
					return err
				}
				t.Logf("文件消息：%+v", messageBody)
			}
			return nil
		},
	}

	ws, _ := websocket.New(instance,
		websocket.WithMessageQueueSize(128),
		websocket.WithMessageHandlers(handlers),
	)
	after(t, ws)
}

// Test_ReceiveDirectMessage 接收私聊消息
func Test_ReceiveDirectMessage(t *testing.T) {
	instance := setup(t)

	// 定义消息处理器
	handlers := &websocket.MessageHandlers{
		// data 是已经解析好的消息事件数据对象，但其中的 MessageBody 是 jsoniter.RawMessage
		PersonalMessage: func(event *websocket.WSEventMessage, data *websocket.PersonalMessageEventBody) error {
			// 可以简单的通过 MessageType，根据不同的消息类型，解析到不同的结构
			switch data.MessageType {
			case model.TextMsg:
				// 文本消息
				messageBody := &model.TextMessage{}
				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
					return err
				}
				t.Logf("文本消息：%s", messageBody.Content)
			case model.ImageMsg:
				// 图片消息
				messageBody := &model.ImageMessage{}
				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
					return err
				}
				t.Logf("图片消息：%+v", messageBody)
			case model.VideoMsg:
				// 视频消息
				messageBody := &model.VideoMessage{}
				if err := tools.JSON.Unmarshal(data.MessageBody, &messageBody); err != nil {
					return err
				}
				t.Logf("视频消息：%+v", messageBody)
			}
			return nil
		},
	}

	ws, _ := websocket.New(instance,
		websocket.WithMessageQueueSize(128),
		websocket.WithMessageHandlers(handlers),
	)
	after(t, ws)
}
