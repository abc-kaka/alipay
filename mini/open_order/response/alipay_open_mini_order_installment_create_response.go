package response

type AlipayOpenMiniOrderInstallmentCreateResponse struct {
	Code               string `json:"code"`
	Msg                string `json:"msg"`
	InstallmentOrderId string `json:"installment_order_id"`
}
