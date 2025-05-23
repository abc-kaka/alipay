package response

// AlipayTradeCreateResponse (统一收单交易创建接口)-响应
type AlipayTradeCreateResponse struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}
