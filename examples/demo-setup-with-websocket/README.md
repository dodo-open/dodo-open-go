# demo-setup-with-websocket

该示例代码演示了如何创建 Bot 实例，并创建 WebSocket 实例监听事件消息。

## 关于 WebSocket 的配置流程

由于我们 SDK 有「允许创建多个实例」的特性，以及「需要自定 WebSocket 消息处理器」的要求，当你创建 Bot 实例后，SDK 不会主动帮你启动 WebSocket 服务。

你需要按照下面的四步来启动 WebSocket 服务：

1. 调用`websocket.New(instance)`创建 WebSocket 实例，其中入参是你在先前创建的 Bot 实例对象；
2. 注册消息处理器（虽然不注册消息处理器也不影响运行，但所有的事件消息都会被忽略）；
3. 调用`Connect()`连接到 WebSocket 服务器；
4. 以协程的方式，调用`Listen()`启动监听。

## 关于监听方法`Listen()`

`Listen()`方法会阻塞当前进程，所以生产环境下不建议直接调用，而是放在协程运行。

## 关于消息处理器的注册时机

虽然我们支持动态注册实例级消息处理器，但为了运行安全，我们建议在启动连接前就注册好消息处理器。
