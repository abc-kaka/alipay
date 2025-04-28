package request

import "encoding/json"

type AlipayTradePayRequest struct {
	NotifyUrl  string
	BizContent AlipayTradePayBizContent
}

// 普通预授权冻结押金转支付接口说明：https://opendocs.alipay.com/open/02cdx8?scene=34&pathHash=dd2813f3
type AlipayTradePayBizContent struct {
	OutTradeNo      string                 `json:"out_trade_no,omitempty"`
	TotalAmount     float64                `json:"total_amount,omitempty"`
	Subject         string                 `json:"subject,omitempty"`
	ProductCode     string                 `json:"product_code,omitempty"`
	AuthNo          string                 `json:"auth_no,omitempty"`
	AuthConfirmMode string                 `json:"auth_confirm_mode,omitempty"`
	ExtendParams    map[string]interface{} `json:"extend_params,omitempty"`
	BusinessParams  map[string]interface{} `json:"business_params,omitempty"`

	// 直付通参数
	BuyerId     string                 `json:"buyer_id,omitempty"`     // 用户支付宝uid
	SellerId    string                 `json:"seller_id,omitempty"`    // 卖家支付宝uid
	SubMerchant map[string]interface{} `json:"sub_merchant,omitempty"` // 二级商户信息
	SettleInfo  map[string]interface{} `json:"settle_info,omitempty"`  // 结算详细信息
}

// ToMap 转map
func (r *AlipayTradePayRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(r.BizContent)
	return map[string]string{
		"notify_url":  r.NotifyUrl,
		"biz_content": string(bizContent),
	}
}
