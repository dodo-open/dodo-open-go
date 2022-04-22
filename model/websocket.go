package model

// GetWebsocketConnectionReq 获取 Websocket 连接 request
type GetWebsocketConnectionReq struct {
}

// GetWebsocketConnectionRsp 获取 Websocket 连接 response
type GetWebsocketConnectionRsp struct {
	Endpoint string `json:"endpoint"` // 连接节点
}
