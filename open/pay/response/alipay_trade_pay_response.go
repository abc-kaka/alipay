package response

// AlipayTradePayResponse 支付宝支付响应
type AlipayTradePayResponse struct {
	Code          string         `json:"code"`
	Msg           string         `json:"msg"`
	TradeNo       string         `json:"trade_no"`
	OutTradeNo    string         `json:"out_trade_no"`
	BuyerLogonId  string         `json:"rest_amount"`
	TotalAmount   string         `json:"total_pay_amount"`
	ReceiptAmount string         `json:"receipt_amount"`
	GmtPayment    string         `json:"gmt_payment"`
	FundBillList  []FundBillList `json:"fund_bill_list"`
	BuyerUserId   string         `json:"buyer_user_id"`
	BuyerOpenId   string         `json:"buyer_open_id"`
}
type FundBillList struct {
	FundChannel string `json:"fund_channel,omitempty"`
	Amount      string `json:"amount,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
}
