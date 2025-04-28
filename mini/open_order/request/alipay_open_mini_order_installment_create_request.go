package request

import (
	"github.com/abc-kaka/alipay/common/req"
)

// AlipayOpenMiniOrderInstallmentCreateRequest 订单分期创建 - 参数
type AlipayOpenMiniOrderInstallmentCreateRequest struct {
	req.BizContentRequest
	OrderId               string              `json:"order_id,omitempty"`
	OutOrderId            string              `json:"out_order_id,omitempty"`
	UserId                string              `json:"user_id,omitempty"`
	OpenId                string              `json:"open_id,omitempty"`
	OutInstallmentOrderId string              `json:"out_installment_order_id,omitempty"`
	TradeNo               string              `json:"trade_no,omitempty"`
	IsFinishPerformance   bool                `json:"is_finish_performance,omitempty"`
	Type                  string              `json:"type,omitempty"`
	InstallmentPrice      string              `json:"installment_price,omitempty"`
	InstallmentNoType     string              `json:"installment_no_type,omitempty"`
	InstallmentNo         string              `json:"installment_no,omitempty"`
	PayChannel            string              `json:"pay_channel,omitempty"`
	StageNo               int                 `json:"stage_no,omitempty"`
	IsSyncPay             bool                `json:"is_sync_pay,omitempty"`
	PayTime               string              `json:"pay_time,omitempty"`
	InstallmentNoInfoList []InstallmentNoInfo `json:"installment_no_info_list,omitempty"`
}

type InstallmentNoInfo struct {
	StageNo       int    `json:"stage_no,omitempty"`
	InstallmentNo string `json:"installment_no,omitempty"`
}
