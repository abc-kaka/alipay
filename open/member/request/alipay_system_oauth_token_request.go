package request

// AlipaySystemOauthTokenRequest 换取授权访问令牌 - 参数
type AlipaySystemOauthTokenRequest struct {
	GrantType    string // 设置授权方式
	Code         string // 设置授权码
	RefreshToken string // 刷新令牌
}

// ToMap 转map
func (r *AlipaySystemOauthTokenRequest) ToMap(v any) map[string]string {
	return map[string]string{
		"grant_type":    r.GrantType,
		"code":          r.Code,
		"refresh_token": r.RefreshToken,
	}
}
