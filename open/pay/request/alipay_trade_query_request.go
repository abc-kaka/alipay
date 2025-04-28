package request

import "github.com/abc-kaka/alipay/common/req"

type AlipayTradeQueryRequest struct {
	req.BizContentRequest
	OutTradeNo   string   `json:"out_trade_no,omitempty"`
	TradeNo      string   `json:"trade_no,omitempty"`
	OrgPid       string   `json:"org_pid,omitempty"`
	QueryOptions []string `json:"query_options,omitempty"`
}
