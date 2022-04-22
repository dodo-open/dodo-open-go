package model

import (
	jsoniter "github.com/json-iterator/go"
)

// OpenAPIRsp wrapped OpenAPI RPC response
type OpenAPIRsp struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    jsoniter.RawMessage `json:"data"`
}
