package response

// AlipayFundAuthOperationDetailQueryResponse 资金授权操作查询接口-响应
type AlipayFundAuthOperationDetailQueryResponse struct {
	Code                    string `json:"code"`
	Msg                     string `json:"msg"`
	AuthNo                  string `json:"auth_no"`
	OutOrderNo              string `json:"out_order_no"`
	OrderStatus             string `json:"order_status"`
	TotalFreezeAmount       string `json:"total_freeze_amount"`
	RestAmount              string `json:"rest_amount"`
	TotalPayAmount          string `json:"total_pay_amount"`
	OrderTitle              string `json:"order_title"`
	PayerLogonId            string `json:"payer_logon_id"`
	PayerUserId             string `json:"payer_user_id"`
	PayerOpenId             string `json:"payer_open_id"`
	ExtraParam              string `json:"extra_param"`
	OperationId             string `json:"operation_id"`
	OutRequestNo            string `json:"out_request_no"`
	Amount                  string `json:"amount"`
	OperationType           string `json:"operation_type"`
	Status                  string `json:"status"`
	Remark                  string `json:"remark"`
	GmtCreate               string `json:"gmt_create"`
	GmtTrans                string `json:"gmt_trans"`
	PreAuthType             string `json:"pre_auth_type"`
	TransCurrency           string `json:"trans_currency"`
	TotalFreezeCreditAmount string `json:"total_freeze_credit_amount"`
	TotalFreezeFundAmount   string `json:"total_freeze_fund_amount"`
	TotalPayCreditAmount    string `json:"total_pay_credit_amount"`
	TotalPayFundAmount      string `json:"total_pay_fund_amount"`
	RestCreditAmount        string `json:"rest_credit_amount"`
	RestFundAmount          string `json:"rest_fund_amount"`
	CreditAmount            string `json:"credit_amount"`
	FundAmount              string `json:"fund_amount"`
	CreditMerchantExt       string `json:"credit_merchant_ext"`
}
