<p align="center">
  <a href="https://open.imdodo.com">
    <img src="https://open.imdodo.com/hero.png" width="200" height="200" alt="dodo-open">
  </a>
</p>

<div align="center">

# dodo-open-go

_✨ 基于最新 GO 开发，内建事件对象解析、支持多实例、支持自定义事件处理中间件。 ✨_

  <a href="https://github.com/dodo-open/dodo-open-go/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/dodo-open/dodo-open-go" alt="license">
  </a>
  <a href="https://github.com/dodo-open/dodo-open-go/releases">
    <img src="https://img.shields.io/github/v/release/dodo-open/dodo-open-go?color=blueviolet&include_prereleases"
      alt="release">
  </a>
	
</div>

## 特性

- DoDo OpenAPI
- WebSocket
- 内建事件对象解析
- 支持多实例
- 支持自定义事件处理中间件

## 起步

```shell
go get github.com/dodo-open/dodo-open-go
```

如果你只需要使用`OpenAPI`功能，使用下面的代码创建`Bot`实例就可以了。

```go
clientId := "your-bot-client-id"
token := "your-bot-token"

// 下面的第三个参数，设定了 resty 的请求超时为 3 秒：
//
//     client.WithTimeout(time.Second*3)
//
instance, err := NewInstance(clientId, token, client.WithTimeout(time.Second*3))

// 获取你的 Bot 加入过的群的列表，可以使用下面的方法
list, err := instance.GetIslandList(context.Background())
```

## 上手 WebSocket

要创建支持 WebSocket 的实例，除了上面[起步](#起步)流程的代码，还需要使用下面的代码来启动 WebSocket 功能：

```go
// 创建 WebSocket 实例，它依赖 instance 对象，即上面创建的 Bot 实例
ws, err := websocket.New(instance)
if err != nil {
	t.Fatal(err)
}

// 主动连接到 WebSocket 服务器
if err = ws.Connect(); err != nil {
	t.Fatal(err)
}

// 开始监听事件消息
if err = ws.Listen(); err != nil {
	t.Fatal(err)
}
```

### 为什么需要我手动发起连接和开始监听事件？

本 SDK 支持两种注册消息处理器中间件的方式，分别是全局消息处理器，和实例级消息处理器。在启动 WebSocket 监听之前，你应该需要开发你自己的消息处理器，来消费我们的 WebSocket 服务器下发的消息。

#### 编写消息处理器

在编写消息处理器前，你可以查看源代码中 `websocket/handler.go` 这个文件，这里会告诉你如何定义一个消息处理器。

以定义「频道消息处理器」为例，可以参考下面的代码:

```go
channelMessageHandler := func(event *WSEventMessage, data *ChannelMessageEventBody) error {
    fmt.Printf("%v\n", data)
	return nil
}
```

#### 注册全局消息处理器

你可以使用 `websocket.RegisterHandlers(handlers ...interface{})` 来注册你自己的消息处理器，该方法会将全局消息处理器改成你定义的，会对所有的`Bot`实例生效。

#### 注册实例级消息处理器

在创建`Bot`实例时，你所使用的 `websocket.New()` 支持传入以 `With` 词缀开头的配置中间件，我们提供了 `websocket.WithMessageHandlers(handlers *MessageHandlers)` 方法来支持你注册实例级的消息处理器。

## 关于协作

我们欢迎社区参与协同开发工作，要想参与到 `dodo-open-go` 的开发工作中，你需要：

1. 会写 golang
2. 基本了解 go-resty
3. fork 本仓库
4. 在您自己的 fork 仓库中完成代码开发
5. 向上游主库（也就是 [`dodo-open/dodo-open-go`](https://github.com/dodo-open/dodo-open-go)）发起 Pull Request
6. 等候我们的审阅，或前往[DoDo 开发者社区-内测版 | DoDo 渡渡语音](https://imdodo.com/s/108015)敲一敲我们

需要注意，在 `dodo-open-go` 中，我们的一些 Restful API 封装，不一定完全使用来自开放平台文档所给的接口名称，例如发送消息 `SetChannelMessageSend`，在本 SDK 中被称作 `SendChannelMessage`。我们将大多数程序化的描述名称，改变成了英文口语上更容易读且朗朗上口的名称。所以当你需要确认一个 API 是否被他人封装，请用 URI 进行全局检索。

## TODO

- 或许我们需要重构封装事件数据的结构，当前`WSEventMessage.Data`和`EventData.EventBody`这两个属性都使用了`jsoniter.RawMessage`来封装。
- WebSocket 的重连流程不是很友好，当前如果需要断线重连，你必须关闭之前的所有资源，包括`conn`、`messageChan`、`closeChan`。
