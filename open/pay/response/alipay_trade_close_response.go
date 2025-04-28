package response

type AlipayTradeCloseResponse struct {
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}
