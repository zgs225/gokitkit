package gokitkit

// APIResponse 通用的接口返回接口
type APIResponse struct {
	Code      int         `json:"code"`
	Timestamp int64       `json:"timestamp"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Error     error       `json:"-"`
}
