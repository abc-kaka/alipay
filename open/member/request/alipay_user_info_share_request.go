package request

// AlipayUserInfoShareRequest 支付宝会员授权信息查询接口 - 参数
type AlipayUserInfoShareRequest struct {
	AuthToken  string
	BizContent string
}

// ToMap 转map
func (r *AlipayUserInfoShareRequest) ToMap(v any) map[string]string {
	return map[string]string{
		"auth_token":  r.AuthToken,
		"biz_content": r.BizContent,
	}
}
