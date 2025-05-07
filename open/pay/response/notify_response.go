package response

type NotifyResponse struct {
	NotifyTime        string `json:"notify_time"`
	NotifyType        string `json:"notify_type"`
	NotifyId          string `json:"notify_id"`
	SignType          string `json:"sign_type"`
	Sign              string `json:"sign"`
	TradeNo           string `json:"trade_no"`
	AppId             string `json:"app_id"`
	AuthAppId         string `json:"auth_app_id"`
	OutTradeNo        string `json:"out_trade_no"`
	OutBizNo          string `json:"out_biz_no"`
	BuyerId           string `json:"buyer_id"`
	BuyerLogonId      string `json:"buyer_logon_id"`
	SellerId          string `json:"seller_id"`
	SellerEmail       string `json:"seller_email"`
	TradeStatus       string `json:"trade_status"`
	TotalAmount       string `json:"total_amount"`
	ReceiptAmount     string `json:"receipt_amount"`
	InvoiceAmount     string `json:"invoice_amount"`
	BuyerPayAmount    string `json:"buyer_pay_amount"`
	PointAmount       string `json:"point_amount"`
	RefundFee         string `json:"refund_fee"`
	SendBackFee       string `json:"send_back_fee"`
	Subject           string `json:"subject"`
	Body              string `json:"body"`
	GmtCreate         string `json:"gmt_create"`
	GmtPayment        string `json:"gmt_payment"`
	GmtRefund         string `json:"gmt_refund"`
	GmtClose          string `json:"gmt_close"`
	FundBillList      string `json:"fund_bill_list"`
	VoucherDetailList string `json:"voucher_detail_list"`
	BizSettleMode     string `json:"biz_settle_mode"`
}
