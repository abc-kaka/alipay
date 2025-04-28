package request

import "github.com/abc-kaka/alipay/common/req"

// AlipayOpenMiniOrderInstallmentCloseRequest 订单分期关闭 - 参数
type AlipayOpenMiniOrderInstallmentCloseRequest struct {
	req.BizContentRequest
	InstallmentOrderId    string `json:"installment_order_id,omitempty"`
	UserId                string `json:"user_id,omitempty"`
	OpenId                string `json:"open_id,omitempty"`
	OutInstallmentOrderId string `json:"out_installment_order_id,omitempty"`
	OutOrderId            string `json:"out_order_id,omitempty"`
	OrderId               string `json:"order_id,omitempty"`
}
