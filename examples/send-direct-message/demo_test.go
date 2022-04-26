package send_direct_message

import (
	"context"
	"fmt"
	"github.com/dodo-open/dodo-open-go/client"
	"github.com/dodo-open/dodo-open-go/model"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// 由于服务器有请求限制，这里的范例将拆成两个方法，以测试的形式展现

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

// Test_SendTextDM 举例：发送文字消息
func Test_SendTextDM(t *testing.T) {
	instance := setup(t)

	content := fmt.Sprintf("example: send-direct-message, time: %s", time.Now().Format("2006-01-02 15:04:05"))
	sendTextDmResp, err := instance.SendDirectMessage(context.Background(), &model.SendDirectMessageReq{
		DodoId:      "5464471",
		MessageBody: &model.TextMessage{Content: content},
	})
	if err != nil {
		t.Fatalf("发送消息失败：%v", err)
		return
	}
	t.Logf("回报消息 ID：%v", sendTextDmResp.MessageId)
}

func Test_SendImageDM(t *testing.T) {
	instance := setup(t)

	// Step 1. 读取文件
	abs, _ := filepath.Abs("../dodo.png")
	bytes, err := os.ReadFile(abs)
	if err != nil {
		t.Fatalf("读取文件失败：%v", err)
		return
	}
	// Step 2. 上传图片资源，获取 CDN 链接和图片宽高
	resourceResp, err := instance.UploadImageByBytes(context.Background(), &model.UploadImageByBytesReq{
		Filename: "dodo.png",
		Bytes:    bytes,
	})
	if err != nil {
		t.Fatalf("上传资源失败：%v", err)
		return
	}
	// Step 3. 发送消息
	sendImageDmResp, err := instance.SendDirectMessage(context.Background(), &model.SendDirectMessageReq{
		DodoId: "5464471",
		MessageBody: &model.ImageMessage{
			Url:        resourceResp.Url,
			Width:      resourceResp.Width,
			Height:     resourceResp.Height,
			IsOriginal: 0,
		},
	})
	if err != nil {
		t.Fatalf("发送消息失败：%v", err)
		return
	}
	t.Logf("回报消息 ID：%v", sendImageDmResp.MessageId)
}
