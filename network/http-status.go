package network

import "net/http"

const (
	OpenAPIStatusOK = 0
)

var successStatuses = map[int]bool{
	http.StatusOK:        true,
	http.StatusNoContent: true,
}

// IsSuccessResponse check the HTTP response code, if it is a success status code, returns true
func IsSuccessResponse(code int) bool {
	if _, ok := successStatuses[code]; ok {
		return true
	}
	return false
}
