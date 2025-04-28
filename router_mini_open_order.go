package alipay

import (
	"github.com/abc-kaka/alipay/mini/open_order/request"
	"github.com/abc-kaka/alipay/mini/open_order/response"
)

// MiniOpenOrderRouter 小程序 - 交易组件 - 路由
type MiniOpenOrderRouter struct {
	client Client
}

func NewMiniOpenOrderRouter(client Client) *MiniOpenOrderRouter {
	return &MiniOpenOrderRouter{client: client}
}

// AlipayOpenMiniOrderDeliveryModify 履约状态变更接口
// https://opendocs.alipay.com/mini/efc03b89_alipay.open.mini.order.delivery.modify
func (r *MiniOpenOrderRouter) AlipayOpenMiniOrderDeliveryModify(request request.AlipayOpenMiniOrderDeliveryModifyRequest) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.open.mini.order.delivery.modify", &request, nil)
	return
}

// AlipayOpenMiniOrderModify 订单修改
// https://opendocs.alipay.com/mini/5440e982_alipay.open.mini.order.modify
func (r *MiniOpenOrderRouter) AlipayOpenMiniOrderModify(request request.AlipayOpenMiniOrderModifyRequest) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.open.mini.order.modify", &request, nil)
	return
}

// AlipayOpenMiniOrderInstallmentCreate 订单分期
// https://opendocs.alipay.com/mini/2cce2a35_alipay.open.mini.order.installment.create
func (r *MiniOpenOrderRouter) AlipayOpenMiniOrderInstallmentCreate(request request.AlipayOpenMiniOrderInstallmentCreateRequest, resp *response.AlipayOpenMiniOrderInstallmentCreateResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.open.mini.order.installment.create", &request, &resp)
	return
}

// AlipayOpenMiniOrderInstallmentClose 订单分期关闭
// https://opendocs.alipay.com/mini/1a9fb015_alipay.open.mini.order.installment.close
func (r *MiniOpenOrderRouter) AlipayOpenMiniOrderInstallmentClose(request request.AlipayOpenMiniOrderInstallmentCloseRequest, resp *response.AlipayOpenMiniOrderInstallmentCreateResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.open.mini.order.installment.close", &request, &resp)
	return
}
