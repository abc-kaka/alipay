package response

type AlipaySystemOauthTokenResponse struct {
	UserId       string `json:"user_id"`
	OpenId       string `json:"open_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	ReExpiresIn  uint64 `json:"re_expires_in"`
	AuthStart    string `json:"auth_start"`
}

type AlipayUserInfoShareResponse struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	UserId   string `json:"user_id"`
	Avatar   string `json:"avatar"`
	City     string `json:"city"`
	NickName string `json:"nick_name"`
	Province string `json:"province"`
	Gender   string `json:"gender"`
}
