package request

import "github.com/abc-kaka/alipay/common/req"

type AlipayTradeCloseRequest struct {
	req.BizContentRequest
	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OperatorId string `json:"operator_id,omitempty"`
}
