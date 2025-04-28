package request

import "encoding/json"

type AlipayFundAuthOrderUnfreezeRequest struct {
	NotifyUrl  string
	BizContent AlipayFundAuthOrderUnfreezeBizContent
}

// 普通预授权冻结押金转支付接口说明：https://opendocs.alipay.com/open/02cdx8?scene=34&pathHash=dd2813f3

type AlipayFundAuthOrderUnfreezeBizContent struct {
	AuthNo       string  `json:"auth_no,omitempty"`
	OutRequestNo string  `json:"out_request_no,omitempty"`
	Amount       float64 `json:"amount,omitempty"`
	Remark       string  `json:"remark,omitempty"`
	ExtraParam   string  `json:"extra_param,omitempty"`
}

// ToMap 转map
func (r *AlipayFundAuthOrderUnfreezeRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(r.BizContent)
	return map[string]string{
		"notify_url":  r.NotifyUrl,
		"biz_content": string(bizContent),
	}
}
