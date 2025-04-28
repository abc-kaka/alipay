package response

type AlipayOpenMiniOrderCreateResponse struct {
	OrderId             string             `json:"order_id"`
	OutOrderId          string             `json:"out_order_id"`
	CustomerDisplayText string             `json:"customer_display_text"`
	PayInfoResponse     PayInfoResponse    `json:"pay_info_response"`
	CreditInfoResponse  CreditInfoResponse `json:"credit_info_response"`
	ServiceType         string             `json:"service_type"`
}

type PayInfoResponse struct {
	NoPayCloseTime string `json:"no_pay_close_time"`
	TradeNo        string `json:"trade_no"`
}

type CreditInfoResponse struct {
	CreditPageLink string `json:"credit_page_link"`
}
