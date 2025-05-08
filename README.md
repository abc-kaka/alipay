# alipay
支付宝 AliPay SDK for Go, 集成简单，功能完善，持续更新，支持公钥证书和普通公钥进行签名和验签。

# 安装
```
go get github.com/abc-kaka/alipay
```
# 使用demo

```go
package main

import (
	"fmt"
	"github.com/abc-kaka/alipay/open/pay/request"
	"github.com/abc-kaka/alipay/open/pay/response"
	"github.com/abc-kaka/alipay"
)

func main() {
	// 获取配置
	config := alipay.NewConfig()
	config.AppId = ""           // 应用app_id
	config.AppAuthToken = ""    // 授权令牌
	config.PrivateKey = ""      // 应用私钥
	config.PublicKey = ""       // 应用公钥
	config.AlipayPublicKey = "" // 支付宝公钥
	// 客户端
	client := alipay.NewClient(config)
	// 支付宝路由，支付宝api通过这个路由调用
	alipayRoute := alipay.NewRoute(client)

	// 支付回调验签、返回回调参数
	var notifyResp response.NotifyResponse
	err := alipayRoute.Open.Pay.VerifyNotify(request.VerifyNotifyRequest{
		Params: "", // 回调参数
	}, &notifyResp)
	if err != nil {
		panic("验签失败！")
	}

	// 预授权转支付
	var fundAuthDetail response.AlipayFundAuthOperationDetailQueryResponse
	reqData, err := alipayRoute.Open.Pay.AlipayFundAuthOperationDetailQuery(request.AlipayFundAuthOperationDetailQueryRequest{
		OutRequestNo: cast.ToString(v2OrderDepositMinus.OutTradeNo),
		OutOrderNo:   cast.ToString(v2OrderDepositMinus.OutTradeNo),
	}, &fundAuthDetail)
	if err != nil {
		panic("预授权转支付失败！")
	}
	fmt.Println("请求参数、结果", reqData)
}
```

# 项目目录结构说明
| 目录 | 说明 |
|:-------|:------|
|common|通用工具目录|
|mini|小程序api的请求参数、请求结果结构体|
|open|开放api的请求参数、请求结果结构体，包含人员、支付|
|client.go|客户端|
|config.go|配置|
|route.go|api路由文件|
# 项目扩展
可以看到目前的接口并不全，我这边只同步了部分接口，但扩展其实也很简单，只需要去支付宝官网把api接口的请求参数、请求结果copy到本项目，再添加对应的路由即可。
## 添加支付接口示例
### 支付宝支付接口
https://opendocs.alipay.com/mini/6039ed0c_alipay.trade.create?pathHash=779dc517&scene=de4d6a1e0c6e423b9eefa7c3a6dcb7a5
### 请求参数
创建请求参数文件：open/pay/request/alipay_trade_pay_request.go
```go
package request

import "encoding/json"

type AlipayTradePayRequest struct {
	NotifyUrl  string
	BizContent AlipayTradePayBizContent
}

// 普通预授权冻结押金转支付接口说明：https://opendocs.alipay.com/open/02cdx8?scene=34&pathHash=dd2813f3
type AlipayTradePayBizContent struct {
	OutTradeNo      string                 `json:"out_trade_no,omitempty"`
	TotalAmount     float64                `json:"total_amount,omitempty"`
	Subject         string                 `json:"subject,omitempty"`
	ProductCode     string                 `json:"product_code,omitempty"`
	AuthNo          string                 `json:"auth_no,omitempty"`
	AuthConfirmMode string                 `json:"auth_confirm_mode,omitempty"`
	ExtendParams    map[string]interface{} `json:"extend_params,omitempty"`
	BusinessParams  map[string]interface{} `json:"business_params,omitempty"`

	// 直付通参数
	BuyerId     string                 `json:"buyer_id,omitempty"`     // 用户支付宝uid
	SellerId    string                 `json:"seller_id,omitempty"`    // 卖家支付宝uid
	SubMerchant map[string]interface{} `json:"sub_merchant,omitempty"` // 二级商户信息
	SettleInfo  map[string]interface{} `json:"settle_info,omitempty"`  // 结算详细信息
}

// ToMap 转map
func (r *AlipayTradePayRequest) ToMap(v any) map[string]string {
	bizContent, _ := json.Marshal(r.BizContent)
	return map[string]string{
		"notify_url":  r.NotifyUrl,
		"biz_content": string(bizContent),
	}
}
```
### 请求结果
创建文件：open/pay/response/alipay_trade_create_response.go
```go
package response

// AlipayTradeCreateResponse (统一收单交易创建接口)-响应
type AlipayTradeCreateResponse struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}
```
### 添加api接口
文件：router_open_pay.go
```go
package alipay

import (
	"github.com/abc-kaka/alipay/open/pay/request"
	"github.com/abc-kaka/alipay/open/pay/response"
)

// OpenPayRouter H5&移动APP - 支付产品 - 路由
type OpenPayRouter struct {
	client Client
}

// NewOpenPayRouter 创建H5&移动APP - 支付产品 - 路由
func NewOpenPayRouter(client Client) *OpenPayRouter {
	return &OpenPayRouter{client: client}
}

// AlipayTradeCreate 统一收单交易创建接口
// https://opendocs.alipay.com/mini/6039ed0c_alipay.trade.create?pathHash=779dc517&ref=api&scene=de4d6a1e0c6e423b9eefa7c3a6dcb7a5
func (r *OpenPayRouter) AlipayTradeCreate(request request.AlipayTradeCreateRequest, response *response.AlipayTradeCreateResponse) (reqData map[string]interface{}, err error) {
	reqData, err = r.client.Exec("alipay.trade.create", &request, response)
	return
}
```
### 绑定到路由上
文件：route.go
```go
package alipay

type Route struct {
	Open *OpenRoute
}

type OpenRoute struct {
	Pay    *OpenPayRouter
}

func NewRoute(client Client) *Route {
	return &Route{
		Open: &OpenRoute{
			Pay:    NewOpenPayRouter(client),
		},
	}
}
```
恭喜你，完成接口的添加！
# 感谢您的支持
如果对您有帮助，请作者喝杯咖啡吧！万分感谢！

![支付宝](https://note.youdao.com/yws/api/image/normal/1644505774761?userId=weixinobU7Vjv5l7n-F3XTOaL8hr3saxAc)
![微信](https://note.youdao.com/yws/api/image/normal/1644505774761?userId=weixinobU7Vjv5l7n-F3XTOaL8hr3saxAc)
