package request

import "encoding/json"

type AlipayTradeCreateRequest struct {
	NotifyUrl  string
	BizContent AlipayTradeCreateBizContent
}

// AlipayTradeCreateBizContent alipay.trade.create(统一收单交易创建接口)
// https://opendocs.alipay.com/mini/6039ed0c_alipay.trade.create?pathHash=779dc517&ref=api&scene=de4d6a1e0c6e423b9eefa7c3a6dcb7a5
// 统一收单交易创建接口 请求参数
type AlipayTradeCreateBizContent struct {
	OutTradeNo         string                 `json:"out_trade_no,omitempty"`        // 商户订单号
	ProductCode        string                 `json:"product_code,omitempty"`        // 产品码
	OpAppId            string                 `json:"op_app_id,omitempty"`           // 商户小程序ID
	TotalAmount        float64                `json:"total_amount,omitempty"`        // 订单总金额
	ExtendParams       *ExtendParams          `json:"extend_params,omitempty"`       // 业务扩展参数
	DiscountableAmount string                 `json:"discountable_amount,omitempty"` // 可打折金额
	Subject            string                 `json:"subject,omitempty"`             // 订单标题
	Body               string                 `json:"body,omitempty"`                // 订单描述
	BuyerId            string                 `json:"buyer_id,omitempty"`            // 买家支付宝用户号
	BuyerOpenId        string                 `json:"buyer_open_id,omitempty"`       // 买家支付宝用户号
	StoreId            string                 `json:"store_id,omitempty"`            // 商户门店编号
	SubMerchant        map[string]interface{} `json:"sub_merchant,omitempty"`        // 二级商户信息
	SettleInfo         map[string]interface{} `json:"settle_info,omitempty"`         // 结算详细信息
	EnablePayChannels  string                 `json:"enable_pay_channels,omitempty"` //指定支付渠道。 用户只能使用指定的渠道进行支付，多个渠道以逗号分割。
}

type ExtendParams struct {
	SysServiceProviderId  string `json:"sys_service_provider_id,omitempty"`  //系统商编号
	HbFqNum               string `json:"hb_fq_num,omitempty"`                //使用花呗分期要进行的分期数
	HbFqSellerPercent     string `json:"hb_fq_seller_percent,omitempty"`     //使用花呗分期需要卖家承担的手续费比例的百分值，传入100代表100%
	TradeComponentOrderId string `json:"trade_component_order_id,omitempty"` //公域商品交易业务订单ID
	EnablePayChannels     string `json:"enable_pay_channels,omitempty"`      // 指定支付渠道
	BankMemo              string `json:"bank_memo,omitempty"`                // 直付通结算说明
}

// ToMap 转map
func (r *AlipayTradeCreateRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(r.BizContent)
	return map[string]string{
		"notify_url":  r.NotifyUrl,
		"biz_content": string(bizContent),
	}
}
