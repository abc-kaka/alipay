package request

import "encoding/json"

type AlipayTradeAppPayRequest struct {
	NotifyUrl  string
	BizContent AlipayTradeAppPayBizContent
}

type AlipayTradeAppPayBizContent struct {
	OutTradeNo          string               `json:"out_trade_no,omitempty"`
	TotalAmount         string               `json:"total_amount,omitempty"`
	Subject             string               `json:"subject,omitempty"`
	ProductCode         string               `json:"product_code,omitempty"`
	TimeExpire          string               `json:"time_expire,omitempty"`
	AgreementSignParams *AgreementSignParams `json:"agreement_sign_params,omitempty"`
}

type AgreementSignParams struct {
	AccessParams        *AccessParams     `json:"access_params,omitempty"`
	PeriodRuleParams    *PeriodRuleParams `json:"period_rule_params,omitempty"`
	SignNotifyUrl       string            `json:"sign_notify_url,omitempty"`
	ExternalLogonId     string            `json:"external_logon_id,omitempty"`
	PersonalProductCode string            `json:"personal_product_code,omitempty"`
	ExternalAgreementNo string            `json:"external_agreement_no,omitempty"`
	ProductCode         string            `json:"product_code,omitempty"`
	SignScene           string            `json:"sign_scene,omitempty"`
	EffectTime          string            `json:"effect_time,omitempty"`
}

type AccessParams struct {
	Channel string `json:"channel,omitempty"`
}

type PeriodRuleParams struct {
	Period        int    `json:"period,omitempty"`
	TotalAmount   string `json:"total_amount,omitempty"`
	ExecuteTime   string `json:"execute_time,omitempty"`
	SingleAmount  string `json:"single_amount,omitempty"`
	TotalPayments int    `json:"total_payments,omitempty"`
	PeriodType    string `json:"period_type,omitempty"`
}

// ToMap 转map
func (r *AlipayTradeAppPayRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(r.BizContent)
	return map[string]string{
		"notify_url":  r.NotifyUrl,
		"biz_content": string(bizContent),
	}
}
