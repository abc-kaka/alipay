package request

import (
	"github.com/abc-kaka/alipay/common/req"
)

type AlipayFundAuthOperationDetailQueryRequest struct {
	req.BizContentRequest
	OutRequestNo string `json:"out_request_no,omitempty"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
}
