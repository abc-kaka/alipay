package response

type AlipayTradeQueryResponse struct {
	Code                  string              `json:"code"`
	Msg                   string              `json:"msg"`
	TradeNo               string              `json:"trade_no"`
	OutTradeNo            string              `json:"out_trade_no"`
	BuyerLogonId          string              `json:"buyer_logon_id"`
	TradeStatus           string              `json:"trade_status"`
	AdditionalStatus      string              `json:"additional_status"`
	TotalAmount           string              `json:"total_amount"`
	TransCurrency         string              `json:"trans_currency"`
	SettleCurrency        string              `json:"settle_currency"`
	SettleAmount          string              `json:"settle_amount"`
	PayCurrency           string              `json:"pay_currency"`
	PayAmount             string              `json:"pay_amount"`
	SettleTransRate       string              `json:"settle_trans_rate"`
	TransPayRate          string              `json:"trans_pay_rate"`
	BuyerPayAmount        string              `json:"buyer_pay_amount"`
	PointAmount           string              `json:"point_amount"`
	InvoiceAmount         string              `json:"invoice_amount"`
	SendPayDate           string              `json:"send_pay_date"`
	ReceiptAmount         string              `json:"receipt_amount"`
	StoreId               string              `json:"store_id"`
	TerminalId            string              `json:"terminal_id"`
	FundBillList          []FundBill          `json:"fund_bill_list"`
	StoreName             string              `json:"store_name"`
	BuyerUserId           string              `json:"buyer_user_id"`
	BuyerOpenId           string              `json:"buyer_open_id"`
	IndustrySepcDetailGov string              `json:"industry_sepc_detail_gov"`
	IndustrySepcDetailAcc string              `json:"industry_sepc_detail_acc"`
	ChargeAmount          string              `json:"charge_amount"`
	ChargeFlags           string              `json:"charge_flags"`
	SettlementId          string              `json:"settlement_id"`
	TradeSettleInfo       TradeSettleInfo     `json:"trade_settle_info"`
	AuthTradePayMode      string              `json:"auth_trade_pay_mode"`
	DiscountAmount        string              `json:"discount_amount"`
	Subject               string              `json:"subject"`
	Body                  string              `json:"body"`
	AlipaySubMerchantId   string              `json:"alipay_sub_merchant_id"`
	ExtInfos              string              `json:"ext_infos"`
	PassbackParams        string              `json:"passback_params"`
	BuyerUserType         string              `json:"buyer_user_type"`
	MdiscountAmount       string              `json:"mdiscount_amount"`
	HbFqPayInfo           HbFqPayInfo         `json:"hb_fq_pay_info"`
	CreditPayMode         string              `json:"credit_pay_mode"`
	CreditBizOrderId      string              `json:"credit_biz_order_id"`
	HybAmount             string              `json:"hyb_amount"`
	BkagentRespInfo       BkagentRespInfo     `json:"bkagent_resp_info"`
	ChargeInfoList        []ChargeInfo        `json:"charge_info_list"`
	BizSettleMode         string              `json:"biz_settle_mode"`
	ReqGoodsDetail        []ReqGoodsDetail    `json:"req_goods_detail"`
	FulfillmentDetailList []FulfillmentDetail `json:"fulfillment_detail_list"`
	PeriodScene           string              `json:"period_scene"`
	AsyncPayApplyStatus   string              `json:"async_pay_apply_status"`
	CashierType           string              `json:"cashier_type"`
	TapPayInfo            TapPayInfo          `json:"tap_pay_info"`
	PreAuthPayAmount      string              `json:"pre_auth_pay_amount"`
}

type FundBill struct {
	FundChannel string `json:"fund_channel"`
	Amount      string `json:"amount"`
	RealAmount  string `json:"real_amount"`
}

type TradeSettleInfo struct {
	TradeSettleDetailList []TradeSettleDetailList `json:"trade_settle_detail_list"`
	TradeUnsettledAmount  string                  `json:"trade_unsettled_amount"`
}

type TradeSettleDetailList struct {
	OperationType     string `json:"operation_type"`
	OperationSerialNo string `json:"operation_serial_no"`
	OperationDt       string `json:"operation_dt"`
	TransOut          string `json:"trans_out"`
	TransIn           string `json:"trans_in"`
	Amount            string `json:"amount"`
	OriTransOut       string `json:"ori_trans_out"`
	OriTransIn        string `json:"ori_trans_in"`
}

type HbFqPayInfo struct {
	UserInstallNum string `json:"user_install_num"`
}

type BkagentRespInfo struct {
	BindtrxId        string `json:"bindtrx_id"`
	BindclrissrId    string `json:"bindclrissr_id"`
	BindpyeracctbkId string `json:"bindpyeracctbk_id"`
	BkpyeruserCode   string `json:"bkpyeruser_code"`
	EstterLocation   string `json:"estter_location"`
}

type ChargeInfo struct {
	ChargeFee               string         `json:"charge_fee"`
	OriginalChargeFee       string         `json:"original_charge_fee"`
	SwitchFeeRate           string         `json:"switch_fee_rate"`
	IsRatingOnTradeReceiver string         `json:"is_rating_on_trade_receiver"`
	IsRatingOnSwitch        string         `json:"is_rating_on_switch"`
	ChargeType              string         `json:"charge_type"`
	SubFeeDetailList        []SubFeeDetail `json:"sub_fee_detail_list"`
}

type SubFeeDetail struct {
	ChargeFee         string `json:"charge_fee"`
	OriginalChargeFee string `json:"original_charge_fee"`
	SwitchFeeRate     string `json:"switch_fee_rate"`
}

type ReqGoodsDetail struct {
	GoodsId        string `json:"goods_id"`
	AlipayGoodsId  string `json:"alipay_goods_id"`
	GoodsName      string `json:"goods_name"`
	Quantity       int    `json:"quantity"`
	Price          string `json:"price"`
	GoodsCategory  string `json:"goods_category"`
	CategoriesTree string `json:"categories_tree"`
	Body           string `json:"body"`
	ShowUrl        string `json:"show_url"`
	OutItemId      string `json:"out_item_id"`
	OutSkuId       string `json:"out_sku_id"`
}

type FulfillmentDetail struct {
	FulfillmentAmount string `json:"fulfillment_amount"`
	OutRequestNo      string `json:"out_request_no"`
	GmtPayment        string `json:"gmt_payment"`
}

type TapPayInfo struct {
	PaymentMediumType string `json:"payment_medium_type"`
}
