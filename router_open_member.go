package alipay

import (
	"github.com/abc-kaka/alipay/open/member/request"
	"github.com/abc-kaka/alipay/open/member/response"
)

// OpenMemberRouter H5&移动APP - 会员产品 - 路由
type OpenMemberRouter struct {
	client Client
}

// NewOpenMemberRouter 创建H5&移动APP - 会员产品 - 路由
func NewOpenMemberRouter(client Client) *OpenMemberRouter {
	return &OpenMemberRouter{client: client}
}

// AlipaySystemOauthToken 换取授权访问令牌
// https://opendocs.alipay.com/open/e6859bb3_alipay.system.oauth.token?scene=common&pathHash=7de84146
func (r *OpenMemberRouter) AlipaySystemOauthToken(request request.AlipaySystemOauthTokenRequest, response *response.AlipaySystemOauthTokenResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.system.oauth.token", &request, &response)
	return
}

// AlipayUserInfoShare 支付宝会员授权信息查询接口
// https://opendocs.alipay.com/open/01834fa5_alipay.user.info.share?scene=common&pathHash=af6ac2d2
func (r *OpenMemberRouter) AlipayUserInfoShare(request request.AlipayUserInfoShareRequest, response *response.AlipayUserInfoShareResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.user.info.share", &request, &response)
	return
}
