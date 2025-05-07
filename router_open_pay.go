package alipay

import (
	"github.com/abc-kaka/alipay/open/pay/request"
	"github.com/abc-kaka/alipay/open/pay/response"
)

// OpenPayRouter H5&移动APP - 支付产品 - 路由
type OpenPayRouter struct {
	client Client
}

// NewMiniPayRouter 创建H5&移动APP - 支付产品 - 路由
func NewMiniPayRouter(client Client) *OpenPayRouter {
	return &OpenPayRouter{client: client}
}

// VerifyNotify 回调校验
// https://opendocs.alipay.com/open/194/103296
func (r *OpenPayRouter) VerifyNotify(request request.VerifyNotifyRequest, response *response.NotifyResponse) (err error) {
	err = r.client.Notify(request, response)
	return
}

// AlipayFundAuthOperationDetailQuery 资金授权操作查询接口
// https://opendocs.alipay.com/open/064jhg?scene=common&pathHash=44be9c20
func (r *OpenPayRouter) AlipayFundAuthOperationDetailQuery(request request.AlipayFundAuthOperationDetailQueryRequest, response *response.AlipayFundAuthOperationDetailQueryResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.fund.auth.operation.detail.query", &request, response)
	return
}

// AlipayFundAuthOrderAppFreeze 线上资金授权冻结接口
// https://opendocs.alipay.com/open/064jhe?pathHash=629fa9a5
func (r *OpenPayRouter) AlipayFundAuthOrderAppFreeze(request request.AlipayFundAuthOrderAppFreezeRequest, response *response.AlipayFundAuthOrderAppFreezeResponse) (reqData map[string]interface{}, err error) {
	response.OrderStr, reqData, err = r.client.GenerateStr("alipay.fund.auth.order.app.freeze", &request)
	return
}

// AlipayFundAuthOrderUnfreeze 资金授权解冻接口
// https://opendocs.alipay.com/open/02f3v8?scene=common&pathHash=c2b17437
func (r *OpenPayRouter) AlipayFundAuthOrderUnfreeze(request request.AlipayFundAuthOrderUnfreezeRequest, response *response.AlipayFundAuthOrderUnfreezeResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.fund.auth.order.unfreeze", &request, response)
	return
}

// AlipayTradeAppPay app支付接口2.0
// https://opendocs.alipay.com/open/e65d4f60_alipay.trade.app.pay
func (r *OpenPayRouter) AlipayTradeAppPay(request request.AlipayTradeAppPayRequest, response *response.AlipayTradeAppPayResponse) (reqData map[string]interface{}, err error) {
	response.OrderStr, reqData, err = r.client.GenerateStr("alipay.trade.app.pay", &request)
	return
}

// AlipayTradeCreate 统一收单交易创建接口
// https://opendocs.alipay.com/mini/6039ed0c_alipay.trade.create?pathHash=779dc517&ref=api&scene=de4d6a1e0c6e423b9eefa7c3a6dcb7a5
func (r *OpenPayRouter) AlipayTradeCreate(request request.AlipayTradeCreateRequest, response *response.AlipayTradeCreateResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.trade.create", &request, response)
	return
}

// AlipayTradeQuery 统一收单交易查询
// https://opendocs.alipay.com/mini/83b6c9a9_alipay.trade.query
func (r *OpenPayRouter) AlipayTradeQuery(request request.AlipayTradeQueryRequest, response *response.AlipayTradeQueryResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.trade.query", &request, response)
	return
}

// AlipayTradePay 线上资金授权冻结接口
// https://opendocs.alipay.com/mini/6039ed0c_alipay.trade.create
func (r *OpenPayRouter) AlipayTradePay(request request.AlipayTradePayRequest, response *response.AlipayTradePayResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.trade.pay", &request, response)
	return
}

// AlipayTradeClose 统一收单交易关闭接口
// https://opendocs.alipay.com/mini/d6f6c80a_alipay.trade.close
func (r *OpenPayRouter) AlipayTradeClose(request request.AlipayTradeCloseRequest, response *response.AlipayTradeCloseResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.trade.close", &request, response)
	return
}

// AlipayUserAgreementQuery (支付宝个人代扣协议查询接口)
// https://opendocs.alipay.com/open/3dab71bc_alipay.user.agreement.query?scene=common&pathHash=6706b504
func (r *OpenPayRouter) AlipayUserAgreementQuery(request request.AlipayUserAgreementQueryRequest, response *response.AlipayUserAgreementQueryResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.user.agreement.query", &request, response)
	return
}

// AlipayUserAgreementPageSign (支付宝个人代扣协议页面签约接口)
func (r *OpenPayRouter) AlipayUserAgreementPageSign(request request.AlipayUserAgreementPageSignRequest, response *response.AlipayUserAgreementPageSignResponse) (reqData map[string]interface{}, err error) {
	response.SignStr, reqData, err = r.client.GenerateStr("alipay.user.agreement.page.sign", &request)
	return
}
