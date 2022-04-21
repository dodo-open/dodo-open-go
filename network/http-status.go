package network

import "net/http"

var successStatuses = map[int]bool{
	http.StatusOK:        true,
	http.StatusNoContent: true,
}

// IsSuccessResponse 判断响应的状态码是否为成功的状态码
func IsSuccessResponse(code int) bool {
	if _, ok := successStatuses[code]; ok {
		return true
	}
	return false
}
