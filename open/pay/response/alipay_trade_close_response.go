package response

type AlipayTradeCloseResponse struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}
