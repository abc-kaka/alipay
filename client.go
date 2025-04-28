package alipay

import (
	"encoding/json"
	"fmt"
	"github.com/abc-kaka/alipay/common/req"
	"github.com/abc-kaka/alipay/common/util"
	"net/url"
	"strings"
)

// Client api请求客户端
type Client struct {
	config *Config
}

// NewClient 创建一个api客户端
func NewClient(config *Config) Client {
	return Client{config: config}
}

// Exec 客户端执行api请求
func (c *Client) Exec(apiCode string, request req.Request, response interface{}) (reqData map[string]interface{}, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("alipay-error：" + err.Error())
		}
	}()
	// 组装参数
	paramsMap, err := c.generateParam(apiCode, request)
	if err != nil {
		return
	}
	// 任务请求
	body, err := util.HttpRequest("POST", c.config.ServerUrl, paramsMap)

	if err != nil {
		return
	}
	reqData = map[string]interface{}{
		"params": paramsMap,
		"body":   body,
	}
	// 映射到结果集
	err = c.mapResult(apiCode, body, response)
	return
}

func (c *Client) GenerateStr(apiCode string, request req.Request) (str string, reqData map[string]any, err error) {
	publicParams, err := c.generateParam(apiCode, request)

	params := url.Values{}
	for key, value := range publicParams {
		params.Add(key, value)
	}
	str = params.Encode()

	reqData = map[string]interface{}{
		"params": publicParams,
		"body":   str,
	}
	return
}

//	-------------------------内部实现----------------------

// generateParam 生成参数
func (c *Client) generateParam(apiCode string, request req.Request) (params map[string]string, err error) {
	// 公共参数
	publicParams := c.config.GetApiPublicParams(apiCode)
	// 请求参数
	requestParam := request.ToMap(request)
	// 合并参数
	params = make(map[string]string)
	for k, v := range publicParams {
		if _, ok := publicParams[k]; ok && v != "" {
			params[k] = v
		}
	}
	for k, v := range requestParam {
		if _, ok := requestParam[k]; ok && v != "" {
			params[k] = v
		}
	}
	// 获取签名
	if c.config.PrivateKey == "" {
		err = fmt.Errorf("支付宝应用配置异常！")
		return
	}
	encryptedDataStr, err := util.GenerateSign(params, c.config.PrivateKey)
	if err != nil {
		return
	}
	params["sign"] = encryptedDataStr
	return
}

// mapResult 请求结果组装
func (c *Client) mapResult(apiCode string, body string, response interface{}) (err error) {
	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return
	}
	data, err := c.extractResult(apiCode, result)
	if err != nil {
		return
	}
	err = util.MapToStruct(data, &response)
	return
}

// extractResult 获取请求结果
func (c *Client) extractResult(apiCode string, result map[string]interface{}) (data map[string]interface{}, err error) {
	// 获取异常code
	errResponse, ok := result["error_response"]
	if ok {
		err = fmt.Errorf(errResponse.(map[string]interface{})["sub_msg"].(string))
		return
	}
	apiCode = strings.ReplaceAll(apiCode, ".", "_") + "_response"
	resultData, ok := result[apiCode]
	if !ok {
		err = fmt.Errorf("请求结果异常：获取key异常：" + apiCode)
		return
	}
	code, ok := resultData.(map[string]interface{})["code"].(string)
	if ok && code != "10000" {
		err = fmt.Errorf(resultData.(map[string]interface{})["sub_msg"].(string))
		return
	}
	data = resultData.(map[string]interface{})
	return
}
