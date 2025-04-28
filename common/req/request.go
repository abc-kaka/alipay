package req

import "encoding/json"

// Request 请求接口
type Request interface {
	ToMap(v any) map[string]string
}

type BizContentRequest struct{}

func (r *BizContentRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(v)
	return map[string]string{
		"biz_content": string(bizContent),
	}
}
