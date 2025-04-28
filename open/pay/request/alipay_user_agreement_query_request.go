package request

import "encoding/json"

// AlipayUserAgreementQueryBizContent  参数描述见：https://opendocs.alipay.com/open/3dab71bc_alipay.user.agreement.query?pathHash=9a0c5949
// AlipayUserAgreementQueryRequest 支付宝个人代扣协议查询接口 - 参数
type AlipayUserAgreementQueryBizContent struct {
	//【描述】支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） ，如果传了该参数，其他参数会被忽略
	//【示例值】20170322450983769228
	AgreementNo string `json:"agreement_no,omitempty"`
}
type AlipayUserAgreementQueryRequest struct {
	BizContent AlipayUserAgreementQueryBizContent
}

// ToMap 转map
func (r *AlipayUserAgreementQueryRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(r.BizContent)
	return map[string]string{
		"biz_content": string(bizContent),
	}
}
