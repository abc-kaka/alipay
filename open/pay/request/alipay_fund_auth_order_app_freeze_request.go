package request

import (
	"encoding/json"
)

// AlipayFundAuthOrderAppFreezeRequest 线上资金授权冻结 - 参数
// 参数描述见：https://opendocs.alipay.com/open/064jhe?pathHash=629fa9a5
type AlipayFundAuthOrderAppFreezeRequest struct {
	NotifyUrl  string
	BizContent AlipayFundAuthOrderAppFreezeBizContent
}

type AlipayFundAuthOrderAppFreezeBizContent struct {
	NotifyUrl          string `json:"notify_url,omitempty"`
	OutOrderNo         string `json:"out_order_no,omitempty"`
	OutRequestNo       string `json:"out_request_no,omitempty"`
	OrderTitle         string `json:"order_title,omitempty"`
	Amount             string `json:"amount,omitempty"`
	ProductCode        string `json:"product_code,omitempty"`
	DepositProductMode string `json:"deposit_product_mode,omitempty"`
	PostPayments       string `json:"post_payments,omitempty"`
	PayeeUserId        string `json:"payee_user_id,omitempty"`
	PayeeLogonId       string `json:"payee_logon_id,omitempty"`
	TimeoutExpress     string `json:"timeout_express,omitempty"`
	EnablePayChannels  string `json:"enable_pay_channels,omitempty"`
	DisablePayChannels string `json:"disable_pay_channels,omitempty"`
	IdentityParams     string `json:"identity_params,omitempty"`
	ExtraParam         string `json:"extra_param,omitempty"`
	BusinessParams     string `json:"business_params,omitempty"`
}

// ToMap 转map
func (r *AlipayFundAuthOrderAppFreezeRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(r.BizContent)
	return map[string]string{
		"notify_url":  r.NotifyUrl,
		"biz_content": string(bizContent),
	}
}
