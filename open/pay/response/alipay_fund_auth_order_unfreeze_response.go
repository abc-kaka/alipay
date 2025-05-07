package response

// AlipayFundAuthOrderUnfreezeResponse 资金授权解冻接口
type AlipayFundAuthOrderUnfreezeResponse struct {
	Code         string `json:"code"`
	Msg          string `json:"msg"`
	AuthNo       string `json:"auth_no"`
	OperationId  string `json:"operation_id"`
	OutRequestNo string `json:"out_request_no"`
	Amount       string `json:"amount"`
	Status       string `json:"status"`
	OutOrderNo   string `json:"out_order_no"`
	GmtTrans     string `json:"gmt_trans"`
	CreditAmount string `json:"credit_amount"`
	FundAmount   string `json:"fund_amount"`
}
