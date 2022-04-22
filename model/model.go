package model

import (
	jsoniter "github.com/json-iterator/go"
)

// OpenApiRpcRsp wrapped OpenAPI RPC response
type OpenApiRpcRsp struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    jsoniter.RawMessage `json:"data"`
}
