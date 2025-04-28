package request

import "encoding/json"

type AlipayUserAgreementPageSignRequest struct {
	NotifyUrl  string
	BizContent AlipayUserAgreementPageSignBizContent
}

// AlipayUserAgreementPageSignBizContent alipay.user.agreement.page.sign(支付宝个人协议页面签约接口)
// https://opendocs.alipay.com/open/8bccfa0b_alipay.user.agreement.page.sign?scene=common&pathHash=725a0634
// 支付宝个人协议页面签约参数
type AlipayUserAgreementPageSignBizContent struct {
	PersonalProductCode string                 `json:"personal_product_code,omitempty"` //个人签约产品码
	SignScene           string                 `json:"sign_scene,omitempty"`            //【描述】请按当前接入的方式进行填充，且输入值必须为文档中的参数取值范围。 扫码或者短信页面签约需要拼装http的请求地址访问中间页面，钱包h5页面签约可直接拼接scheme的请求地址
	ExternalAgreementNo string                 `json:"external_agreement_no,omitempty"` //【描述】用户在商户网站的登录账号，
	ThirdPartyType      string                 `json:"third_party_type,omitempty"`      //【描述】签约第三方主体类型,默认为PARTNER。
	AccessParams        map[string]interface{} `json:"access_params,omitempty"`         //【描述】协议中标注的商户服务名称，由商户自定义。
}

// ToMap 转map
func (r *AlipayUserAgreementPageSignRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(r.BizContent)
	return map[string]string{
		"notify_url":  r.NotifyUrl,
		"biz_content": string(bizContent),
	}
}
